#!/usr/bin/env bash

go mod download
go build -o ops-insights-exporter ./cmd/ops-insights-exporter