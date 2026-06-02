package mqtt

import "fmt"

func SensorReadings(sensorID string) string {
	return fmt.Sprintf("sensors/%s/readings", sensorID)
}
