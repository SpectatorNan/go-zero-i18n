package goi18nx

import i18n2 "github.com/nicksnyder/go-i18n/v2/i18n"

type Cache interface {
	GetLocalizer() (*i18n2.Localizer, bool)
	SetLocalizer(l *i18n2.Localizer)
}
