package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
	toast "gopkg.in/toast.v1"
)

func main() {
	log.Println("Application Started.")
	defer exit()

	ws := &wifiStats{}

	dBattery := newData(batteryQuantity)
	dChargeStatus := newData(chargeStatus)
	ws.l = append(ws.l, dBattery, dChargeStatus)

	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Minute().Do(ws.Get)
	s.Every(1).StartAt(time.Now().Add(1 * time.Minute)).Minute().Do(dBattery.checkData)
	s.Every(1).StartAt(time.Now().Add(1 * time.Minute)).Minute().Do(dChargeStatus.checkData)
}

func exit() {
	exitChannel := make(chan os.Signal)
	signal.Notify(exitChannel, os.Kill, os.Interrupt)
	select {
	case <-exitChannel:
		log.Println("Application stopped.")
		os.Exit(0)
	}
}

type Label string

const (
	batteryQuantity Label = "lDashBatteryQuantity"
	chargeStatus    Label = "lDashChargeStatus"
)

const (
	JIOFIURL string = "http://jiofi.local.html/cgi-bin/en-jio/mStatus.html"
)

type Notification struct {
	Title       string
	Description string
}

var toastData = map[string]Notification{
	fmt.Sprintf("%s-%s", chargeStatus, "Fully Charged"): {Title: "Charging Status", Description: "Charging Status is %s, Please disconnect charger."},
	fmt.Sprintf("%s-%s", batteryQuantity, "100%"):       {Title: "Battery Level", Description: "Battery Level is %s, Please disconnect charger."},
	fmt.Sprintf("%s-%s", batteryQuantity, "50%"):        {Title: "Battery Level", Description: "Battery Level is %s."},
	fmt.Sprintf("%s-%s", batteryQuantity, "10%"):        {Title: "Battery Level", Description: "Low Battery Level. Please connect charger"},
	fmt.Sprintf("%s-%s", batteryQuantity, "5%"):         {Title: "Battery About to die", Description: "Battery Level critical. Please connect charger immediately"},
}

type data struct {
	label Label
	value sync.Map
	Rgx   *regexp.Regexp
}

type wifiStats struct {
	l []*data
}

func newData(label Label) *data {
	d := &data{}
	d.label = label
	rgxStr := fmt.Sprintf("<label id=\"%s\">(.*?)</label>", d.label)
	d.Rgx = regexp.MustCompile(rgxStr)
	return d
}

func (ws *wifiStats) Get() {
	resp, err := http.Get(JIOFIURL)
	checkError(err)
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	htmlSource := string(html)
	for _, d := range ws.l {
		d.value.Store(d.label, d.Rgx.FindStringSubmatch(htmlSource)[1])
	}
}

func (d *data) checkData() {
	val, ok := d.value.Load(d.label)
	if !ok {
		log.Fatalf("Failed to get value")
	}
	log.Println("Got a Value: ", val)
	toastData, ok := toastData[fmt.Sprintf("%s-%s", d.label, val)]
	if !ok {
		return
	}
	notification := toast.Notification{
		AppID:   "{1AC14E77-02E7-4E5D-B744-2EB1AE5198B7}\\WindowsPowerShell\\v1.0\\powershell.exe",
		Title:   toastData.Title,
		Message: fmt.Sprintf(toastData.Description, val),
		Audio:   toast.Default,
		Actions: []toast.Action{
			{"protocol", "Cancel", ""},
		},
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}

// checkError is utility function which checks for error.
func checkError(err error) {
	if err != nil {
		log.Fatal("Failure: ", err)
	}
}
