package domain

import (
	"context"

	"github.com/google/uuid"
)

type ISensorRepository interface {
	GetByID(ctx context.Context, ID uuid.UUID) (Sensor, error)
}
