package clients

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	merchant "github.com/bd878/merchant_bot"
)

func (m Module) StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text: "ok",
	})
	if err != nil {
		m.log.Errorw("send message returns error", "error", err)
	}
}

func (m Module) MemberKickedHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	m.log.Infow("kicked", "chat_id", update.Message.Chat.ID)
}

func (m Module) MemberRestoredHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	m.log.Infow("restored", "chat_id", update.Message.Chat.ID)
}

// https://core.telegram.org/bots/payments-stars#live-checklist
func (m Module) TermsHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	chat, ok := ctx.Value(&chatKey{}).(*merchant.Chat)
	if !ok {
		m.log.Errorw("no chat key", "chat_id", update.Message.Chat.ID)
		return
	}

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text: merchant.LangRu.Text("terms"),
		ReplyMarkup: BackKeyboard(chat.Lang, chat.ID),
	})
	if err != nil {
		m.log.Errorw("failed to send terms", "error", err)
	}
}

func (m Module) ChangeLanguageHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	chat, ok := ctx.Value(&chatKey{}).(*merchant.Chat)
	if !ok {
		m.log.Errorw("no chat key", "chat_id", update.Message.Chat.ID)
		return
	}

	lang, ok := ctx.Value(&langKey{}).(merchant.LangCode)
	if !ok {
		m.log.Errorw("no lang key", "chat_id", update.Message.Chat.ID)
		return
	}

	chat.Lang = lang
	err := m.repo.Update(ctx, chat)
	if err != nil {
		m.log.Errorw("repo failed to update lang code", "chat_id", chat.ID, "error", err)
		return
	}

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text: "ok",
	})
	if err != nil {
		m.log.Errorw("failed to send message", "error", err)
	}
}