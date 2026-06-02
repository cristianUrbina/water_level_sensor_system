package domain

import "context"

type IReadingRepository interface {
	Add(ctx context.Context, s *SensorReading) error
}
