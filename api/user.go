package api

import (
	"WAF/middlewares"
	"WAF/models"
	"WAF/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Tags User
// @Summary 用户信息
// accept	json
// produce	json
// @Success 200 {string} string	"ok"
// @Router /user/ [get]

func User(c *gin.Context) {
	claims, _ := c.Get("claims")
	waitUse := claims.(*middlewares.UserClaim)
	utils.OkData(waitUse, c)

}

func AddUser(c *gin.Context) {
	var user models.User
	var err error
	err = c.ShouldBindJSON(&user)
	utils.Initvalidate()
	err = utils.Validate.Struct(user)
	if err != nil {
		utils.HandleError(err, c)
		return
	}
	err = models.NewUser(user)
	if err != nil {
		utils.FailMessage(err.Error(), c)
		return
	}
	utils.Ok(c)
}

func CheckUser(c *gin.Context) {
	var user models.User
	var err error
	v := c.Param("id")
	id := StringToInt64(v)
	user, err = models.CheckUser(id)
	if err != nil {
		utils.FailMessage(err.Error(), c)
		return
	}
	utils.OkData(gin.H{
		"id":       user.Id,
		"name":     user.Name,
		"password": user.Passwd,
		"state":    user.State,
	}, c)
}

func EditUser(c *gin.Context) {
	var user models.User
	var err error
	err = c.ShouldBindJSON(&user)
	if user.Id == 0 {
		utils.FailMessage("请求参数错误", c)
		return
	}
	user.Passwd = models.EncryptPass(user.Passwd)
	err = models.UpdateUser(user)
	if err != nil {
		utils.FailMessage(err.Error(), c)
		return
	}
	utils.Ok(c)

}

func DelUser(c *gin.Context) {
	v := c.Param("id")
	id := StringToInt64(v)
	err := models.DelUser(id)
	if err != nil {
		utils.FailMessage(err.Error(), c)
		return
	}
	utils.Ok(c)
}

func StringToInt64(s string) (i int64) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	} else {
		return i
	}
}
