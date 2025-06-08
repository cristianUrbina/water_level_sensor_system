package sensormeasurement


type ISensorMeasurementRepository interface {
	AddSensoreMeasurement(*SensorMeasurement) error
}
