# goi18nx
 
## Usage

### Interceptor
 
#### Go zero api server
 

```go

import "github.com/SpectatorNan/go-zero-i18n/goi18nx"

    ......
	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf, errorx.RouteMethodNotAllow())
 
    // LocalizationFiles is a slice of toml files path
	// register i18n
	server.Use(goi18nx.NewI18nMiddleware([]language.Tag{
        language.English,
        language.Chinese,
    }, []string{".../active.en.toml",".../active.zh.toml"}).Handle)
	
	......

```

#### Go zero rpc server
 
```go
import "github.com/SpectatorNan/go-zero-i18n/goi18nx"

    ......

    i18n := goi18nx.NewI18nGrpcInterceptor([]language.Tag{
            language.English,
            language.Chinese,
        }, c.LocalizationFiles)
    s.AddUnaryInterceptors(i18n.Interceptor)
	
    ......
	
```

#### Format string

```go

    // check context has i18n translator
    goi18nx.IsHasI18n(ctx)

    // context.Context
	// msgKey: localization string's key
	// defaultMsg: default string (if not found in localization file)
    goi18nx.FormatText(ctx, serr.MsgKey, serr.DefaultMsg)

    
    
```

### Support From DB
```go
// Shop is a struct from db
func (s *Shop) Name(ctx context.Context) string {
	langMap := map[language.Tag]string{
		language.English: s.NameEn,
		language.Chinese: s.NameCn,
		// more language mapping
	}
	return goi18nx.LocalizedString(ctx, s.NameCn, langMap)
}

```