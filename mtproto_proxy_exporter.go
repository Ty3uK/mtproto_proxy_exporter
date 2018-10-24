package main

import (
	"flag"
	"fmt"
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
		err := stats.GetData(config.StatsAddress)
		if err != nil {
			log.Printf("error: could not get stats data: %v\n", err)
		} else {
			for _, item := range metrics.List {
				item.SetValue(
					stats.GetNumberItem(item.StatName),
				)
			}
		}
		time.Sleep(time.Duration(config.Interval) * time.Second)
	}
}

func initFromConfig() {
	fmt.Println("LISTENING ON  :", config.Address)
	fmt.Println("SCAN INTERVAL :", config.Interval)
	fmt.Println()

	for _, configItem := range config.Metrics {
		metrics.AddItem(
			configItem.StatName,
			configItem.Name,
			configItem.Help,
		)

		fmt.Println("FROM MTPROTO  :", configItem.StatName)
		fmt.Println("TO PROMETHEUS :", configItem.Name)
		fmt.Println()
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
