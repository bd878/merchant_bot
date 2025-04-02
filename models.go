package merchant_bot

import (
	"github.com/go-telegram/bot/models"
)

type Payment struct {
	*models.SuccessfulPayment
	ID uint32
	UserID int64
	Refunded bool
	CreatedAt int64
}

type Chat struct {
	*models.Chat
	Lang LangCode
}