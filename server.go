/**
* @Author: mo tan
* @Description:
* @Date 2021/1/1 21:58
 */
package main

import (
	"gin-server/routes"
	"github.com/gin-gonic/gin"
	"github.com/maotan/go-truffle/web"
	log "github.com/sirupsen/logrus"
)

func main() {
	
	err, registryClient:= web.ConsulInit()
	if err != nil{
		panic(err)
	}

	router := gin.Default()
	web.RouterInit(router)
	routes.AddUserRoutes(router)
	routes.AddPingRoutes(router)

	log.Info("app run...")
	err = router.Run(":5000")
	if err != nil{
		registryClient.Deregister()
	}
}