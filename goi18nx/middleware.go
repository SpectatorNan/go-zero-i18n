package goi18nx

import (
	"golang.org/x/text/language"
	"net/http"
)

type I18nMiddleware struct {
	supportTags       []language.Tag
	localizationFiles []string
}

func NewI18nMiddleware(supportTags []language.Tag, localizationFiles []string) *I18nMiddleware {
	return &I18nMiddleware{
		supportTags:       supportTags,
		localizationFiles: localizationFiles,
	}
}

func (m *I18nMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lang := r.Header.Get(defaultLangHeaderKey)
		langTag := FetchCurrentLanguageTag(lang, m.supportTags)
		bundle := NewBundle(langTag, m.localizationFiles...)
		next(w, withRequest(r, bundle))
	}
}
