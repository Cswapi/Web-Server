package main

import (
	"os"

	"github.com/Cswapi/Web-Server/service"
	flag "github.com/spf13/pflag"
)

// 设置默认端口8080
const (
	PORT string = "8080"
)

func main() {
	// 如果没有监听到其他端口，则使用默认端口8080
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	// 端口号解析，用户使用-p设置端口号
	pPort := flag.StringP("port", "p", PORT, "PORT for http to listen")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	// 创建一个server结构对象
	server := service.NewServer()
	// 启动服务器监听端口
	server.Run(":" + port)
}
