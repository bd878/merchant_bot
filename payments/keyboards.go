package payments

import (
	"fmt"
	"github.com/go-telegram/bot/models"
	merchant "github.com/bd878/merchant_bot"
)

func TransactionsKeyboard(code merchant.LangCode, trans *models.StarTransactions) *models.InlineKeyboardMarkup {
	keyboards := make([][]models.InlineKeyboardButton, 0)
	for _, tr := range trans.Transactions {
		keyboards = append(keyboards, []models.InlineKeyboardButton{
			{Text: fmt.Sprintf("%d - %d", tr.Amount, tr.Date), CallbackData: "trans"},
		})
	}
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: keyboards,
	}
}