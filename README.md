# goi18nx
 
## Usage
 
```go

import "github.com/SpectatorNan/go-zero-i18n/goi18nx"

    ......
	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf, errorx.RouteMethodNotAllow())
 
    // LocalizationFiles is a slice of toml files path
	// register i18n
	server.Use(goi18nx.NewI18nMiddleware(func(r *http.Request) language.Tag {
        tags := []language.Tag{
            language.English,
            language.Chinese,
        }
        accept := r.Header.Get("Accept-Language")
        return localizationx.FetchCurrentLanguageTag(accept, tags)
	}, []string{".../active.en.toml",".../active.zh.toml"}...).Handle)
	
	......

```