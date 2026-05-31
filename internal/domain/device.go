package domain

import "github.com/google/uuid"

type Device struct {
	ID       uuid.UUID
	Name     string
	Metadata DeviceMetadata
	Sensors  []Sensor
}

type DeviceMetadata struct {
	Location        string
	HardwareModel   string
	FirmwareVersion string
	SerialNumber    string
}
