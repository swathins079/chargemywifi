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

func (n *say) ID(id string) Alert {
	n.id = id
	return n
}

func (n *say) Title(title string) Alert {
	n.title = title
	return n
}

func (n *say) Message(msg string) Alert {
	n.message = msg
	return n
}

func (n *say) Audio(audio string) Alert {
	n.audio = audio
	return n
}

func (n *say) Count(count int) Alert {
	n.count = count
	return n
}

func (n *say) Push() {
	cmd := exec.Command("say", n.message)
	cmd.Run()
}
