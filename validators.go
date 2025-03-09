package telegram

import (
	"regexp"
	"slices"

	"github.com/go-playground/validator/v10"
)

var (
	botTokenRe    = regexp.MustCompile(`^\d{8,10}:[\w-]{35}$`)
	secretTokenRe = regexp.MustCompile(`^[\w-]{1,256}$`)
	updateTypes   = []string{
		"message",
		"edited_message",
		"channel_post",
		"edited_channel_post",
		"business_connection",
		"business_message",
		"edited_business_message",
		"deleted_business_message",
		"message_reaction",
		"message_reaction_count",
		"inline_query",
		"chosen_inline_result",
		"callback_query",
		"shipping_query",
		"pre_checkout_query",
		"purchased_paid_media",
		"poll",
		"poll_answer",
		"my_chat_member",
		"chat_member",
		"chat_join_request",
		"chat_boost",
		"removed_chat_boost",
	}
	usernameRe = regexp.MustCompile(`^\w{4,32}$`)
)

func ValidateBotToken(fl validator.FieldLevel) bool {
	return botTokenRe.MatchString(fl.Field().String())
}

func ValidateBotWebhookSecretToken(fl validator.FieldLevel) bool {
	return secretTokenRe.MatchString(fl.Field().String())
}

func ValidateBotUpdateType(fl validator.FieldLevel) bool {
	return slices.Contains(updateTypes, fl.Field().String())
}

func ValidateUsername(fl validator.FieldLevel) bool {
	return usernameRe.MatchString(fl.Field().String())
}
