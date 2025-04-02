package merchant_bot

import (
	"time"
	"github.com/go-telegram/bot/models"
)

type Payment struct {
	*models.SuccessfulPayment
	ID uint32
	UserID int64
	Refunded bool
	CreatedAt time.Time
}

type Chat struct {
	*models.Chat
	Lang LangCode
}