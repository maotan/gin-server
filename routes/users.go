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
	// gin.H{} 用于直接生成json对象
	users := router.Group("/v1/users")

	users.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)
		userId := session.Get("userId")
		userDb := &model.User{}
		gosql.Model(userDb).Where("id=?", userId).Get()
		if userDb.Id == 0{
			panic(truffle.NewWarnError(40400, "不存在该用户"))
		}
		c.JSON(http.StatusOK, truffle.Success(userDb))
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

		ctx.JSON(http.StatusOK, truffle.Success(user))
	})

	// 没有前端模拟登录
	users.GET("/login", func(ctx *gin.Context) {
		mobile := ctx.Query("mobile")
		if mobile == ""{
			panic(truffle.NewWarnError(40000, "参数错误"))
		}
		//Get
		userDb := &model.User{}
		gosql.Model(userDb).Where("mobile=?", mobile).Get()
		if userDb.Id == 0{
			panic(truffle.NewWarnError(40400, "不存在该用户"))
		}

		session := sessions.Default(ctx)
		session.Set("userId", userDb.Id)
		session.Save()
		ctx.JSON(http.StatusOK, truffle.Success(userDb))
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
		if userDb.Id == 0{
			panic(truffle.NewWarnError(40400, "不存在该用户"))
		}

		session := sessions.Default(ctx)
		session.Set("userId", userDb.Id)
		session.Save()
		ctx.JSON(http.StatusOK, truffle.Success(userDb))
	})
}