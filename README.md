<p align="center">
<img height="256" src="logo.svg" alt="oteldb svg logo">
</p>

# oteldb [![codecov](https://img.shields.io/codecov/c/github/go-faster/oteldb?label=cover)](https://codecov.io/gh/go-faster/oteldb) [![experimental](https://img.shields.io/badge/-experimental-blueviolet)](https://go-faster.org/docs/projects/status#experimental)

The OpenTelemetry-compatible telemetry aggregation, storage and processing.

> [!WARNING]  
> Work in progress.

## Storage

The oteldb is stateless and uses external storage systems for data persistence, processing and aggregation.

Currently, we focus on ClickHouse for realtime queries on hot/warm data.

## Query

Supported query languages:
- LogQL (Loki) for logs
- TraceQL (Tempo) for traces
- PromQL (Prometheus) for metrics

### Prometheus Compatibility

```console
$ promql-compliance-tester -config-file promql-test-queries.yml -config-file test.oteldb.yml
Total: 547 / 548 (99.82%) passed, 0 unsupported
```

## Local development

Setups oteldb, clickhouse server, grafana, and telemetry generators:

```shell
docker compose -f dev/local/ch/docker-compose.yml up -d
```

You can open Grafana dashboard at http://localhost:3000/d/oteldb/oteldb
