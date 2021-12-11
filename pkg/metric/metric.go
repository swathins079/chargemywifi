package metric

import (
	"fmt"
	"log"
	"regexp"
	"sync"

	"github.com/swathinsankaran/chargemywifi/pkg/alert"
	"github.com/swathinsankaran/chargemywifi/pkg/model"
)

type Metric interface {
	Observer
	Setter
	Getter
}

type Observer interface {
	OnNotify()
}

type Getter interface {
	Rgx() *regexp.Regexp
}

type Setter interface {
	Value(string)
}

type metric struct {
	label model.Label
	value sync.Map
	rgx   *regexp.Regexp
}

func New(label model.Label) Metric {
	m := &metric{}
	m.label = label
	rgxStr := fmt.Sprintf("<label id=\"%s\">(.*?)</label>", m.label)
	m.rgx = regexp.MustCompile(rgxStr)
	return m
}

func (d *metric) OnNotify() {
	val, ok := d.value.Load(d.label)
	if !ok {
		log.Println("Failed to get value")
		return
	}
	log.Println("Got a Value:", val)
	alert, ok := alert.AlertPool[fmt.Sprintf("%s-%s-%s", model.OperatingSystem, d.label, val)]
	if !ok {
		return
	}
	alert.Push()
}

func (d *metric) Rgx() *regexp.Regexp {
	return d.rgx
}

func (d *metric) Value(value string) {
	d.value.Store(d.label, value)
}
