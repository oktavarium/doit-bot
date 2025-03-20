package command

import "context"

type CommandHandler[C any] interface {
	Handle(ctx context.Context, cmd C) error
}

type CommandHandlerHint[C any, R any] interface {
	Handle(ctx context.Context, cmd C) (R, error)
}
