package main

import (
	"context"
	"flag"
	"os"

	"github.com/go-faster/errors"
	"github.com/go-faster/sdk/app"
	"go.uber.org/zap"

	"github.com/go-faster/oteldb/internal/autologs"
	"github.com/go-faster/oteldb/internal/autopyro"
	"github.com/go-faster/oteldb/internal/autozpages"
)

func main() {
	app.Run(func(ctx context.Context, lg *zap.Logger, m *app.Metrics) error {
		shutdown, err := autozpages.Setup(m.TracerProvider())
		if err != nil {
			return errors.Wrap(err, "setup zPages")
		}
		defer func() {
			_ = shutdown(context.Background())
		}()
		if ctx, err = autologs.Setup(ctx, m); err != nil {
			return errors.Wrap(err, "setup logs")
		}

		set := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
		cfgPath := set.String("config", "", "Path to config (defaults to oteldb.yml)")
		if err := set.Parse(os.Args[1:]); err != nil {
			return err
		}

		cfg, err := loadConfig(*cfgPath)
		if err != nil {
			return errors.Wrap(err, "load config")
		}
		{
			stop, err := autopyro.Setup(ctx)
			if err != nil {
				return errors.Wrap(err, "setup pyroscope")
			}
			defer func() {
				_ = stop(context.Background())
			}()
		}

		root, err := newApp(ctx, cfg, m)
		if err != nil {
			return errors.Wrap(err, "setup")
		}
		return root.Run(ctx)
	})
}
