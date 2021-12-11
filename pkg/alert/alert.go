package alert

import (
	"fmt"

	"github.com/swathinsankaran/chargemywifi/pkg/model"
)

type Alert interface {
	ID(id string) Alert
	Title(title string) Alert
	Message(msg string) Alert
	Audio(string) Alert
	Count(int) Alert
	Push()
}

var AlertPool = map[string]Alert{
	// Windows OS
	fmt.Sprintf("%d-%s-%s", model.Windows, model.LabelChargeStatus, "Fully Charged"): New(model.Toast).Title("Charging Status").Message("Charging Status is Fully Charged, Please disconnect charger.").Audio(model.Reminder).Count(1),
	fmt.Sprintf("%d-%s-%s", model.Windows, model.LabelBatteryQuantity, "100%"):       New(model.Toast).Title("Battery Level").Message("Battery Level is 100%, Please disconnect charger.").Audio(model.Reminder).Count(1),
	fmt.Sprintf("%d-%s-%s", model.Windows, model.LabelBatteryQuantity, "50%"):        New(model.Toast).Title("Battery Level").Message("Battery Level is 50%.").Audio(model.Reminder).Count(1),
	fmt.Sprintf("%d-%s-%s", model.Windows, model.LabelBatteryQuantity, "10%"):        New(model.Toast).Title("Battery Level").Message("Low Battery Level. Please connect charger").Audio(model.Alarm).Count(1),
	fmt.Sprintf("%d-%s-%s", model.Windows, model.LabelBatteryQuantity, "5%"):         New(model.Toast).Title("Battery About to die").Message("Battery Level critical. Please connect charger immediately").Audio(model.Alarm).Count(10),

	// Linux OS
	fmt.Sprintf("%d-%s-%s", model.Linux, model.LabelChargeStatus, "Fully Charged"): New(model.Notify).Title("Charging Status").Message("Charging Status is Fully Charged, Please disconnect charger.").Count(1),
	fmt.Sprintf("%d-%s-%s", model.Linux, model.LabelBatteryQuantity, "100%"):       New(model.Notify).Title("Battery Level").Message("Battery Level is 100%, Please disconnect charger.").Count(1),
	fmt.Sprintf("%d-%s-%s", model.Linux, model.LabelBatteryQuantity, "50%"):        New(model.Notify).Title("Battery Level").Message("Battery Level is 50%.").Count(1),
	fmt.Sprintf("%d-%s-%s", model.Linux, model.LabelBatteryQuantity, "10%"):        New(model.Notify).Title("Battery Level").Message("Low Battery Level. Please connect charger").Count(1),
	fmt.Sprintf("%d-%s-%s", model.Linux, model.LabelBatteryQuantity, "5%"):         New(model.Notify).Title("Battery About to die").Message("Battery Level critical. Please connect charger immediately").Count(10),

	// Mac OS
	fmt.Sprintf("%d-%s-%s", model.MacOS, model.LabelChargeStatus, "Fully Charged"): New(model.Say).Title("Charging Status").Message("Charging Status is Fully Charged, Please disconnect charger.").Count(1),
	fmt.Sprintf("%d-%s-%s", model.MacOS, model.LabelBatteryQuantity, "100%"):       New(model.Say).Title("Battery Level").Message("Battery Level is 100%, Please disconnect charger.").Count(1),
	fmt.Sprintf("%d-%s-%s", model.MacOS, model.LabelBatteryQuantity, "50%"):        New(model.Say).Title("Battery Level").Message("Battery Level is 50%.").Count(1),
	fmt.Sprintf("%d-%s-%s", model.MacOS, model.LabelBatteryQuantity, "10%"):        New(model.Say).Title("Battery Level").Message("Low Battery Level. Please connect charger").Count(1),
	fmt.Sprintf("%d-%s-%s", model.MacOS, model.LabelBatteryQuantity, "5%"):         New(model.Say).Title("Battery About to die").Message("Battery Level critical. Please connect charger immediately").Count(10),
}

func New(alertType model.AlertType) Alert {
	switch alertType {
	case model.Toast:
		return &toast{}
	case model.Notify:
		return &notify{}
	case model.Say:
		return &say{}
	}
	return nil
}
