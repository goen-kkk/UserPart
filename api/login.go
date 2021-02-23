package api

import (
	"WAF/middlewares"
	"WAF/models"
	"WAF/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// @Tags User
// @Summary 用户登录
// accept	json
// produce	json
// @Param   name      query    string     true        "用户名"
// @Param   passwd      query    string     true        "密码"
// @Success 200 {string} string	"ok"
// @Router /login/ [post]

func Login(c *gin.Context) {
	var user models.User
	var user2 models.User
	var err error
	err = c.ShouldBindJSON(&user)

	utils.Initvalidate()
	err = utils.Validate.Struct(user)
	if err != nil {
		utils.HandleError(err, c)
		return
	}
	user2, err = models.GetUserByName(user.Name)

	if user2.Passwd != models.EncryptPass(user.Passwd) {
		utils.FailMessage("用户名或密码不对", c)
		return
	}
	// 修改登录状态
	user2.State = 1
	models.UpdateUser(user2)

	token, ExpiresAt := CreateToken(c, &user2)
	utils.OkData(gin.H{
		"token":      token,
		"user":       user2,
		"expires_at": ExpiresAt,
	}, c)

}

func CreateToken(c *gin.Context, user *models.User) (string, int64) {
	//自定义claim
	claim := middlewares.UserClaim{
		Id:   user.Id,
		Name: user.Name,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),       // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 60*60*24*7), // 过期时间 一周
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, _ := token.SignedString([]byte("12345678"))
	return tokenString, claim.StandardClaims.ExpiresAt
}
