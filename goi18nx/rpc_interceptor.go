package goi18nx

import (
	"context"
	i18n2 "github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
)

var rpcInterceptorDebug = false

func WithRpcInterceptorDebug(debug bool) {
	rpcInterceptorDebug = debug
}

type I18nGrpcInterceptor struct {
	supportTags       []language.Tag
	localizationFiles []string
}

func NewI18nGrpcInterceptor(supportTags []language.Tag, localizationFiles []string) *I18nGrpcInterceptor {
	if len(supportTags) == 0 {
		panic("supportTags can not be empty")
	}
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
	//lang := "en-US" //language.English
	langTag := i.supportTags[0]
	lang := langTag.String()
	if ok {
		langs := respHeader.Get(defaultLangHeaderKey)
		if len(langs) == 0 {
			//return nil, status.Error(codes.Code(defaultErrCode), "can not correct get language")
			if rpcInterceptorDebug {
				log.Printf("can not correct get language")
			}
		} else {
			lang = langs[0]
			langTag = MatchCurrentLanguageTag(lang, i.supportTags)
		}
	} else {
		if rpcInterceptorDebug {
			log.Printf("can not correct get metadata1")
		}
		//return nil, status.Error(codes.Code(defaultErrCode), "can not correct get metadata1")
	}

	//lang := langs[0]
	//langTag := MatchCurrentLanguageTag(lang, i.supportTags)
	bundle := NewBundle(langTag, i.localizationFiles...)
	localizer := i18n2.NewLocalizer(bundle, lang)
	ctx = setLocalizer(ctx, localizer)

	// Append the language metadata to the outgoing context
	ctx = metadata.AppendToOutgoingContext(ctx, defaultLangHeaderKey, lang)
	return ctx, nil
}
