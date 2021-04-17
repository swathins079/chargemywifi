package model

type Label string

const (
	LabelBatteryQuantity Label = "lDashBatteryQuantity"
	LabelChargeStatus    Label = "lDashChargeStatus"
)

const (
	JIOFIURL string = "http://jiofi.local.html/cgi-bin/en-jio/mStatus.html"
)

type AlertType uint8

const (
	Toast AlertType = iota
	Notify
)

const (
	Reminder = "reminder"
	Alarm    = "loopingalarm"
)
