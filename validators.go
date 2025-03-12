package telegram

import (
	"errors"
	"regexp"
	"slices"
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

func ValidateBotToken(value any) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("must be a string")
	}

	if !botTokenRe.MatchString(s) {
		return errors.New("invalid secret token format")
	}

	return nil
}

func ValidateBotWebhookSecretToken(value any) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("must be a string")
	}
	if !secretTokenRe.MatchString(s) {
		return errors.New("invalid secret token format")
	}
	return nil
}

func ValidateBotUpdateType(value any) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("must be a string")
	}
	if !slices.Contains(updateTypes, s) {
		return errors.New("invalid update type")
	}
	return nil
}

func ValidateUsername(value any) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("must be a string")
	}
	if !usernameRe.MatchString(s) {
		return errors.New("invalid username format")
	}
	return nil
}
