package decorator

import "context"

type QueryHandler[Q any, R any] interface {
	Handle(ctx context.Context, cmd Q) (R, error)
}
