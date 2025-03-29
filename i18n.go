package merchant_bot

import (
	_ "embed"
	"reflect"
	"encoding/json"
)

type langs struct {
	En map[string]string  `json:"en,omitempty"`
	Ru map[string]string  `json:"ru,omitempty"`
}

//go:embed texts.json
var textsFile []byte
var texts langs

func init() {
	if err := json.Unmarshal(textsFile, &texts); err != nil {
		panic(err)
	}
}

type TextHandler func(key string) string
func(h TextHandler) Text(key string) string {
	return h(key)
}

type LangCode string

const (
	LangRu LangCode = "Ru"
	LangEn = "En"
)

func (code LangCode) String() string {
	return string(code)
}

func CreateHandler(code LangCode) TextHandler {
	return TextHandler(func(key string) string {
		value := reflect.ValueOf(texts)
		fieldValue := value.FieldByName(code.String())
		if !fieldValue.IsValid() {
			log.Errorw("field value is invalid", "value", fieldValue.String())
			return ""
		}

		translations, ok := fieldValue.Interface().(map[string]string)
		if !ok {
			log.Error("not ok")
			return ""
		}

		text, ok := translations[key]
		if !ok {
			log.Errorln("cannot convert field value to map[string]string")
			return ""
		}

		return text
	})
}