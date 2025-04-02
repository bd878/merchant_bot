package payments

import (
	"fmt"
	"github.com/go-telegram/bot/models"
	merchantModels "github.com/bd878/merchant_bot/models"
	merchant "github.com/bd878/merchant_bot"
)

func TransactionsKeyboard(code merchant.LangCode, transactions []*merchantModels.Payment) *models.InlineKeyboardMarkup {
	keyboards := make([][]models.InlineKeyboardButton, 0)
	for _, tr := range transactions {
		keyboards = append(keyboards, []models.InlineKeyboardButton{
			{Text: fmt.Sprintf("%d - %d", tr.TotalAmount, tr.ID), CallbackData: fmt.Sprintf("tr_%d", tr.ID)},
		})
	}
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: keyboards,
	}
}