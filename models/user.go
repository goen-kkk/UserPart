package models

import (
	"crypto/md5"
	"errors"
	"fmt"
)

//state: 登录状态

type User struct {
	Id     int64  `xrom:"notnull unique autoincr"`
	Name   string `json:"name" validate:"required" xrom:"varchar(255) notnull"`
	Passwd string `json:"passwd" xrom:"varchar(255) notnull"`
	State  int    `json:"state" xrom:"notnull"`
}

func CheckUser(id int64) (user User, err error) {
	_, err = Engine.Where("id=?", id).Get(&user)
	return user, err
}

func NewUser(user User) (err error) {
	has, _ := Engine.Where("name=?", user.Name).Exist(&user)
	if has {
		return errors.New("用户名已注册")
	}
	user.Passwd = EncryptPass(user.Passwd)
	_, err = Engine.Insert(&user)
	return err
}

func UpdateUser(user User) (err error) {
	newuser := new(User)
	_, err = Engine.Where("id=?", user.Id).Get(newuser)
	newuser.Name = user.Name
	newuser.Passwd = user.Passwd
	newuser.State = user.State
	_, err = Engine.Where("id=?", user.Id).Update(newuser)

	return err
}

func DelUser(id int64) (err error) {
	_, err = Engine.Delete(&User{Id: id})
	return err
}

func GetUserByName(name string) (user User, err error) {
	_, err = Engine.Where("name=?", name).Desc("Id").Get(&user)
	return
}

func MakeMd5(srcStr string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(srcStr)))
}

func EncryptPass(src string) string {
	return fmt.Sprintf("%s", MakeMd5(MakeMd5(src)[5:10]))
}
