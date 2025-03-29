package merchant_bot_test

import (
	"testing"
	bot "github.com/bd878/merchant_bot"
)

func TestI18n(t *testing.T) {
	handlerRu := bot.CreateHandler(bot.LangRu)
	text := handlerRu("test")
	if text == "" {
		t.Error("no text for ru key \"test\", expected non empty string")
	}
	t.Log(text)

	handlerEn := bot.CreateHandler(bot.LangEn)
	text = handlerEn("test")
	if text == "" {
		t.Error("no text for en key \"test\", expected non empty string")
	}
	t.Log(text)
}