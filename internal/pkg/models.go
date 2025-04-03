package pkg

import (
	"time"
	"github.com/go-telegram/bot/models"
	"github.com/bd878/merchant_bot/internal/i18n"
)

type ChatKey struct {}
type LangKey struct {}

type Payment struct {
	*models.SuccessfulPayment
	ID uint32
	UserID int64
	Refunded bool
	CreatedAt time.Time
}

type Chat struct {
	*models.Chat
	Lang i18n.LangCode
}