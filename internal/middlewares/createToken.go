package middlewares

import (
	"fmt"
	"gin-go-bl/internal/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)

func CreateToken(c *gin.Context, sessionInfo *models.SessionUserInfo) string {
	//生成token信息
	j := NewJWT()
	claims := CustomClaims{
		ID:       sessionInfo.UserID,
		NickName: sessionInfo.UserName,
		UUID:     sessionInfo.UUID,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			// TODO 设置token过期时间
			ExpiresAt: time.Now().Unix() + 60*60*24*7, //token -->7天过期
			//发放用户
			Issuer: "test",
			// 发放时间
			IssuedAt: time.Now().Unix(),
		},
	}
	//生成token
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":   401,
			"msg":    "token生成失败,重新再试",
			"isUser": "test",
		})
		return ""
	}
	return token
}
func CreateToken_t(Id int, NickName string, uuid uuid.UUID) string {
	//生成token信息
	j := NewJWT()
	claims := CustomClaims{
		ID:       uint(Id),
		NickName: NickName,
		UUID:     uuid,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			// TODO 设置token过期时间
			ExpiresAt: time.Now().Unix() + 60*60, //token -->30天过期
			//发放用户
			Issuer: "test",
			// 发放时间
			IssuedAt: time.Now().Unix(),
		},
	}
	//生成token
	token, err := j.CreateToken(claims)
	if err != nil {
		fmt.Println("token生成失败,重新再试  :   ", err)
		return ""
	}
	return token
}
