package alert

import (
	"reflect"
	"testing"

	"github.com/swathinsankaran/chargemywifi/pkg/model"
)

func TestNew(t *testing.T) {
	cases := []struct {
		testcaseName string
		expected     Alert
		got          Alert
	}{
		{
			testcaseName: "Toast Object creation",
			expected:     &toast{},
			got:          New(model.Toast),
		},
		{
			testcaseName: "Notify Object creation",
			expected:     &notify{},
			got:          New(model.Notify),
		},
		{
			testcaseName: "Invalid Object creation",
			expected:     nil,
			got:          New(model.AlertType(255)),
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
