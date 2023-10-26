# goi18nx
 
## Usage
 
### Go zero api server
 
```shell
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

### Go zero rpc server
 
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