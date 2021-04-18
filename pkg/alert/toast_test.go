package alert

import (
	"bytes"
	"log"
	"strings"
	"testing"
)

func TestToast(t *testing.T) {
	cases := []struct {
		testcaseName string
		expected     toast
		got          toast
	}{
		{
			testcaseName: "Test Toast Methods",
			expected: toast{
				id:      "test",
				title:   "title",
				message: "message",
				audio:   "audio",
			},
			got: toast{},
		},
	}
	for _, c := range cases {
		t.Run(c.testcaseName, func(t *testing.T) {
			_ = c.got.ID(c.expected.id)
			_ = c.got.Title(c.expected.title)
			_ = c.got.Message(c.expected.message)
			_ = c.got.Audio(c.expected.audio)
			if c.got.id != c.expected.id ||
				c.got.title != c.expected.title ||
				c.got.message != c.expected.message ||
				c.got.audio != c.expected.audio {
				t.Fatalf("Failed: Expected: %v, Got: %v", c.expected, c.got)
			}
		})
	}
}

func TestToastPush(t *testing.T) {
	cases := []struct {
		testcaseName string
	}{
		{
			testcaseName: "Test Toast Push Method",
		},
	}
	for _, c := range cases {
		t.Run(c.testcaseName, func(t *testing.T) {
			var buffer bytes.Buffer
			log.SetOutput(&buffer)
			toast := &toast{}
			toast.Push()
			if !strings.Contains(buffer.String(), "PowerShell\": executable file not found in $PATH") {
				t.Fatalf("Failed: Got: %s", buffer.String())
			}
		})
	}
}
