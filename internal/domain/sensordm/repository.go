package sensordm

import "github.com/google/uuid"

type ISensorRepository interface {
	GetById(uuid.UUID) (*Sensor, error)
}
