package main

import (
	"context"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-faster/errors"
	sdkapp "github.com/go-faster/sdk/app"
	"github.com/go-faster/sdk/zctx"
	"github.com/prometheus/prometheus/promql"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/confmap/converter/expandconverter"
	"go.opentelemetry.io/collector/otelcol"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"

	"github.com/go-faster/oteldb/internal/httpmiddleware"
	"github.com/go-faster/oteldb/internal/logql"
	"github.com/go-faster/oteldb/internal/logql/logqlengine"
	"github.com/go-faster/oteldb/internal/lokiapi"
	"github.com/go-faster/oteldb/internal/lokihandler"
	"github.com/go-faster/oteldb/internal/otelreceiver"
	"github.com/go-faster/oteldb/internal/promapi"
	"github.com/go-faster/oteldb/internal/promhandler"
	"github.com/go-faster/oteldb/internal/tempoapi"
	"github.com/go-faster/oteldb/internal/tempohandler"
	"github.com/go-faster/oteldb/internal/traceql/traceqlengine"
)

// App contains application dependencies and services.
type App struct {
	cfg Config

	services map[string]func(context.Context) error
	otelStorage

	lg      *zap.Logger
	metrics *sdkapp.Metrics
}

func newApp(ctx context.Context, cfg Config, m *sdkapp.Metrics) (_ *App, err error) {
	cfg.setDefaults()

	app := &App{
		cfg:      cfg,
		services: map[string]func(context.Context) error{},
		lg:       zctx.From(ctx),
		metrics:  m,
	}

	{
		dsn := os.Getenv("CH_DSN")
		if dsn == "" {
			dsn = cfg.DSN
		}
		store, err := setupCH(ctx, dsn, app.lg, m)
		if err != nil {
			return nil, errors.Wrapf(err, "create storage")
		}
		app.otelStorage = store
	}

	if err := app.setupCollector(); err != nil {
		return nil, errors.Wrap(err, "otelcol")
	}
	if err := app.trySetupTempo(); err != nil {
		return nil, errors.Wrap(err, "tempo")
	}
	if err := app.trySetupLoki(); err != nil {
		return nil, errors.Wrap(err, "loki")
	}
	if err := app.trySetupProm(); err != nil {
		return nil, errors.Wrap(err, "prom")
	}

	return app, nil
}

func addOgen[
	R httpmiddleware.OgenRoute,
	Server interface {
		httpmiddleware.OgenServer[R]
		http.Handler
	},
](
	app *App,
	name string,
	server Server,
	defaultPort string,
) {
	lg := app.lg.Named(name)

	addr := os.Getenv(strings.ToUpper(name) + "_ADDR")
	if addr == "" {
		addr = defaultPort
	}

	app.services[name] = func(ctx context.Context) error {
		lg := lg.With(zap.String("addr", addr))
		lg.Info("Starting HTTP server")

		routeFinder := httpmiddleware.MakeRouteFinder(server)
		httpServer := &http.Server{
			Addr: addr,
			Handler: httpmiddleware.Wrap(server,
				httpmiddleware.InjectLogger(zctx.From(ctx)),
				httpmiddleware.Instrument("oteldb", routeFinder, app.metrics),
				httpmiddleware.LogRequests(routeFinder),
			),
			ReadHeaderTimeout: 15 * time.Second,
		}

		parentCtx := ctx
		g, ctx := errgroup.WithContext(ctx)
		g.Go(func() error {
			<-ctx.Done()
			lg.Info("Shutting down")

			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			return httpServer.Shutdown(ctx)
		})
		g.Go(func() error {
			if err := httpServer.ListenAndServe(); err != nil {
				if errors.Is(err, http.ErrServerClosed) && parentCtx.Err() != nil {
					lg.Info("HTTP server closed gracefully")
					return nil
				}
				return errors.Wrap(err, "http server")
			}
			return nil
		})
		return g.Wait()
	}
}

