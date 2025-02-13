# ejourney-auth-middleware

Middleware untuk autentikasi JWT yang dapat digunakan di berbagai service berbasis Golang dengan framework Gin.

## Instalasi
Tambahkan middleware ini ke proyek Anda dengan perintah berikut:

```sh
go get github.com/Primeskills-Web-Team/ejourney-auth-middleware
```

## Penggunaan

Import middleware di dalam proyek Gin Anda dan gunakan secara global untuk semua route:

```go
import (
    "github.com/gin-gonic/gin"
    middleware "github.com/Primeskills-Web-Team/ejourney-auth-middleware"
    "log"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    
    cfg, err := configJwt.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }
    
    // Pasang middleware JWT untuk semua route
    r.Use(middleware.EjourneyMiddlewareUserJWT(cfg.Secret))
    
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })
    
    return r
}
```

Atau, jika hanya ingin digunakan dalam grup tertentu:

```go
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

