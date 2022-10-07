package main

import (
	"net/http"

	"github.com/j18e/hs110"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := run(); err != nil {
		logrus.Fatal(err)
	}
}

func run() error {
	http.HandleFunc("/probe", handlerFunc)

	listenAddr := ":8080"
	logrus.Infof("listening on %s", listenAddr)
	return http.ListenAndServe(listenAddr, nil)
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	target := params.Get("target")
	if target == "" {
		http.Error(w, "Target parameter is missing", http.StatusBadRequest)
		return
	}

	voltage := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "hs110_voltage_mv",
		Help: "current voltage of smart plug",
	})
	current := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "hs110_current_ma",
		Help: "current current of smart plug",
	})
	power := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "hs110_power_mw",
		Help: "current power output of smart plug",
	})
	total := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "hs110_total_wh",
		Help: "total consumption of smart plug",
	})

	registry := prometheus.NewRegistry()
	registry.MustRegister(voltage)
	registry.MustRegister(current)
	registry.MustRegister(power)
	registry.MustRegister(total)

	plug := hs110.NewPlug(target)
	readout, err := plug.Energy()
	if err != nil {
		http.Error(w, "Could not get readout from plug", http.StatusInternalServerError)
		logrus.Warn(err)
		return
	}

	voltage.Set(float64(readout.VoltageMV))
	current.Set(float64(readout.CurrentMA))
	power.Set(float64(readout.PowerMW))
	total.Add(float64(readout.TotalWH))

	h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	h.ServeHTTP(w, r)
}
