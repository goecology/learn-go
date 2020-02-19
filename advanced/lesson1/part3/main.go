package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goecology/muses"
	"github.com/goecology/muses/pkg/cmd"
	musesgin "github.com/goecology/muses/pkg/server/gin"
	"net/http"
)

func main() {
	app := muses.Container(
		cmd.Register,
		musesgin.Register,
	)
	app.SetGinRouter(InitRouter)
	err := app.Run()
	if err != nil {
		panic(err)
	}
}

func InitRouter() *gin.Engine {
	r := musesgin.Caller()
	r.GET("/hello", func(context *gin.Context) {
		context.String(http.StatusOK, "hello ecology")
	})
	return r
}
