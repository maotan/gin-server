package routes

import (
	"gin-server/model"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/ilibs/gosql/v2"
	"github.com/maotan/go-truffle/httpresult"
	"github.com/maotan/go-truffle/util"
)

func AddUserRoutes(router *gin.Engine) {
	// gin.H{} 用于直接生成json对象
	users := router.Group("/v1/users")

	users.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)
		userId := session.Get("userId")
		userDb := &model.User{}
		gosql.Model(userDb).Where("id=?", userId).Get()
		if userDb.Id == 0 {
			panic(httpresult.NewWarnError(40400, "不存在该用户"))
		}
		c.JSON(http.StatusOK, httpresult.Success(userDb))
	})

	// 网页创建用户
	users.POST("/", func(ctx *gin.Context) {
		var user model.User
		if err := ctx.BindJSON(&user); err != nil {
			panic(httpresult.NewWarnError(400, err.Error()))
		}
		user.Account = user.Mobile
		id := util.GenSnowFlakeId()
		pwd := util.GenMd5(user.Password)
		user.Id = id
		user.Password = pwd
		gosql.Model(&user).Create()

		ctx.JSON(http.StatusOK, httpresult.Success(user))
	})

	// 没有前端模拟登录
	users.GET("/login", func(ctx *gin.Context) {
		mobile := ctx.Query("mobile")
		if mobile == "" {
			panic(httpresult.NewWarnError(40000, "参数错误"))
		}
		//Get
		userDb := &model.User{}
		gosql.Model(userDb).Where("mobile=?", mobile).Get()
		if userDb.Id == 0 {
			panic(httpresult.NewWarnError(40400, "不存在该用户"))
		}

		session := sessions.Default(ctx)
		session.Set("userId", userDb.Id)
		session.Save()
		ctx.JSON(http.StatusOK, httpresult.Success(userDb))
	})

	// 登录
	users.POST("/login", func(ctx *gin.Context) {
		var user model.User
		if err := ctx.BindJSON(&user); err != nil {
			panic(httpresult.NewWarnError(40000, "参数错误"))
		}
		//Get
		userDb := &model.User{}
		gosql.Model(userDb).Where("mobile=?", user.Mobile).Get()
		if userDb.Id == 0 {
			panic(httpresult.NewWarnError(40400, "不存在该用户"))
		}

		session := sessions.Default(ctx)
		session.Set("userId", userDb.Id)
		session.Save()
		ctx.JSON(http.StatusOK, httpresult.Success(userDb))
	})
}
