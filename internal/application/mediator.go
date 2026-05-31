package application

import "context"

type Mediator interface {
	Send(ctx context.Context, query any) error
}
