// 项目的总入口
package main

import (
	"golang_project_scaffold/config"
	"golang_project_scaffold/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 加载程序的配置
	// 2. 配置gin
	r := gin.Default()
	// // 测试jwt生成token
	// ss, _ := jwtutils.GenToken("liyaohui")
	// fmt.Println("token值为:", ss)
	// // 测试jwt解析token
	// c1, _ := jwtutils.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImxpeWFvaHVpIiwiaXNzIjoiZG90YmFsbyIsInN1YiI6ImR1a3VhbiIsImV4cCI6MTc3MTQyODE4OSwibmJmIjoxNzcxNDIwOTg5LCJpYXQiOjE3NzE0MjA5ODl9.JJWLjWDNf_Zyw-ok7eM8yHtP35lF1FVq0ti-VdB8STM")
	// fmt.Println(c1)
	// // 3. 查看日志输出效果
	// testMap := map[string]interface{}{
	// 	"字段1": "值1",
	// 	"字段2": []int{1, 2, 3, 4, 5},
	// }
	// logs.Warning(testMap, "确认下效果")
	// 9. 启动程序
	routers.RegistrerRouters(r)
	r.Run(config.Port)
}
