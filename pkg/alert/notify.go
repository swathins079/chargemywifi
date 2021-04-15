package alert

import (
	"os/exec"
)

type notify struct {
	id      string
	title   string
	message string
}

func (n *notify) ID(id string) Alert {
	n.id = id
	return n
}

func (n *notify) Title(title string) Alert {
	n.title = title
	return n
}

func (n *notify) Message(msg string) Alert {
	n.message = msg
	return n
}

func (n *notify) Push() {
	cmd := exec.Command("notify-send", "-t", "5000", n.title, n.message)
	cmd.Run()
}
