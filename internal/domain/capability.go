package domain

const (
	CapabilityDistance    Capability = "distance"
	CapabilityPressure    Capability = "pressure"
	CapabilityTemperature Capability = "temperature"
	CapabilityHumidity    Capability = "humidity"
	CapabilityVoltage     Capability = "voltage"
)

func ParseCapability(
	value string,
) (
	Capability,
	error,
) {
	switch value {

	case string(CapabilityDistance):
		return CapabilityDistance, nil

	case string(CapabilityPressure):
		return CapabilityPressure, nil

	case string(CapabilityTemperature):
		return CapabilityTemperature, nil

	default:
		return "", ErrInvalidCapability
	}
}
