package alert

import (
	"os/exec"
)

type say struct {
	id      string
	title   string
	message string
	audio   string
	count   int
}

func (s *say) ID(id string) Alert {
	s.id = id
	return s
}

func (s *say) Title(title string) Alert {
	s.title = title
	return s
}

func (s *say) Message(msg string) Alert {
	s.message = msg
	return s
}

func (s *say) Audio(audio string) Alert {
	s.audio = audio
	return s
}

func (s *say) Count(count int) Alert {
	s.count = count
	return s
}

func (s *say) Push() {
	if s.count == 0 {
		return
	}
	cmd := exec.Command("say", s.message)
	cmd.Run()
	s.count--
}
