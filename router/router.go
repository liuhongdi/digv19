package router

import (
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv19/controller"
	"github.com/liuhongdi/digv19/global"
	"github.com/liuhongdi/digv19/pkg/result"
	"log"
	"net/http"
	"runtime/debug"
)

func Router() *gin.Engine {
	router := gin.Default()
	//处理异常
	router.NoRoute(HandleNotFound)
	router.NoMethod(HandleNotFound)
	router.Use(Recover)

	//static
	router.StaticFS("/static", http.Dir(global.StaticSetting.StaticDir))

	// 路径映射
	idc:=controller.NewIdController()
	router.GET("/id/getone", idc.GetOne);
	router.POST("/id/verify", idc.Verify);
	return router
}

//404
func HandleNotFound(c *gin.Context) {
	result.NewResult(c).Error(404,"资源未找到")
	return
}

//500
func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			result.NewResult(c).Error(500,"服务器内部错误")
		}
	}()
	//继续后续接口调用
	c.Next()
}