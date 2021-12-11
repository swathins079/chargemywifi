package metric

import (
	"bytes"
	"log"
	"reflect"
	"regexp"
	"strings"
	"sync"
	"testing"

	"github.com/swathinsankaran/chargemywifi/pkg/model"
)

func TestNew(t *testing.T) {
	cases := []struct {
		testcaseName string
		expected     Metric
		got          Metric
	}{
		{
			testcaseName: "Battery Quantity Metric Object creation",
			expected:     &metric{},
			got:          New(model.LabelBatteryQuantity),
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

func TestMetric(t *testing.T) {
	cases := []struct {
		testcaseName string
		expected     metric
		got          Metric
	}{
		{
			testcaseName: "Test Metric Methods",
			expected: metric{
				label: "label",
				value: sync.Map{},
				rgx:   regexp.MustCompile("<label id=\"label\">(.*?)</label>"),
			},
			got: New("label"),
		},
	}

	for _, c := range cases {
		t.Run(c.testcaseName, func(t *testing.T) {
			c.expected.value.Store(c.expected.label, "test")
			c.got.Value("test")
			expectedValue, ok := c.expected.value.Load(c.expected.label)
			if !ok {
				t.Fatalf("Failed. Value not found")
			}
			if expectedValue.(string) != "test" || c.expected.rgx.String() != c.got.Rgx().String() {
				t.Fatalf("Failed. Expected: %v, Got: %v", c.expected.rgx, c.got.Rgx())
			}
		})
	}
}

func TestOnNotify(t *testing.T) {
	cases := []struct {
		testcaseName string
		m            *metric
		expected     string
		storeValue   bool
	}{
		{
			testcaseName: "Test OnNotify Method No Key",
			m:            &metric{value: sync.Map{}, label: ""},
			expected:     "Failed to get value",
			storeValue:   false,
		},
		{
			testcaseName: "Test OnNotify Method Key Present",
			m:            &metric{value: sync.Map{}, label: "test2"},
			expected:     "Got a Value: 5%",
			storeValue:   true,
		},
	}
	for _, c := range cases {
		t.Run(c.testcaseName, func(t *testing.T) {
			var buffer bytes.Buffer
			log.SetOutput(&buffer)
			if c.storeValue {
				c.m.value.Store(c.m.label, "5%")
			}
			c.m.OnNotify()
			if !strings.Contains(buffer.String(), c.expected) {
				t.Fatalf("Failed: Got: %s", buffer.String())
			}
		})
	}
}
