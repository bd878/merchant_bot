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
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: keyboards,
	}
}