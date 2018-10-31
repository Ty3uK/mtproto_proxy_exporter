package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/Ty3uK/mtproto_proxy_exporter/config"
	metricsPkg "github.com/Ty3uK/mtproto_proxy_exporter/metrics"
	statsPkg "github.com/Ty3uK/mtproto_proxy_exporter/stats"
)

var (
	configPath = flag.String("config", "", "YML config path")
	help       = flag.Bool("help", false, "prints help")
	versionCmd = flag.Bool("version", false, "prints information about version and build")

	stats   = statsPkg.Stats{}
	metrics = metricsPkg.Metrics{
		List: make(map[string]metricsPkg.Item),
	}

	version = "current"
	build   = "master build"
)

func run() {
	for {
		err := stats.GetData(config.Config.StatsAddress)
		if err != nil {
			log.Printf("error: could not get stats data: %v\n", err)
		} else {
			for _, item := range metrics.List {
				item.SetValue(
					stats.GetNumberItem(item.StatName),
				)
			}
		}
		time.Sleep(time.Duration(config.Config.Interval) * time.Second)
	}
}

func initFromConfig() {
	fmt.Println("LISTENING ON    :", config.Config.Address)
	fmt.Println("SCAN INTERVAL   :", config.Config.Interval)
	fmt.Println("REQUEST TIMEOUT :", config.Config.RequestTimeout)
	fmt.Println()

	for _, configItem := range config.Config.Metrics {
		err := metrics.AddItem(
			configItem.StatName,
			configItem.Name,
			configItem.Help,
		)

		if err != nil {
			log.Printf("could not add metrics item: %v\n", err)
			continue
		}

		fmt.Println("FROM MTPROTO    :", configItem.StatName)
		fmt.Println("TO PROMETHEUS   :", configItem.Name)
		fmt.Println()
	}
}

func main() {
	flag.Parse()

	if *versionCmd {
		config.PrintVersion(version, build)
		return
	}

	if *help {
		config.PrintHelp()
		return
	}

	err := config.InitFromFile(*configPath)
	if err != nil {
		log.Fatalf("could not init config from file: %v", err)
	}
	initFromConfig()

	http.Handle("/metrics", promhttp.Handler())
	go run()
	log.Fatal(http.ListenAndServe(config.Config.Address, nil))
}
