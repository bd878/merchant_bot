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
