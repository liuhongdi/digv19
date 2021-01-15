package main

import (
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv19/global"
	"github.com/liuhongdi/digv19/router"
	"log"
)

//init
func init() {
	//setting
	err := global.SetupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
    //redis
	err = global.SetupRedisDb()
	if err != nil {
		log.Fatalf("init.SetupRedisDb err: %v", err)
	}
}

func main() {
	//设置运行模式
	gin.SetMode(global.ServerSetting.RunMode)
	//引入路由
	r := router.Router()
	//run
	r.Run(":"+global.ServerSetting.HttpPort)
}




