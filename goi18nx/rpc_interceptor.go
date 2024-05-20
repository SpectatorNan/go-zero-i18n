package goi18nx

import (
	"context"
	i18n2 "github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type I18nGrpcInterceptor struct {
	supportTags       []language.Tag
	localizationFiles []string
}

func NewI18nGrpcInterceptor(supportTags []language.Tag, localizationFiles []string) *I18nGrpcInterceptor {
	return &I18nGrpcInterceptor{
		supportTags:       supportTags,
		localizationFiles: localizationFiles,
	}
}

func (i *I18nGrpcInterceptor) Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	ctx, err = i.saveLocalize(ctx)
	if err != nil {
		return nil, err
	}
	return handler(ctx, req)
}

func (i *I18nGrpcInterceptor) saveLocalize(ctx context.Context) (context.Context, error) {
	respHeader, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Code(defaultErrCode), "can not correct get metadata1")
	}
	langs := respHeader.Get(defaultLangHeaderKey)
	if len(langs) == 0 {
		return nil, status.Error(codes.Code(defaultErrCode), "can not correct get language")
	}
	lang := langs[0]
	langTag := MatchCurrentLanguageTag(lang, i.supportTags)
	bundle := NewBundle(langTag, i.localizationFiles...)
	localizer := i18n2.NewLocalizer(bundle, lang)
	return setLocalizer(ctx, localizer), nil
}
