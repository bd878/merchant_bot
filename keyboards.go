package merchant_bot

import (
	"fmt"
	"github.com/go-telegram/bot/models"
)

func TransactionsKeyboard(code LangCode, trans *models.StarTransactions) *models.InlineKeyboardMarkup {
	keyboards := make([][]models.InlineKeyboardButton, 1)
	for _, tr := range trans.Transactions {
		keyboards[0] = append(keyboards[0], models.InlineKeyboardButton{
			Text: fmt.Sprintf("%d - %d", tr.Amount, tr.Date), CallbackData: "tr_" + tr.ID,
		})
	}
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: keyboards,
	}
}