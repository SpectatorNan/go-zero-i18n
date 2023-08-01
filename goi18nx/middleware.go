package goi18nx

import (
	"golang.org/x/text/language"
	"net/http"
)

type I18nMiddleware struct {
	fetchTag func(r *http.Request) language.Tag
	configs  []string
}

func NewI18nMiddleware(fetchTag func(r *http.Request) language.Tag, configs ...string) *I18nMiddleware {
	return &I18nMiddleware{
		fetchTag: fetchTag,
		configs:  configs,
	}
}

func (m *I18nMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tag := m.fetchTag(r)
		bundle := NewBundle(tag, m.configs...)
		next(w, withRequest(r, bundle))
	}
}
