package i18n_test

import (
	"testing"
	"github.com/bd878/merchant_bot/internal/i18n"
)

func TestI18n(t *testing.T) {
	text := i18n.LangRu.Text("test")
	if text == "" {
		t.Error("no text for ru key \"test\", expected non empty string")
	}
	t.Log(text)

	text = i18n.LangEn.Text("test")
	if text == "" {
		t.Error("no text for en key \"test\", expected non empty string")
	}
	t.Log(text)

	decls := i18n.LangRu.Decl("sah")
	if len(decls) == 0 {
		t.Error("no decls for ru key \"test\", expected non empty list")
	}
	t.Log(decls)

	decls = i18n.LangEn.Decl("sah")
	if len(decls) == 0 {
		t.Error("no decls for en key \"test\", expected non empty list")
	}
	t.Log(decls)
}