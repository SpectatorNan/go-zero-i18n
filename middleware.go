package goi18nx

import (
	"context"
	i18n2 "github.com/nicksnyder/go-i18n/v2/i18n"
	"net/http"
)

type I18nMiddleware struct {
	bundle *i18n2.Bundle
}

func NewI18nMiddleware(bundle *i18n2.Bundle) *I18nMiddleware {
	return &I18nMiddleware{
		bundle: bundle,
	}
}

func (m *I18nMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, withRequest(r, m.bundle))
	}
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
