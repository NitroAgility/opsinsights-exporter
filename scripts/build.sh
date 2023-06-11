#!/usr/bin/env bash

go mod download
go build -o ops-insights-exporter ./pkg/cmd/ops-insights-exporter