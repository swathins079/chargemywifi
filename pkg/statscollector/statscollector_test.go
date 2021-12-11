package statscollector

import (
	"bytes"
	"errors"
	"log"
	"reflect"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/swathinsankaran/chargemywifi/pkg/metric"
	"github.com/swathinsankaran/chargemywifi/pkg/model"
)

func TestNew(t *testing.T) {
	cases := []struct {
		testcaseName string
		expected     StatsCollector
		got          StatsCollector
	}{
		{
			testcaseName: "WifiStats Object creation",
			expected:     &wifiStats{},
			got:          New(),
		},
	}

	for _, c := range cases {
		t.Run(c.testcaseName, func(t *testing.T) {
			if reflect.TypeOf(c.expected) != reflect.TypeOf(c.got) {
				t.Fatalf("Failed. Expected: %v, Got: %v", reflect.TypeOf(c.expected), reflect.TypeOf(c.expected))
			}
		})
	}
}

func TestRegister(t *testing.T) {
	cases := []struct {
		testcaseName string
	}{
		{
			testcaseName: "Register Method Functionality Test",
		},
	}

	for _, c := range cases {
		t.Run(c.testcaseName, func(t *testing.T) {
			ws := &wifiStats{stats: make(map[metric.Metric]struct{})}
			newMetric := metric.New(model.LabelBatteryQuantity)
			ws.Register(newMetric)
			if _, ok := ws.stats[newMetric]; !ok {
				t.Fatalf("Failed. Register method not ok")
			}
		})
	}
}

func TestDeregister(t *testing.T) {
	cases := []struct {
		testcaseName string
	}{
		{
			testcaseName: "Deregister Method Functionality Test",
		},
	}

	for _, c := range cases {
		t.Run(c.testcaseName, func(t *testing.T) {
			ws := &wifiStats{stats: make(map[metric.Metric]struct{})}
			newMetric := metric.New(model.LabelBatteryQuantity)
			ws.Register(newMetric)
			ws.Deregister(newMetric)
			if _, ok := ws.stats[newMetric]; ok {
				t.Fatalf("Failed. Deregister method not ok")
			}
		})
	}
}

type mockMetric struct {
	value string
	ch    chan string
}

func (mm *mockMetric) OnNotify() {
	mm.ch <- mm.value
}

func (mm *mockMetric) Rgx() *regexp.Regexp {
	return &regexp.Regexp{}
}

func (mm *mockMetric) Value(value string) {
	mm.value = value
}

func TestNotify(t *testing.T) {
	cases := []struct {
		testcaseName string
	}{
		{
			testcaseName: "Notify Method Functionality Test",
		},
	}

	for _, c := range cases {
		t.Run(c.testcaseName, func(t *testing.T) {
			ws := &wifiStats{stats: make(map[metric.Metric]struct{})}
			ch := make(chan string)
			ws.Register(&mockMetric{value: "test", ch: ch})
			go func() {
				select {
				case msg := <-ch:
					if msg != "test" {
						t.Fatalf("Failed. Notify method not ok")
					}
				case <-time.After(100 * time.Millisecond):
					t.Fatalf("Failed. Notify method not ok")
				}
			}()
			time.Sleep(10 * time.Millisecond)
			ws.Notify()
		})
	}
}

func TestCheckError(t *testing.T) {
	var buffer bytes.Buffer
	log.SetOutput(&buffer)
	checkError(errors.New("test"))
	if !strings.Contains(buffer.String(), "Failure:  test") {
		t.Fatalf("Failed. checkError Function not ok -%s", buffer.String())
	}
}

// func TestCollect(t *testing.T) {
// 	router := mux.NewRouter()
// 	router.HandleFunc(model.JIOFIURL, func(rw http.ResponseWriter, r *http.Request) {
// 		fmt.Println("testing")
// 	})
// 	log.Fatal(http.ListenAndServe("192.168.225.1:80", router))
// }
