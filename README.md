# goi18nx
 
## Usage
 
```go

    ......
	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf, errorx.RouteMethodNotAllow())
 
    // LocalizationFiles is a slice of toml files path
	// register i18n
	server.Use(i18nx.NewI18nMiddleware(func(r *http.Request) language.Tag {
	    accept := r.Header.Get("Accept-Language")
	    if accept == "zh-CN" {
	        return language.Chinese
	    } else {
	        return language.English
	    }
	}, []string{".../active.en.toml",".../active.zh.toml"}...).Handle)
	
	......

```