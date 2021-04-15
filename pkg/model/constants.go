package model

type Label string

const (
	LabelBatteryQuantity Label = "lDashBatteryQuantity"
	LabelChargeStatus    Label = "lDashChargeStatus"
)

const (
	JIOFIURL string = "http://jiofi.local.html/cgi-bin/en-jio/mStatus.html"
)

type NotifierType uint8

const (
	Toast NotifierType = iota
	Notify
)
