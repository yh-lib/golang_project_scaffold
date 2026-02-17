// 项目的总入口
package main

import (
	"golang_project_scaffold/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// step1 开始加载程序配置
	// import golang_project_scaffold/config

	// step2 配置 gin
	r := gin.Default()
	r.Run(config.Port)
}
