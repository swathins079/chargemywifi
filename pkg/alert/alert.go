package alert

import (
	"fmt"

	"github.com/swathins079/chargemywifi/pkg/model"
)

type Alert interface {
	ID(id string) Alert
	Title(title string) Alert
	Message(msg string) Alert
	Push()
}

var AlertPool = map[string]Alert{
	fmt.Sprintf("%s-%s", model.LabelChargeStatus, "Fully Charged"): &toast{title: "Charging Status", message: "Charging Status is Fully Charged, Please disconnect charger."},
	fmt.Sprintf("%s-%s", model.LabelBatteryQuantity, "100%"):       &toast{title: "Battery Level", message: "Battery Level is 100%, Please disconnect charger."},
	fmt.Sprintf("%s-%s", model.LabelBatteryQuantity, "50%"):        &toast{title: "Battery Level", message: "Battery Level is 50%."},
	fmt.Sprintf("%s-%s", model.LabelBatteryQuantity, "10%"):        &toast{title: "Battery Level", message: "Low Battery Level. Please connect charger"},
	fmt.Sprintf("%s-%s", model.LabelBatteryQuantity, "5%"):         &toast{title: "Battery About to die", message: "Battery Level critical. Please connect charger immediately"},
}

func New(notifierType string) Alert {
	switch notifierType {
	default:
		return &toast{}
	}
}
