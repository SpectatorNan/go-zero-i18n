package goi18nx

import (
	"context"
	i18n2 "github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"google.golang.org/grpc/metadata"
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

func withRequest(r *http.Request, currentLang language.Tag, bundle *i18n2.Bundle) *http.Request {

	accept := r.Header.Get(defaultLangHeaderKey)
	localizer := i18n2.NewLocalizer(bundle, accept)
	ctx := setLocalizer(r.Context(), localizer)
	ctx = setCurrentLang(ctx, currentLang)
	ctx = metadata.AppendToOutgoingContext(ctx, defaultLangHeaderKey, accept)
	return r.WithContext(ctx)
}

func setCurrentLang(ctx context.Context, currentLang language.Tag) context.Context {
	return context.WithValue(ctx, I18nCurrentLangKey, currentLang)
}

func setLocalizer(ctx context.Context, l *i18n2.Localizer) context.Context {
	return context.WithValue(ctx, I18nKey, l)
}

func IsHasI18n(ctx context.Context) bool {
	v := ctx.Value(I18nKey)
	if v != nil {
		return true
	}
	return false
}

//func isHasI18n(ctx context.Context) bool {
//	if use, exist := ctx.Value(isUseI18n).(bool); exist {
//		return use
//	}
//	return false
//}
//
//func setHasI18n(ctx context.Context, use bool) context.Context {
//	return context.WithValue(ctx, isUseI18n, use)
//}
