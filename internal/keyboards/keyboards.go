package keyboards

import (
	"fmt"
	"github.com/go-telegram/bot/models"
	"github.com/bd878/merchant_bot/internal/i18n"
)

func BackKeyboard(code i18n.LangCode, clientID int64) *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "« " + code.Text("back"), CallbackData: fmt.Sprintf("back:%d", clientID)},
			},
		},
	}
}

func SettingsKeyboard(code i18n.LangCode, clientID int64) *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "🇺🇸 " + code.Text("en"), CallbackData: fmt.Sprintf("en:%d", clientID)},
				{Text: "🇷🇺 " + code.Text("ru"), CallbackData: fmt.Sprintf("ru:%d",clientID)},
			},
			{
				{Text: "« " + code.Text("back"), CallbackData: fmt.Sprintf("back:%d", clientID)},
			},
		},
	}
}

func StartKeyboard(code i18n.LangCode, clientID int64) *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "⚙️ " + code.Text("settings"), CallbackData: fmt.Sprintf("settings:%d", clientID)},
			},
		},
	}
}