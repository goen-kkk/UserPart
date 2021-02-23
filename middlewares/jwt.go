package middlewares

import (
	"WAF/models"
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserClaim struct {
	Id                 int64
	Name               string `json:"name"`
	State              int    `json:"state"`
	jwt.StandardClaims        //嵌套了这个结构体就实现了Claim接口
}

/*
使用 JWT 做用户认证
- 基于 token 的认证方式
1. 用户向服务器发送用户名和密码。
2. 服务器将用户 ID，认证有效期等信息签名后生成 token 返回给客户端。
3. 客户端将 token 写入本地存储。
4. 用户随后的每一次请求，都将 token 附加到 header 中。
5. 服务端获取到用户请求的 header，拿到用户数据并且做签名校验
*/

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "未登录或非法访问",
			})
			c.Abort()
			return
		}
		claims, err := ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}

}

var (
	TokenExpired     error = errors.New("Token is expired")
	TokenNotValidYet error = errors.New("Token not active yet")
	TokenMalformed   error = errors.New("That's not even a token")
	TokenInvalid     error = errors.New("Couldn't handle this token:")
)

func ParseToken(tokenString string) (claim *UserClaim, err error) {
	// token, err := jwt.Parse(tokenString, secret())
	token, err := jwt.ParseWithClaims(tokenString, &UserClaim{}, secret())
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
	claim = token.Claims.(*UserClaim)
	if token.Valid {
		return claim, nil
	}
	return nil, TokenInvalid
}

func secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(models.Cfg.Section("server").Key("JWtSECRET").String()), nil
	}
}
