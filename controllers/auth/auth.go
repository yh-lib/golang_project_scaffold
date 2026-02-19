package auth

import (
	"golang_project_scaffold/config"
	"golang_project_scaffold/utils/jwtutils"
	"golang_project_scaffold/utils/logs"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `json:username`
	Password string `json:password`
}

func Login(r *gin.Context) {
	// 1. 获取前端传递的用户名和密码
	userInfo := UserInfo{}
	if err := r.ShouldBindJSON(&userInfo); err != nil {
		r.JSON(200, gin.H{
			"message": err.Error(),
			"status":  401,
		})
		return
	}
	logs.Debug(map[string]interface{}{"用户名": userInfo.Username, "密码": userInfo.Password}, "开始验证登录信息")
	// 验证用户名和密码是否正确
	// 认证成功
	if userInfo.Username == config.Username && userInfo.Password == config.Password {

		// 生成 JWT 的 Token
		jwtToken, err := jwtutils.GenToken(userInfo.Username)
		if err != nil {
			logs.Error(map[string]interface{}{"用户名": userInfo.Username, "错误信息": err}, "用户名密码正确,但生成toke失败.")
			r.JSON(200, gin.H{
				"status":  401,
				"message": "生成token失败:",
			})
			return
		}
		// token 正常生成，返回给前端
		logs.Info(map[string]interface{}{"用户名": userInfo.Username}, "登录成功")
		data := make(map[string]interface{})
		data["token"] = jwtToken
		r.JSON(200, gin.H{
			"status":  200,
			"message": "生成token成功",
			"data":    data,
		})
		return
	} else {
		// 认证失败,用户名密码错误
		r.JSON(200, gin.H{
			"status":  401,
			"message": "认证失败：用户名或密码错误.",
		})
	}
}

func Logout(r *gin.Context) {
	// 退出
	r.JSON(200, gin.H{
		"message": "退出成功",
		"status":  200,
	})
	logs.Debug(nil, "用户已退出.")
}
