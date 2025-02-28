package handlers

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/oktavarium/doit-bot/internal/server/internal/model"
)

type Handlers struct {
	model *model.Model
}

func New(model *model.Model) *Handlers {
	return &Handlers{
		model: model,
	}
}

func (h *Handlers) DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	switch {
	case update.Message != nil:
		Message(ctx, b, update)
	case update.ChannelPost != nil:
		NotSupported(ctx, b, update)
	case update.EditedChannelPost != nil:
		NotSupported(ctx, b, update)
	case update.BusinessConnection != nil:
		NotSupported(ctx, b, update)
	case update.BusinessMessage != nil:
		NotSupported(ctx, b, update)
	case update.EditedBusinessMessage != nil:
		NotSupported(ctx, b, update)
	case update.DeletedBusinessMessages != nil:
		NotSupported(ctx, b, update)
	case update.MessageReaction != nil:
		NotSupported(ctx, b, update)
	case update.MessageReactionCount != nil:
		NotSupported(ctx, b, update)
	case update.InlineQuery != nil:
		InlineQuery(ctx, b, update)
	case update.ChosenInlineResult != nil:
		NotSupported(ctx, b, update)
	case update.CallbackQuery != nil:
		CallbackQuery(ctx, b, update)
	case update.ShippingQuery != nil:
		NotSupported(ctx, b, update)
	case update.PreCheckoutQuery != nil:
		NotSupported(ctx, b, update)
	case update.PurchasedPaidMedia != nil:
		NotSupported(ctx, b, update)
	case update.Poll != nil:
		NotSupported(ctx, b, update)
	case update.PollAnswer != nil:
		NotSupported(ctx, b, update)
	case update.MyChatMember != nil:
		h.MyChatMember(ctx, b, update.MyChatMember)
	case update.ChatMember != nil:
		ChatMember(ctx, b, update)
	case update.ChatJoinRequest != nil:
		NotSupported(ctx, b, update)
	case update.RemovedChatBoost != nil:

	}
}
