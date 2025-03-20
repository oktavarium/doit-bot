package planner

import "context"

type TGClient interface {
	SendMessage(
		ctx context.Context,
		messageText string,
		chatID int64,
	) error
}
