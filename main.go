package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/swathinsankaran/chargemywifi/pkg/metric"
	"github.com/swathinsankaran/chargemywifi/pkg/model"
	"github.com/swathinsankaran/chargemywifi/pkg/statscollector"
)

func main() {
	log.Println("Application Started.")
	defer exit()

	dBattery := metric.New(model.LabelBatteryQuantity)
	dChargeStatus := metric.New(model.LabelChargeStatus)

	sc := statscollector.New()
	sc.Register(dBattery)
	sc.Register(dChargeStatus)

	log.Println("Starting cron jobs.")
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Minute().Do(sc.Collect)
	s.Every(1).StartAt(time.Now().Add(30 * time.Second)).Minute().Do(sc.Notify)
	s.StartAsync()
}

func exit() {
	exitChannel := make(chan os.Signal)
	signal.Notify(exitChannel, os.Kill, os.Interrupt)
	select {
	case <-exitChannel:
		log.Println("Application stopped.")
	}
}
