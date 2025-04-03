package payments

import (
	"fmt"
	"github.com/go-telegram/bot/models"
	merchant "github.com/bd878/merchant_bot"
)

func TransactionsKeyboard(code merchant.LangCode, transactions []*merchant.Payment) *models.InlineKeyboardMarkup {
	keyboards := make([][]models.InlineKeyboardButton, 0)
	for _, tr := range transactions {
		keyboards = append(keyboards, []models.InlineKeyboardButton{
			{Text: fmt.Sprintf("%d - %d", tr.TotalAmount, tr.ID), CallbackData: fmt.Sprintf("tr:%d", tr.ID)},
		})
	}

	keyboards = append(keyboards, []models.InlineKeyboardButton{
		{Text: fmt.Sprintf("%s", code.Text("back")), CallbackData: "back"},
	})

	return &models.InlineKeyboardMarkup{
		InlineKeyboard: keyboards,
	}
}

func RefundKeyboard(code merchant.LangCode, trID uint32) *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: code.Text("refund"), CallbackData: fmt.Sprintf("refund:%d", trID)},
			},
			{
				{Text: code.Text("back"), CallbackData: "back"},
			},
		},
	}
}