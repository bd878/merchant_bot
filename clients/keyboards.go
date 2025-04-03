package clients

import (
	"github.com/go-telegram/bot/models"
	merchant "github.com/bd878/merchant_bot"
)

func BackKeyboard(code merchant.LangCode) *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: code.Text("back"), CallbackData: "back"},
			},
		},
	}
}