package model

type OSType uint8

const (
	Windows OSType = iota
	Linux
	MacOS
)

var OperatingSystem string
