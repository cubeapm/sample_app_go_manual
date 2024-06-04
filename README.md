# Go manual instrumentation with OpenTelemetry

This is a sample app to demonstrate how to instrument Go code with OpenTelemetry.

This repository is inentionally designed to work with any OpenTelemetry backend, not just CubeAPM. In fact, it can even work without any OpenTelemetry backend (by dumping traces to console, which is also the default behaviour).

## Setup

Clone this repository and go to the project directory. Then run the following commands

```
go mod download -x

export OTEL_SERVICE_NAME=cube_sample_go_manual
export OTEL_EXPORTER_OTLP_COMPRESSION=gzip
# print traces on console
export OTEL_LOG_LEVEL=debug
# send traces to CubeAPM
# export OTEL_EXPORTER_OTLP_TRACES_ENDPOINT=http://<ip_address_of_cubeapm_server>:4318/v1/traces
go run .
```

## Contributing

Please feel free to raise PR for any enhancements - additional service integrations, library version updates, documentation updates, etc.
