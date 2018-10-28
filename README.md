# MTProto Proxy Exporter

[![Go Report Card](https://goreportcard.com/badge/github.com/Ty3uK/mtproto_proxy_exporter)](https://goreportcard.com/report/github.com/Ty3uK/mtproto_proxy_exporter)

Prometheus exporter for MTProto Proxy Stats. At this moment working only with simple numeric values.

## Configuration

Configuration must be stored in YAML file and passed through `-config` command.
Sample configuration can be viewed with `-help` command.

```yaml
address:         ":8080"
stats_address:   "http://localhost:2398/stats"
interval:        5
request_timeout: 10

metrics:
  - stat_name: "total_special_connections"
    name:      "mtproto_proxy_users_count"
    help:      "Users count"

  - stat_name: "active_connections"
    name:      "mtproto_active_connections_count"
    help:      "Active connections count"

  - stat_name: "uptime"
    name:      "mtproto_proxy_uptime"
    help:      "Uptime"
```

| Property      | Type     | Description             |
|---------------|----------|-------------------------|
| address       | String   | Listening address       |
| stats_address | String   | MTProto Proxy stats URL |
| interval      | Int      | Scan interval           |
| metrics       | Metric[] | Mapping items           |

| Property  | Type   | Description                   |
|-----------|--------|-------------------------------|
| stat_name | String | Input MTProto Proxy stat name |
| name      | String | Output Prometheus metric name |
| help      | String | Output Prometheus metric help |

## Building and running

```sh
go get github.com/Ty3uK/mtproto_proxy_exporter
cd $GOPATH/github.com/Ty3uK/mtproto_proxy_exporter
go build mtproto_proxy_exporter.go
./mtproto_proxy_exporter
```

To see config file help:

```sh
./mtproto_proxy_exporter -help
```
