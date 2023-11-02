package middle

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/x14n/x14n-gin-api/global"
	"github.com/x14n/x14n-gin-api/internal/model"
	"github.com/x14n/x14n-gin-api/internal/response"
	"go.uber.org/zap"
)

// token gin 中间件
func JWTAuthMiddle(ctx *gin.Context) {
	token := GetTokenFromHeader(ctx)
	global.GLogger.Sugar().Infof("token: %s", token)
	if token == "" {
		response.ErrorWithMsg(ctx, "token不能为空")
		ctx.Abort()
		return
	}
	//设置到上下文

	ctx.Next()
}

// 设置token到上下文
// func SetTokenToContext(ctx *gin.Context, userClaim *model.UserCliams, token string) {
// 	//
// }

// 从请求头中获取token
func GetTokenFromHeader(ctx *gin.Context) string {
	var token string
	token = ctx.Request.Header.Get("Authorization")
	if token != "" {
		return token
	}
	if ctx.Request.Method == http.MethodGet {
		token, ok := ctx.GetQuery("Authorization")
		if ok {
			return token
		}
	}

	if ctx.Request.Method == http.MethodPost {
		postParam := make(map[string]interface{})
		_ = ctx.ShouldBindJSON(&postParam)
		token, ok := postParam["Authorization"]
		if ok {
			return token.(string)
		}
	}
	return ""
}

// 创建token
func CreateToken(uid uint) (string, error) {
	newWithClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.UserCliams{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(global.GConfig.Jwt.ExpireTime).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    global.GConfig.Jwt.Issure,
		},
		Uid: uid,
	})
	return newWithClaim.SignedString([]byte(global.GConfig.Jwt.Securt))
}

// 解析Token
func ParseToken(tokenString string) (*model.UserCliams, error) {
	var err error
	var token *jwt.Token

	token, err = jwt.ParseWithClaims(tokenString, &model.UserCliams{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.GConfig.Jwt.Securt), nil
	})
	if err != nil {
		global.GLogger.Error("解析token失败", zap.String("err", err.Error()))
		return nil, err
	}

	//断言
	userCliams, ok := token.Claims.(*model.UserCliams)

	if !ok || token.Valid {
		return nil, errors.New("Jwt验证失败")
	}

	return userCliams, nil

}
