package payments

import (
	"fmt"
	"github.com/go-telegram/bot/models"
	"github.com/bd878/merchant_bot/internal/i18n"
	"github.com/bd878/merchant_bot/internal/pkg"
)

func TransactionsKeyboard(code i18n.LangCode, transactions []*pkg.Payment, clientID int64) *models.InlineKeyboardMarkup {
	keyboards := make([][]models.InlineKeyboardButton, 0)
	for _, tr := range transactions {
		keyboards = append(keyboards, []models.InlineKeyboardButton{
			{Text: fmt.Sprintf("%d - %d", tr.TotalAmount, tr.ID), CallbackData: fmt.Sprintf("tr:%d", tr.ID)},
		})
	}

	keyboards = append(keyboards, []models.InlineKeyboardButton{
		{Text: "« " + code.Text("back"), CallbackData: fmt.Sprintf("back:%d", clientID)},
	})

	return &models.InlineKeyboardMarkup{
		InlineKeyboard: keyboards,
	}
}

func RefundKeyboard(code i18n.LangCode, trID uint32, clientID int64) *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: code.Text("refund"), CallbackData: fmt.Sprintf("refund:%d", trID)},
			},
			{
				{Text: "« " + code.Text("back"), CallbackData: fmt.Sprintf("back:%d", clientID)},
			},
		},
	}
}