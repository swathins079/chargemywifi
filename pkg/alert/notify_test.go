package alert

import (
	"bytes"
	"log"
	"strings"
	"testing"
)

func TestNotify(t *testing.T) {
	cases := []struct {
		testcaseName string
		expected     notify
		got          notify
	}{
		{
			testcaseName: "Test Notify Methods",
			expected: notify{
				id:      "test",
				title:   "title",
				message: "message",
				audio:   "audio",
			},
			got: notify{},
		},
	}
	for _, c := range cases {
		t.Run(c.testcaseName, func(tt *testing.T) {
			_ = c.got.ID(c.expected.id)
			_ = c.got.Title(c.expected.title)
			_ = c.got.Message(c.expected.message)
			_ = c.got.Audio(c.expected.audio)
			if c.got.id != c.expected.id ||
				c.got.title != c.expected.title ||
				c.got.message != c.expected.message ||
				c.got.audio != c.expected.audio {
				tt.Fatalf("Failed: Expected: %v, Got: %v", c.expected, c.got)
			}
		})
	}
}

func TestNotifyPush(t *testing.T) {
	cases := []struct {
		testcaseName string
	}{
		{
			testcaseName: "Test Notify Push Method",
		},
	}
	for _, c := range cases {
		t.Run(c.testcaseName, func(t *testing.T) {
			var buffer bytes.Buffer
			log.SetOutput(&buffer)
			notify := &notify{}
			notify.Push()
			if !strings.Contains(buffer.String(), "") {
				t.Fatalf("Failed: Got: %s", buffer.String())
			}
		})
	}
}
