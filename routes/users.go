package routes

import (
	"gin-server/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/ilibs/gosql/v2"
	"github.com/maotan/go-truffle/truffle"
	"github.com/maotan/go-truffle/util"
	"net/http"
)

func AddUserRoutes(router *gin.Engine) {
	users := router.Group("/v1/users")

	users.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)
		u := session.Get("user")
		c.JSON(http.StatusOK, gin.H{"user":u})
	})
	users.GET("/comments", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users comments")
	})

	// 创建用户
	users.POST("/", func(ctx *gin.Context) {
		var user model.User
		if err := ctx.BindJSON(&user); err != nil {
			panic(truffle.NewWarnError(400, "参数错误"))
		}
		user.Account = user.Mobile
		id := util.GenSnowFlakeId()
		user.Id = id
		gosql.Model(&user).Create()

		ctx.JSON(http.StatusOK, user)
	})

	// 登录
	users.POST("/login", func(ctx *gin.Context) {
		var user model.User
		if err := ctx.BindJSON(&user); err != nil {
			panic(truffle.NewWarnError(400, "参数错误"))
		}
		//Get
		userDb := &model.User{}
		gosql.Model(userDb).Where("mobile=?",user.Mobile).Get()

		session := sessions.Default(ctx)
		session.Set("user", userDb.Id)
		session.Save()
		ctx.JSON(http.StatusOK, userDb)
	})
}