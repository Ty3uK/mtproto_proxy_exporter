package main

import (
  "flag"
  "log"
  "net/http"
  "time"

  "github.com/prometheus/client_golang/prometheus/promhttp"

  configPkg "github.com/Ty3uK/mtproto_proxy_exporter/config"
  metricsPkg "github.com/Ty3uK/mtproto_proxy_exporter/metrics"
  statsPkg "github.com/Ty3uK/mtproto_proxy_exporter/stats"
)

var (
  configPath = flag.String("config", "", "YML config path")
  help       = flag.Bool("help", false, "prints help")

  stats   = statsPkg.Stats{}
  metrics = metricsPkg.Metrics{
    List: make(map[string]metricsPkg.Item),
  }
)

var config configPkg.Config

func run() {
  for {
    stats.GetData(config.StatsAddress)

    for _, item := range metrics.List {
      item.SetValue(
        stats.GetNumberItem(item.StatName),
      )
    }

    time.Sleep(time.Duration(config.Interval) * time.Second)
  }
}

func initFromConfig() {
  println("LISTENING ON  :", config.Address)
  println("SCAN INTERVAL :", config.Interval)
  println()

  for _, configItem := range config.Metrics {
    metrics.AddItem(
      configItem.StatName,
      configItem.Name,
      configItem.Help,
    )

    println("FROM MTPROTO  :", configItem.StatName)
    println("TO PROMETHEUS :", configItem.Name)
    println()
  }
}

func main() {
  flag.Parse()

  if *help {
    configPkg.PrintHelp()
    return
  }

  config = configPkg.InitFromFile(*configPath)
  initFromConfig()

  http.Handle("/metrics", promhttp.Handler())
  go run()
  log.Fatal(http.ListenAndServe(config.Address, nil))
}
