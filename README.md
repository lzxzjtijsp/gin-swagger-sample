# Gin + Swaggerの導入


goを初期化
```:sh
go mod init app 
```

適当にパッケージを用意
```:sh
mkdir -p app/{controller,database,helper,middleware,model,service}
touch README.md main.go
```

ginをインストール
```:sh
go get -u github.com/gin-gonic/gin
```

ここまでで以下のような状態
```:sh
~/IdeaProjects/gin-swagger $ tree    
.
├── README.md
├── app
│   ├── controller
│   ├── database
│   ├── helper
│   ├── middleware
│   ├── model
│   └── service
├── go.mod
├── go.sum
└── main.go
```


main.goで以下のように記載する。

```go:main.go
package main

import "github.com/gin-gonic/gin"

import "net/http"

func main() {
    engine:= gin.Default()
    engine.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "hello world",
        })
    })
    engine.Run(":8080")
}

```

Goを起動
```sh
go run main.go
```

以下のコマンドをターミナル上で実行してレスポンスを確認します。
```sh
curl http://localhost:8080/
{"message":"hello world"}% 
```
hello worldが返ってきたらOK


#ここから導入

swagインストール
```sh
go install github.com/swaggo/swag/cmd/swag@latest
```

初期化
```sh
 swag init
2023/09/14 23:17:33 Generate swagger docs....
2023/09/14 23:17:33 Generate general API Info, search dir:./
2023/09/14 23:17:33 create docs.go at docs/docs.go
2023/09/14 23:17:33 create swagger.json at docs/swagger.json
2023/09/14 23:17:33 create swagger.yaml at docs/swagger.yaml
```

これで関連ファイルが作成される
```sh
~/IdeaProjects/gin-swagger $ tree
.
├── README.md
├── app
│   ├── controller
│   ├── database
│   ├── helper
│   ├── middleware
│   ├── model
│   └── service
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
└── main.go
```

swaggerをインストール
```sh
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```

コントローラーを作成し、コメントを書く
```go:controller/hello.go
package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary liveness probe
// @Schemes
// @Description do ping
// @Tags Hello World
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /hello [get]
func HelloWorld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}
```


main.goを以下のようにする
```go:main.go
package main

import (
	"app/app/controller"
	"app/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	route := SetupRoutes()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := route.Group("/api/v1")
	{
		eg := v1.Group("/")
		{
			eg.GET("/hello", controller.HelloWorld)
		}
	}
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	route.Run(":8080")
}

func SetupRoutes() *gin.Engine {
	route := gin.Default()
	route.GET("/", controller.HelloWorld)
	return route
}
```

swag initを実行する
```sh
swag init
2023/09/14 23:29:09 Generate swagger docs....
2023/09/14 23:29:09 Generate general API Info, search dir:./
2023/09/14 23:29:09 create docs.go at docs/docs.go
2023/09/14 23:29:09 create swagger.json at docs/swagger.json
2023/09/14 23:29:09 create swagger.yaml at docs/swagger.yaml
```
Goを立ち上げる
```:sh
go run main.go    
```

http://localhost:8080/api/v1/hello にアクセスして動作していることを確認

![](https://storage.googleapis.com/zenn-user-upload/96fe9b125e5b-20230914.png)


http://localhost:8080/swagger/index.html にアクセスしてswaggerが動作していることを確認
![](https://storage.googleapis.com/zenn-user-upload/90f0fa2cb813-20230914.png)

