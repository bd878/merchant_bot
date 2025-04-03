package clients

import (
	"fmt"
	"github.com/go-telegram/bot/models"
	merchant "github.com/bd878/merchant_bot"
)

func BackKeyboard(code merchant.LangCode, clientID int64) *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "< " + code.Text("back"), CallbackData: fmt.Sprintf("back:%d", clientID)},
			},
		},
	}
}

func SettingsKeyboard(code merchant.LangCode, clientID int64) *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "🇺🇸 " + code.Text("en"), CallbackData: fmt.Sprintf("en:%d", clientID)},
				{Text: "🇷🇺 " + code.Text("ru"), CallbackData: fmt.Sprintf("ru:%d",clientID)},
			},
			{
				{Text: "< " + code.Text("back"), CallbackData: fmt.Sprintf("back:%d", clientID)},
			},
		},
	}
}