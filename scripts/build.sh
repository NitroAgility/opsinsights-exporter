#!/usr/bin/env bash

go mod download
go build -o opsinsights-exporter ./pkg/cmd/opsinsights-exporter