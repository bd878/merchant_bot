package merchant_bot

import (
	"github.com/go-telegram/bot/models"
)

func MemberKickedMatch(update *models.Update) bool {
	if update.MyChatMember != nil {
		return (
			update.MyChatMember.NewChatMember.Type == models.ChatMemberTypeBanned ||
			update.MyChatMember.NewChatMember.Type == models.ChatMemberTypeLeft)
	}
	return false
}

func MemberRestoredMatch(update *models.Update) bool {
	if update.MyChatMember != nil {
		return (
			update.MyChatMember.NewChatMember.Type == models.ChatMemberTypeMember ||
			update.MyChatMember.NewChatMember.Type == models.ChatMemberTypeAdministrator)
	}
	return false
}

func PreCheckoutUpdateMatch(update *models.Update) bool {
	return update.PreCheckoutQuery != nil
}

func SuccessfullPaymentMatch(update *models.Update) bool {
	if update.Message != nil {
		return update.Message.SuccessfulPayment != nil
	}
	return false
}