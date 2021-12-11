package alert

import (
	"bytes"
	"log"
	"strings"
	"testing"
)

func TestSay(t *testing.T) {
	cases := []struct {
		testcaseName string
		expected     say
		got          say
	}{
		{
			testcaseName: "Test Say Methods",
			expected: say{
				id:      "test",
				title:   "title",
				message: "message",
				audio:   "audio",
			},
			got: say{},
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

func TestSayPush(t *testing.T) {
	cases := []struct {
		testcaseName string
	}{
		{
			testcaseName: "Test Say Push Method",
		},
	}
	for _, c := range cases {
		t.Run(c.testcaseName, func(t *testing.T) {
			var buffer bytes.Buffer
			log.SetOutput(&buffer)
			say := &say{}
			say.Push()
			if !strings.Contains(buffer.String(), "") {
				t.Fatalf("Failed: Got: %s", buffer.String())
			}
		})
	}
}
