# ejourney-auth-middleware

Middleware untuk autentikasi JWT yang dapat digunakan di berbagai service berbasis Golang dengan framework Gin.

## Instalasi
Tambahkan middleware ini ke proyek Anda dengan perintah berikut:

```sh
go get github.com/Primeskills-Web-Team/ejourney-auth-middleware
```

## Penggunaan

Import middleware di dalam proyek Gin Anda:

```go
import (
    "github.com/gin-gonic/gin"
    middleware "github.com/Primeskills-Web-Team/ejourney-auth-middleware"
)

func Register(r *gin.Engine) {
    cfg, err := configJwt.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    modules := r.Group("/module")
    modules.Use(middleware.EjourneyMiddlewareUserJWT(cfg.Secret))
    {
        modules.GET("/all", moduleCtrl.GetListModule)
        modules.POST("/create", moduleCtrl.CreateModule)
        // Tambahkan endpoint lainnya
    }
}
```


