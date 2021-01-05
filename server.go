/**
* @Author: mo tan
* @Description:
* @Date 2021/1/1 21:58
 */
package main

import (
	"gin-server/routes"
	"github.com/maotan/go-truffle/web"
)

func main() {

	/*consulConf :=yaml_config.YamlConf.ConsulConf
	registryDiscoveryClient, err := serviceregistry.NewConsulServiceRegistry(consulConf.Host,
		consulConf.Port, consulConf.Token)
	feign.Init(registryDiscoveryClient)

	ip, err := util.GetLocalIP()
	if err != nil {
		panic(err)
	}
	rand.Seed(time.Now().UnixNano())

	serverConf := yaml_config.YamlConf.ServerConf
	si, _ := cloud.NewDefaultServiceInstance(serverConf.Name, ip, serverConf.Port,
		false, map[string]string{"user": "zyn2"}, "")
	registryDiscoveryClient.Register(si)*/
	err, registryClient:= web.WebInit()
	if err != nil{
		panic(err)
	}
	err = routes.Run()
	if err != nil{
		registryClient.Deregister()
	}
}