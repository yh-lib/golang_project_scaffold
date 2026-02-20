package auth

import (
	"golang_project_scaffold/config"
	"golang_project_scaffold/utils/jwtutils"
	"golang_project_scaffold/utils/logs"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var returnData = config.NewRetrunData()

func Login(r *gin.Context) {
	// 1. 获取前端传递的用户名和密码
	userInfo := UserInfo{}
	if err := r.ShouldBindJSON(&userInfo); err != nil {
		returnData.Status = 401
		returnData.Message = err.Error()
		r.JSON(200, returnData)
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
			returnData.Status = 401
			returnData.Message = "生成token失败:"
			r.JSON(200, returnData)
			return
		}
		// token 正常生成，返回给前端
		logs.Info(map[string]interface{}{"用户名": userInfo.Username}, "登录成功")
		returnData.Status = 200
		returnData.Message = "生成token成功"
		returnData.Data["token"] = jwtToken
		r.JSON(200, returnData)
		return
	} else {
		// 认证失败,用户名密码错误
		returnData.Status = 401
		returnData.Message = "认证失败：用户名或密码错误."
		r.JSON(200, returnData)
	}
}

func Logout(r *gin.Context) {
	// 退出
	returnData.Status = 200
	returnData.Message = "退出成功"
	r.JSON(200, returnData)
	logs.Debug(nil, "用户已退出.")
}
