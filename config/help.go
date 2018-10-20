package config

// PrintHelp displays help message about config file
func PrintHelp() {
	println("Config file must be written in YAML format and passed through `-config` option. Example config file file described below.")
	println()
	println(`address:       ":8080"`)
	println(`stats_address: "http://localhost:2398/stats"`)
	println(`interval:       5`)
	println()
	println(`metrics:`)
	println(`  - stat_name: "total_special_connections"`)
	println(`    name:      "mtproto_proxy_users_count"`)
	println(`    help:      "Users count"`)
	println()
	println(`  - stat_name: "active_connections"`)
	println(`    name:      "mtproto_active_connections_count"`)
	println(`    help:      "Active connections count"`)
	println()
	println(`  - stat_name: "uptime"`)
	println(`    name:      "mtproto_proxy_uptime"`)
	println(`    help:      "Uptime"`)
}
