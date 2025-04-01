package models

import "github.com/go-telegram/bot/models"

type Payment struct {
	*models.SuccessfulPayment
	UserID int64
	Refunded bool
}