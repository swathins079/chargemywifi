package model

type OSType uint8

const (
	Windows OSType = iota
	Linux
)

var OperatingSystem string
