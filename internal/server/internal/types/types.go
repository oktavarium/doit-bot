package types

import "context"

type BotAddedCallBack func(ctx context.Context, chatID int64, userID int64, username string) error
