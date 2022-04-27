package goi18nx

import (
	"github.com/BurntSushi/toml"
	i18n2 "github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func NewBundle(tag language.Tag, configs ...string) *i18n2.Bundle {
	bundle := i18n2.NewBundle(tag)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	for _, file := range configs {
		bundle.LoadMessageFile(file)
	}
	return bundle
}
