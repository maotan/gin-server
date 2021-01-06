/**
* @Author: mo tan
* @Description:
* @Date 2021/1/1 21:58
 */
package main

import (
	"gin-server/routes"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/maotan/go-truffle/logger"
	"github.com/maotan/go-truffle/truffle"
	"github.com/maotan/go-truffle/web"
	log "github.com/sirupsen/logrus"
)

func main() {
	
	err, registryClient:= web.WebInit()
	if err != nil{
		panic(err)
	}
	//err = routes.Run()
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("my-session", store))
	router.Use(logger.LogerMiddleware())
	router.Use(truffle.Recover)
	// --------router ---------
	router.GET("/actuator/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	routes.AddUserRoutes(router)
	routes.AddPingRoutes(router)
	// --------router end------


	log.Info("app run...")
	err = router.Run(":5000")
	if err != nil{
		registryClient.Deregister()
	}
}