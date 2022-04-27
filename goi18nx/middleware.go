package goi18nx

import (
	i18n2 "github.com/nicksnyder/go-i18n/v2/i18n"
	"net/http"
)

const I18nKey = "SpectatorNan/goi18nx"

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
