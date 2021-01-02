/**
* @Author: mo tan
* @Description:
* @Date 2021/1/1 21:58
 */
package main

import (
	"gin-server/routes"
	"github.com/maotan/go-truffle/cloud"
	"github.com/maotan/go-truffle/cloud/serviceregistry"
	"github.com/maotan/go-truffle/feign"
	"github.com/maotan/go-truffle/util"
	"github.com/maotan/go-truffle/yaml_config"
	"math/rand"
	"time"
)

func main() {

	consulConf :=yaml_config.YamlConf.ConsulConf
	registryDiscoveryClient, err := serviceregistry.NewConsulServiceRegistry(consulConf.Host,
		consulConf.Port, consulConf.Token)
	feign.Init(registryDiscoveryClient)

	ip, err := util.GetLocalIP()
	if err != nil {
		panic(err)
	}
	rand.Seed(time.Now().UnixNano())

	ginConf := yaml_config.YamlConf.GinConf
	si, _ := cloud.NewDefaultServiceInstance(ginConf.Name, ip, ginConf.Port,
		false, map[string]string{"user": "zyn2"}, "")
	registryDiscoveryClient.Register(si)

	err = routes.Run()
	if err != nil{
		registryDiscoveryClient.Deregister()
	}
}