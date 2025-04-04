package history

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

	"github.com/bd878/merchant_bot/internal/keyboards"
	"github.com/bd878/merchant_bot/internal/logger"
	"github.com/bd878/merchant_bot/internal/chats"
	"github.com/bd878/merchant_bot/internal/i18n"
)

type History struct {
	chats *chats.Chats
}

func NewHistory(chats *chats.Chats) *History {
	return &History{chats}
}

func (h History) StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message != nil {
		chat, ok := h.chats.Get(update.Message.Chat.ID)
		var lang i18n.LangCode
		if ok {
			lang = chat.Lang
		} else {
			lang = i18n.LangEn
		}

		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chat.ID,
			Text: lang.Text("start_text"),
			ReplyMarkup: keyboards.StartKeyboard(lang, chat.ID),
		})
		if err != nil {
			logger.Log.Errorw("failed to send message", "error", err)
			return
		}
	} else {
		logger.Log.Warnln("update.message is nil")
	}
}

func (h History) SettingsCallbackHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.CallbackQuery != nil {
		ok, err := b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})
		if err != nil {
			logger.Log.Errorw("failed to answer callback query", "error", err)
			return
		}
		if !ok {
			logger.Log.Errorln("failed to answer callback query")
			return
		}
		chat, ok := h.chats.Get(update.CallbackQuery.From.ID)
		if !ok {
			logger.Log.Errorw("cannot find chat", "chat_id", chat.ID)
			return
		}

		_, err = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chat.ID,
			Text: chat.Lang.Text("select_lang"),
			ReplyMarkup: keyboards.SettingsKeyboard(chat.Lang, chat.ID),
		})
		if err != nil {
			logger.Log.Errorw("failed to send message", "error", err)
		}
	} else {
		logger.Log.Warnln("update.callbackQuery is nil")
	}
}

func (h History) SettingsHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message != nil {
		chat, ok := h.chats.Get(update.Message.Chat.ID)
		if !ok {
			logger.Log.Errorw("cannot find chat", "chat_id", chat.ID)
			return
		}

		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chat.ID,
			Text: chat.Lang.Text("settings"),
			ReplyMarkup: keyboards.SettingsKeyboard(chat.Lang, chat.ID),
		})
		if err != nil {
			logger.Log.Errorw("failed to send message", "error", err)
		}
	} else {
		logger.Log.Warnln("update.message is nil")
	}
}

// https://core.telegram.org/bots/payments-stars#live-checklist
func (h History) TermsHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message != nil {
		chat, ok := h.chats.Get(update.Message.Chat.ID)
		if !ok {
			logger.Log.Errorw("cannot find chat", "chat_id", chat.ID)
			return
		}

		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text: chat.Lang.Text("terms"),
			ReplyMarkup: keyboards.BackKeyboard(chat.Lang, chat.ID),
		})
		if err != nil {
			logger.Log.Errorw("failed to send terms", "error", err)
		}
	} else {
		logger.Log.Warnln("udpate.message is nil")
	}
}

func (h History) StepBackHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.CallbackQuery != nil {
		_, err := b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})
		if err != nil {
			logger.Log.Errorw("failed to answer callback query", "error", err)
			return
		}

		chat, ok := h.chats.Get(update.CallbackQuery.From.ID)
		if !ok {
			logger.Log.Errorw("cannot find chat", "chat_id", chat.ID)
			return
		}
		_, err = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chat.ID,
			Text: fmt.Sprintln(chat.Lang.Text("start_text")),
			ReplyMarkup: keyboards.StartKeyboard(chat.Lang, chat.ID),
		})
		if err != nil {
			logger.Log.Errorw("cannot send message", "chat_id", chat.ID, "error", err)
			return
		}
	} else {
		logger.Log.Warnln("update.callbackQuery is nil")
	}
}
