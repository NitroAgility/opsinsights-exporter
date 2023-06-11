# OpsInsights-Exporter

[![Build and Publish](https://github.com/NitroAgility/ops-insights-exporter/actions/workflows/ci-docker.yml/badge.svg)](https://github.com/NitroAgility/ops-insights-exporter/actions/workflows/ci-docker.yml) [![Docker Pulls](https://img.shields.io/docker/pulls/nitroagility/opsinsights-exporter)](https://hub.docker.com/r/nitroagility/opsinsights-exporter)

Prometheus Exporter for OpsInsights, detect anomalies before they become business incidents.

<p align="center">
    <img src="https://github.com/nitroagility/ops-insights-exporter/blob/main/logo.png" width="150"/>
</p>

## Overview

`OpsInsights-Exporter` is a tool you can use to get insights to make sure your apps are working as intended.

It can often happen that the application runs correctly without raising any kind of error but at the same time it does not produce the expected data or business operations. Those anomalies can be difficult to discover before they impact business operations.

For instance let's suppose we have deployed an API to be used to collect data from external edge devices and we do expect at the least one update every 30 minutes by each device. Our API is running as expected however because of a firewall change the edge devices are unable to make requests to our API. How can we detect this issue?

`OpsInsights-Exporter` would allow our operations team to be notified if there are no records in the database for more than 30 minutes for each device.

This would allow us to detect anomalies before they become business incidents.

Notifications are fired if data expectations are not met. Data exepectations can query different data sources (for instance DB, S3, API).

## Configuration File

Below a sample `OpsInsights-Exporter`config file:

```yaml
settings:
  app_name: edge-devices
  sleep_for: 30
datasources:
    orders_db:
        type: DB
        arguments:
            dialect: PG
            host: env:PG_DATABASE_HOST
            port: env:PG_DATABASE_PORT
            username: env:PG_DATABASE_USER
            password: env:PG_DATABASE_PASSWORD
            database: env:PG_DATABASE_NAME
metrics:
    edge_device_up:
        type: counter
        description: Edge device is up and running
        labels: device_name, device_ip
    edge_device_down:
        type: counter
        description: Edge device is missing
        labels: device_name
expectations:
    devices_up:
        datasource: pg_order_db
        metrices:
            - edge_device_up
        sql: >
            SELECT 
                device AS device_name,
                ip     AS device_ip
            FROM edge_devices
            WHERE last_update >= NOW() - INTERVAL '30 minutes'
    devices_down:
        datasource: pg_order_db
        metrices:
            - edge_device_up
        sql: >
            SELECT 
                device AS device_name,
                ip     AS device_ip
            FROM edge_devices
            WHERE last_update < NOW() - INTERVAL '30 minutes'
```
