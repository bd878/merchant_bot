package merchant_bot_test

import (
	"testing"
	bot "github.com/bd878/merchant_bot"
)

func TestI18n(t *testing.T) {
	text := bot.LangRu.Text("test")
	if text == "" {
		t.Error("no text for ru key \"test\", expected non empty string")
	}
	t.Log(text)

	text = bot.LangEn.Text("test")
	if text == "" {
		t.Error("no text for en key \"test\", expected non empty string")
	}
	t.Log(text)

	decls := bot.LangRu.Decl("sah")
	if len(decls) == 0 {
		t.Error("no decls for ru key \"test\", expected non empty list")
	}
	t.Log(decls)

	decls = bot.LangEn.Decl("sah")
	if len(decls) == 0 {
		t.Error("no decls for en key \"test\", expected non empty list")
	}
	t.Log(decls)
}