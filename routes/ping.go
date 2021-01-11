package routes

import (
	"errors"
	"gin-server/domain"
	"github.com/gin-gonic/gin"
	"github.com/maotan/go-truffle/truffle"
	"net/http"
)
type Person struct {
	Name  string
	Age     int
	Email string
}

func test() (int, error) {
	return 8, errors.New("678899")
}

func AddPingRoutes(router *gin.Engine) {
	ping := router.Group("/v1/ping")
	ping.POST("", func(ctx *gin.Context) {
		var pingDo domain.PingDo
		if err := ctx.BindJSON(&pingDo); err != nil {
			panic(truffle.NewWarnError(40000, "参数错误"))
		}
		ctx.JSON(http.StatusCreated, truffle.Success(pingDo))
	})

	ping.GET("", func(c *gin.Context) {
		//panic(truffle.NewWarnError(500,"12345"))
		test()
		var p Person
		p.Name = "123"
		p.Age = 3
		base := truffle.Success(p)
		c.JSON(http.StatusCreated, base)
	})
}