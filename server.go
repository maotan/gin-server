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
	
	err, registryClient:= web.WebInit()
	if err != nil{
		panic(err)
	}
	err = routes.Run()
	if err != nil{
		registryClient.Deregister()
	}
}