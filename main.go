package main

import (
	"hummer/config"
	"hummer/router"
	"hummer/storage"
)

func main() {
	config.InitConfig()
	storage.InitMySQL() // 初始化数据库连接
	r := router.SetupRouter()
	r.Run(":8101")
}
