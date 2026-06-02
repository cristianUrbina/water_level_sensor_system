package mediator

import "context"

type MediatrMediator struct{}

func (m *MediatrMediator) Send(ctx context.Context, query any) error {
	return nil
}
