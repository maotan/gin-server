package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/maotan/go-truffle/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddUserRoutes(router *gin.Engine) {
	users := router.Group("/users")

	users.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)
		u := session.Get("user")
		c.JSON(http.StatusOK, gin.H{"user":u})
	})
	users.GET("/comments", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users comments")
	})
	users.GET("/pictures", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users pictures")
	})
	users.GET("/set", func(c *gin.Context) {
		session := sessions.Default(c)
		id := util.GenSnowFlakeId()
		session.Set("user", id)
		session.Save()
		c.JSON(http.StatusOK, "users set")
	})
}