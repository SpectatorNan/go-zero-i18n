package goi18nx

import (
	"context"
	i18n2 "github.com/nicksnyder/go-i18n/v2/i18n"
	"net/http"
)

/*
type Cache interface {
	GetLocalizer() (*i18n2.Localizer, bool)
	SetLocalizer(l *i18n2.Localizer)
}
*/

func getLocalizer(ctx context.Context) (*i18n2.Localizer, bool) {
	v := ctx.Value(I18nKey)
	if l, b := v.(*i18n2.Localizer); b {
		return l, true
	}
	return nil, false
}

func withRequest(r *http.Request, bundle *i18n2.Bundle) *http.Request {

	lang := r.FormValue("lang")
	accept := r.Header.Get("Accept-Language")
	localizer := i18n2.NewLocalizer(bundle, lang, accept)
	return r.WithContext(setLocalizer(r.Context(), localizer))
}

func setLocalizer(ctx context.Context, l *i18n2.Localizer) context.Context {
	return context.WithValue(ctx, I18nKey, l)
}
