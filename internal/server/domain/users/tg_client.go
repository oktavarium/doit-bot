package users

import "context"

type TGClient interface {
	SendMessage(
		ctx context.Context,
		messageText string,
		chatID int64,
	) error
	SendStartupButton(
		ctx context.Context,
		buttonText string,
		messageText string,
		chatID int64,
	) error
}
