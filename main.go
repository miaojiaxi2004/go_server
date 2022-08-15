package main

import (
	"embed"
	"github.com/miaojiaxi2004/go_server/core"
	"github.com/miaojiaxi2004/go_server/global"
	"github.com/miaojiaxi2004/go_server/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

var f embed.FS

func main() {
	global.VIPER = core.Viper()          // 初始化Viper
	global.LOG = core.Zap()              // 初始化zap日志库
	global.DB = initialize.Gorm()        //初始化mysql
	global.REDIS = initialize.Redis()    //初始化redis
	global.Router = initialize.Routers() //初始化路由
	////global.Router.StaticFS("/static", http.FS(f)) // 为用户头像和文件提供静态地址 //go:embed static/*
	global.Router.Run(":8888")
}
