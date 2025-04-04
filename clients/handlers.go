package clients

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/bd878/merchant_bot/internal/pkg"
	"github.com/bd878/merchant_bot/internal/i18n"
)

func (m Module) MemberKickedHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	m.app.Log().Infow("kicked", "chat_id", update.Message.Chat.ID)
}

func (m Module) MemberRestoredHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	m.app.Log().Infow("restored", "chat_id", update.Message.Chat.ID)
}

func (m Module) ChangeLanguageHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	chat, ok := ctx.Value(&pkg.ChatKey{}).(*pkg.Chat)
	if !ok {
		m.app.Log().Errorw("no chat key", "chat_id", update.CallbackQuery.From.ID)
		return
	}

	lang, ok := ctx.Value(&pkg.LangKey{}).(i18n.LangCode)
	if !ok {
		m.app.Log().Errorw("no lang key", "chat_id", update.CallbackQuery.From.ID)
		return
	}

	chat.Lang = lang
	err := m.repo.Update(ctx, chat)
	if err != nil {
		m.app.Log().Errorw("repo failed to update lang code", "chat_id", chat.ID, "error", err)
		return
	}

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chat.ID,
		Text: "ok",
	})
	if err != nil {
		m.app.Log().Errorw("failed to send message", "error", err)
	}
}