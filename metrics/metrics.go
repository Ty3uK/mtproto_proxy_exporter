package metrics

import "github.com/prometheus/client_golang/prometheus"

// Item represents single metrics item
type Item struct {
	StatName string
	handler  prometheus.Gauge
}

// SetValue writes new value to prometheus metric
func (metricsItem *Item) SetValue(value float64) {
	metricsItem.handler.Set(value)
}

// Metrics represents list of metrics items
type Metrics struct {
	List map[string]Item
}

// AddItem adds new metrics item to list
func (metric *Metrics) AddItem(statName string, name string, help string) {
	if _, ok := metric.List[name]; ok {
		prometheus.Unregister(metric.List[name].handler)
		delete(metric.List, name)
	}

	metric.List[name] = Item{
		StatName: statName,
		handler: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: name,
				Help: help,
			},
		),
	}
	prometheus.Register(metric.List[name].handler)
}
