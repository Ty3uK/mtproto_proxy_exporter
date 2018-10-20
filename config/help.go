package config

// PrintHelp displays help message about config file
func PrintHelp() {
	println(`
Config file must be written in YAML format and passed through -config option. Example config file file described below.

address:       ":8080"
stats_address: "http://localhost:2398/stats"
interval:       5

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
  `)
}
