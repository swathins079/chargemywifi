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
	fmt.Sprintf("%s-%s", model.LabelChargeStatus, "Fully Charged"): New(model.Toast).Title("Charging Status").Message("Charging Status is Fully Charged, Please disconnect charger."),
	fmt.Sprintf("%s-%s", model.LabelBatteryQuantity, "100%"):       New(model.Toast).Title("Battery Level").Message("Battery Level is 100%, Please disconnect charger."),
	fmt.Sprintf("%s-%s", model.LabelBatteryQuantity, "50%"):        New(model.Toast).Title("Battery Level").Message("Battery Level is 50%."),
	fmt.Sprintf("%s-%s", model.LabelBatteryQuantity, "10%"):        New(model.Toast).Title("Battery Level").Message("Low Battery Level. Please connect charger"),
	fmt.Sprintf("%s-%s", model.LabelBatteryQuantity, "5%"):         New(model.Toast).Title("Battery About to die").Message("Battery Level critical. Please connect charger immediately"),
}

func New(notifierType model.NotifierType) Alert {
	switch notifierType {
	case model.Toast:
		return &toast{}
	case model.Notify:
		return &notify{}
	default:
	}
	return nil
}
