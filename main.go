package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
	"go-gin-sample/handler"
	//"google.golang.org/appengine"
	"net/http"
)

func main() {
	router := gin.Default()
	http.Handle("/", router)

	s := rpc.NewServer()
	s.RegisterCodec(json2.NewCodec(), "application/json")
	_ = s.RegisterService(&handler.App{}, "")
	r := router.Group("/")
	route(r, s)

	// GAEで実行
	//appengine.Main()

	// ローカル実行
	_ = http.ListenAndServe(":8080", nil)
}

func route(r *gin.RouterGroup, s *rpc.Server) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "this is app")
	})
	r.POST("/jsonrpc", func(g *gin.Context) {
		s.ServeHTTP(g.Writer, g.Request)
	})
}
