package alert

import (
	"fmt"

	"github.com/swathins079/chargemywifi/pkg/model"
)

type Alert interface {
	ID(id string) Alert
	Title(title string) Alert
	Message(msg string) Alert
	Audio(string) Alert
	Push()
}

var AlertPool = map[string]Alert{
	fmt.Sprintf("%d-%s-%s", model.Windows, model.LabelChargeStatus, "Fully Charged"): New(model.Toast).Title("Charging Status").Message("Charging Status is Fully Charged, Please disconnect charger.").Audio(model.Reminder),
	fmt.Sprintf("%d-%s-%s", model.Windows, model.LabelBatteryQuantity, "100%"):       New(model.Toast).Title("Battery Level").Message("Battery Level is 100%, Please disconnect charger.").Audio(model.Reminder),
	fmt.Sprintf("%d-%s-%s", model.Windows, model.LabelBatteryQuantity, "50%"):        New(model.Toast).Title("Battery Level").Message("Battery Level is 50%.").Audio(model.Reminder),
	fmt.Sprintf("%d-%s-%s", model.Windows, model.LabelBatteryQuantity, "10%"):        New(model.Toast).Title("Battery Level").Message("Low Battery Level. Please connect charger").Audio(model.Alarm),
	fmt.Sprintf("%d-%s-%s", model.Windows, model.LabelBatteryQuantity, "5%"):         New(model.Toast).Title("Battery About to die").Message("Battery Level critical. Please connect charger immediately").Audio(model.Alarm),

	fmt.Sprintf("%d-%s-%s", model.Linux, model.LabelChargeStatus, "Fully Charged"): New(model.Notify).Title("Charging Status").Message("Charging Status is Fully Charged, Please disconnect charger."),
	fmt.Sprintf("%d-%s-%s", model.Linux, model.LabelBatteryQuantity, "100%"):       New(model.Notify).Title("Battery Level").Message("Battery Level is 100%, Please disconnect charger."),
	fmt.Sprintf("%d-%s-%s", model.Linux, model.LabelBatteryQuantity, "50%"):        New(model.Notify).Title("Battery Level").Message("Battery Level is 50%."),
	fmt.Sprintf("%d-%s-%s", model.Linux, model.LabelBatteryQuantity, "10%"):        New(model.Notify).Title("Battery Level").Message("Low Battery Level. Please connect charger"),
	fmt.Sprintf("%d-%s-%s", model.Linux, model.LabelBatteryQuantity, "5%"):         New(model.Notify).Title("Battery About to die").Message("Battery Level critical. Please connect charger immediately"),
}

func New(alertType model.AlertType) Alert {
	switch alertType {
	case model.Toast:
		return &toast{}
	case model.Notify:
		return &notify{}
	}
	return nil
}
