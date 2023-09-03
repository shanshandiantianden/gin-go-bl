package middlewares

import (
	"errors"
	"gin-go-bl/global"
	"github.com/dgrijalva/jwt-go"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"strings"
	"time"
)

type CustomClaims struct {
	ID       uint      //
	NickName string    //
	UUID     uuid.UUID //
	jwt.StandardClaims
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localSstorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("x-token")
		color.Blue(token)
		//校验token是否为空
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "token为空,请登录",
			})
			c.Abort()
			return
		}
		//校验token是否携带Bearer
		if !strings.HasPrefix(token, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "请登录",
			})
			c.Abort()
			return
		}
		// 提取token的有效部分,即"Bearer "之后的部分
		token = token[7:]
		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, TokenExpired) {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code": 401,
					"msg":  "授权已过期",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  err.Error(),
			})
			c.Abort()
			return
		}
		// gin的上下文记录claims和userId的值
		c.Set("claims", claims)
		c.Set("user_uuid", claims.UUID)
		c.Next()
	}
}

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)
var (
	TokenExpired_zh     = errors.New("token已过期")
	TokenNotValidYet_zh = errors.New("token尚未激活")
	TokenMalformed_zh   = errors.New("不是一个正确的token")
	TokenInvalid_zh     = errors.New("无法处理此token")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.SigningKey),
	}
}

// 创建一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {

		return j.SigningKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid

	}

}

// 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	//jwt.TimeFunc = func() time.Time {
	//	fmt.Println(time.Unix(0, 0))
	//	return time.Unix(0, 0)
	//}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		//fmt.Println(err)
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		//claims.StandardClaims.ExpiresAt = claims.StandardClaims.ExpiresAt + (time.Now().Unix() - claims.StandardClaims.NotBefore)
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
