package application

import "errors"

var ErrRecordNotFound = errors.New("record not found")
var ErrSensorNotFound = errors.New("sensor not found")
var ErrInvalidEntity = errors.New("sensor not found")
var ErrInvalidUUID = errors.New("invalid uuid")
