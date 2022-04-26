package goi18nx

import (
	"github.com/BurntSushi/toml"
	i18n2 "github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

const I18nKey = "SpectatorNan/goi18nx"

func NewBundle(tag language.Tag, configs ...string) *i18n2.Bundle {
	bundle := i18n2.NewBundle(tag)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	for _, file := range configs {
		bundle.LoadMessageFile(file)
	}
	return bundle
}

func FormatText(cache Cache, message *i18n2.Message) string {
	return FormatMessage(cache, message, nil)
}

func FormatMessage(cache Cache, message *i18n2.Message, args map[string]interface{}) string {
	if localizer, ok := cache.GetLocalizer(); ok {
		return localizer.MustLocalize(&i18n2.LocalizeConfig{
			DefaultMessage: message,
			TemplateData:   args,
		})
	}

	return formatInternalMessage(message, args)
}

func formatInternalMessage(message *i18n2.Message, args map[string]interface{}) string {
	if args == nil {
		return message.Other
	}
	tpl := i18n2.NewMessageTemplate(message)
	msg, err := tpl.Execute("other", args, nil)
	if err != nil {
		panic(err)
	}
	return msg
}
