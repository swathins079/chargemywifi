package alert

import (
	"log"

	toastlib "gopkg.in/toast.v1"
)

type toast struct {
	id      string
	title   string
	message string
}

func (t *toast) ID(id string) Alert {
	t.id = id
	return t
}

func (t *toast) Title(title string) Alert {
	t.title = title
	return t
}

func (t *toast) Message(msg string) Alert {
	t.message = msg
	return t
}

func (t *toast) Push() {
	notification := toastlib.Notification{
		AppID:   "{1AC14E77-02E7-4E5D-B744-2EB1AE5198B7}\\WindowsPowerShell\\v1.0\\powershell.exe", // hard coding for testing
		Title:   t.title,
		Message: t.message,
		Audio:   toastlib.Default,
		Actions: []toastlib.Action{
			{"protocol", "Cancel", ""},
		},
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}
