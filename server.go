/**
* @Author: mo tan
* @Description:
* @Date 2021/1/1 21:58
 */
package main

import (
	"fmt"
	"gin-server/routes"
	"github.com/gin-gonic/gin"
	"github.com/maotan/go-truffle/web"
	"github.com/maotan/go-truffle/yamlconf"
	log "github.com/sirupsen/logrus"
)

func main() {
	
	registryClient, err := web.ConsulInit(map[string]string{"user": "zyn2"})
	if err != nil{
		panic(err)
	}

	router := gin.New()
	web.RouterInit(router)
	web.DatabaseInit()  // db

	routes.AddUserRoutes(router)
	routes.AddPingRoutes(router)

	serverConf := yamlconf.YamlConf.ServerConf
	runHostPort := fmt.Sprintf(":%d", serverConf.Port)
	log.Info("app run...")
	err = router.Run(runHostPort)
	if err != nil{
		registryClient.Deregister()
	}
}