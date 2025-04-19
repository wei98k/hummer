package main

import (
	"hummer/router"
	"hummer/storage"
)

func main() {
	storage.InitMySQL() // 初始化数据库连接
	r := router.SetupRouter()
	r.Run(":8080")
}
