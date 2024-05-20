package goi18nx

import (
	"context"
	i18n2 "github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func FormatText(ctx context.Context, msgId string, defaultText string) string {
	return FormatTextWithData(ctx, msgId, defaultText, nil)
}

func FormatTextWithData(ctx context.Context, msgId string, defaultText string, args map[string]interface{}) string {
	return FormatMessage(ctx, &i18n2.Message{
		ID:    msgId,
		Other: defaultText,
	}, args)
}

func FormatMessage(ctx context.Context, message *i18n2.Message, args map[string]interface{}) string {
	if localizer, ok := getLocalizer(ctx); ok {
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

func FetchCurrentLanguageFromCtx(ctx context.Context) (*language.Tag, bool) {
	v := ctx.Value(I18nCurrentLangKey)
	if l, b := v.(language.Tag); b {
		return &l, true
	}
	return nil, false
}