func (app *App) trySetupTempo() error {
	q := app.traceQuerier
	if q == nil {
		return nil
	}
	cfg := app.cfg.Tempo
	cfg.setDefaults()

	engine := traceqlengine.NewEngine(app.traceQuerier, traceqlengine.Options{
		TracerProvider: app.metrics.TracerProvider(),
	})
	tempo := tempohandler.NewTempoAPI(q, engine)

	s, err := tempoapi.NewServer(tempo,
		tempoapi.WithTracerProvider(app.metrics.TracerProvider()),
		tempoapi.WithMeterProvider(app.metrics.MeterProvider()),
	)
	if err != nil {
		return err
	}

	addOgen[tempoapi.Route](app, "tempo", s, cfg.Bind)
	return nil
}

func (app *App) trySetupLoki() error {
	q := app.logQuerier
	if q == nil {
		return nil
	}
	cfg := app.cfg.LokiConfig
	cfg.setDefaults()

	engine := logqlengine.NewEngine(q, logqlengine.Options{
		ParseOptions: logql.ParseOptions{
			AllowDots: true,
		},
		LookbackDuration: cfg.LookbackDelta,
		TracerProvider:   app.metrics.TracerProvider(),
	})
	loki := lokihandler.NewLokiAPI(q, engine)

	s, err := lokiapi.NewServer(loki,
		lokiapi.WithTracerProvider(app.metrics.TracerProvider()),
		lokiapi.WithMeterProvider(app.metrics.MeterProvider()),
	)
	if err != nil {
		return err
	}

	addOgen[lokiapi.Route](app, "loki", s, cfg.Bind)
	return nil
}

func (app *App) trySetupProm() error {
	q := app.metricsQuerier
	if q == nil {
		return nil
	}
	cfg := app.cfg.Prometheus
	cfg.setDefaults()

	engine := promql.NewEngine(promql.EngineOpts{
		// NOTE: zero-value MaxSamples and Timeout makes
		// all queries to fail with error.
		MaxSamples:           cfg.MaxSamples,
		Timeout:              cfg.Timeout,
		LookbackDelta:        cfg.LookbackDelta,
		EnableAtModifier:     cfg.EnableAtModifier,
		EnableNegativeOffset: *cfg.EnableNegativeOffset,
		EnablePerStepStats:   cfg.EnablePerStepStats,
	})
	prom := promhandler.NewPromAPI(engine, q, q, promhandler.PromAPIOptions{})

	s, err := promapi.NewServer(prom,
		promapi.WithTracerProvider(app.metrics.TracerProvider()),
		promapi.WithMeterProvider(app.metrics.MeterProvider()),
		promapi.WithMiddleware(promhandler.TimeoutMiddleware()),
	)
	if err != nil {
		return err
	}

	addOgen[promapi.Route](app, "prom", s, cfg.Bind)
	return nil
}

func (app *App) setupCollector() error {
	conf, err := otelcol.NewConfigProvider(otelcol.ConfigProviderSettings{
		ResolverSettings: confmap.ResolverSettings{
			URIs: []string{"oteldb:/"},
			Providers: map[string]confmap.Provider{
				"oteldb": otelreceiver.NewMapProvider("oteldb", app.cfg.Collector),
			},
			Converters: []confmap.Converter{
				expandconverter.New(),
			},
		},
	})
	if err != nil {
		return errors.Wrap(err, "create otelcol config provider")
	}

	col, err := otelcol.NewCollector(otelcol.CollectorSettings{
		Factories: otelreceiver.Factories,
		BuildInfo: component.NewDefaultBuildInfo(),
		LoggingOptions: []zap.Option{
			zap.WrapCore(func(zapcore.Core) zapcore.Core {
				return app.lg.Core()
			}),
		},
		DisableGracefulShutdown: false,
		ConfigProvider:          conf,
		SkipSettingGRPCLogger:   false,
	})
	if err != nil {
		return errors.Wrap(err, "create collector")
	}

	app.services["otelcol"] = func(ctx context.Context) error {
		return col.Run(ctx)
	}
	return nil
}

// Run runs application.
func (app *App) Run(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)
	for _, s := range app.services {
		s := s
		g.Go(func() error {
			return s(ctx)
		})
	}
	return g.Wait()
}
