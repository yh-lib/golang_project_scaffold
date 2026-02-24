
# 使用说明：

## 1. 日志输出规范
```go
// 使用方法
logs.Info(map[string]interface{}{"key1": "value1", "key2": "value2"}, "message...")
// 显示效果
{"file":"logs.go","func":"golang_project_scaffold/utils/logs.Info","key1":"value1","key2":"value2","level":"info","msg":"message...","time":"2026-02-25 03:13:03"}
```
## 2. 加载程序配置
```go
// 使用方法
import "golang_project_scaffold/config"
// 效果
// 1. 配置_默认_日志输出级别_"debug"
// 2. 配置_默认_程序监听端口_":8080"
// 3. 配置_默认_JWT加密的SECRET_"加密用的SECRET"
// 4. 配置_默认_JWT过期时间_"120分钟"
// 5. 配置_默认_用户名_Admin
// 6. 配置_默认_密码_Admin123
```
## 3. 路由注册功能
```go
// 使用方法
// 示例：添加路由/api/product/postAdd
// 1. 注册根路由/api: routers/routers.go
func RegistrerRouters(r *gin.Engine) {
	... ...
    product.RegisterSubRouter(apiGroup)
}
// 2. 注册子路由/api/product/postAdd: routers/product/product.go
func postAdd(productGroup *gin.RouterGroup) {
    // 将 /api/product/postAdd 的 POST 请求，转给 product.postAdd 控制器做逻辑处理。
	productGroup.POST("/postAdd", product.postAdd)
}
func RegisterSubRouter(g *gin.RouterGroup) {
	// 注册根路由
	productGroup := g.Group("product")
	// 添加功能
	postAdd(productGroup)
}
// 3. 实现路由的逻辑处理: controllers/product.go
var returnData = config.NewRetrunData()
func PostAdd(c *gin.Context) {
	returnData.Data["效果"] = "实现路由的逻辑处理"
	c.JSON(200, returnData)
}
// 4. 验证效果
// 4.1 post /api/auth/login 获取 jwttoken
{
    "username":"YH",
    "password":"E64B78FC3BC91BCBC7DC232BA8EC59E0"
}
// 4.2 post /api/product/postAdd
// Headers
Authorization
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IllIIiwiaXNzIjoiZG90YmFsbyIsInN1YiI6ImR1a3VhbiIsImV4cCI6MTc3MTk3MTE3MCwibmJmIjoxNzcxOTYzOTcwLCJpYXQiOjE3NzE5NjM5NzB9.sQ9-RNsIRmmvg4T07X1JzzOJgHrJOVss5MXEpvlYK4Y
// 4.3 效果
// 第一段是中间件返回给前端的数据
// 第二段是路由/api/product/postAdd返回给前端的数据
{
    "status": 200,
    "message": "请求成功。",
    "data": {}
}{
    "status": 200,
    "message": "",
    "data": {
        "效果": "实现路由的逻辑处理"
    }
}
```
## 4. 登录认证功能
```go
// 使用方法
// 登录 && 获取JWTtoken
post 127.0.0.1:8080/api/auth/login
{
    "username":"YH",
    "password":"E64B78FC3BC91BCBC7DC232BA8EC59E0"
}
// 登出
post 127.0.0.1:8080/api/auth/logout
// 设置默认用户名/密码
config/config.go
viper.SetDefault("USERNAME", "E3AFED0047B08059D0FADA10F400C1E5") // 默认用户名：Admin	    通过MD5加密
viper.SetDefault("PASSWORD", "E64B78FC3BC91BCBC7DC232BA8EC59E0") // 默认用户名：Admin123    通过MD5加密
// 部署时通过环境变量修改密码
USERNAME
PASSWORD
// 请求携带 JWTtoken
Authorization   获取到的JWTtoken
```