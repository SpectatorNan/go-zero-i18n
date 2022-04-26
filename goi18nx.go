package goi18nx

import (
	"context"
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

func FormatText(ctx context.Context, message *i18n2.Message) string {
	return FormatMessage(ctx, message, nil)
}

func FormatMessage(ctx context.Context, message *i18n2.Message, args map[string]interface{}) string {
	if localizer, ok := GetLocalizer(ctx); ok {
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

func GetLocalizer(ctx context.Context) (*i18n2.Localizer, bool) {
	v := ctx.Value(I18nKey)
	if l, b := v.(*i18n2.Localizer); b {
		return l, true
	}
	return nil, false
}
