package core

import (
	"fmt"
	"go-web/global"
	"go-web/initialize"
	"go.uber.org/zap"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

//endless 是一个可以用于重新加载和优雅关闭HTTP服务器的库。它可以在运行时更新服务器代码而无需停止正在运行的HTTP服务器。
//这使得服务器能够在生产环境下无缝地进行更新和维护，同时不影响当前正在运行的请求和连接。
//使用 endless，可以在代码修改后，通过发送信号量通知服务器进行重载，新的代码会被加载并运行，旧的连接会继续服务，新的连接将使用新的代码进行处理。
//当需要关闭服务器时，endless 会等待所有当前处理的请求完成后再关闭服务器，这样可以确保所有请求都能得到处理，避免数据丢失和用户体验下降。
//在 Gin 中使用 endless 可以提高服务器的可靠性和稳定性，同时也能提高开发效率，减少服务器维护和更新的停机时间。

// 配置解析到了全局变量中，使用到服务启动逻辑
type server interface {
	ListenAndServe() error
}

// RunServer gin 启动！
func RunServer() {
	// 初始化redis服务
	initialize.Redis()
	// 初始化路由
	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.WEB_CONFIG.App.Port)
	s := initServer(address, Router)

	global.WEB_LOG.Info("server run success on ", zap.String("address", address))

	// 保证文本顺序输出
	time.Sleep(10 * time.Microsecond)

	global.WEB_LOG.Error(s.ListenAndServe().Error())
}

func initServer(address string, router *gin.Engine) server {
	// 使用endless库创建一个HTTP服务器，其中address是服务器的监听地址（如:8080），router是HTTP请求路由器。
	s := endless.NewServer(address, router)

	// 设置HTTP请求头的读取超时时间为20秒，如果在20秒内未读取到请求头，则会返回一个超时错误。
	s.ReadHeaderTimeout = 20 * time.Second

	// 设置HTTP响应体的写入超时时间为20秒，如果在20秒内未将响应体写入完成，则会返回一个超时错误。
	s.WriteTimeout = 20 * time.Second

	// 设置HTTP请求头的最大字节数为1MB。如果请求头超过1MB，则会返回一个错误。
	s.MaxHeaderBytes = 1 << 20

	return s
}
