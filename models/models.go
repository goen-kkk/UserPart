package models

import (
	"fmt"
	"log"

	"github.com/go-ini/ini"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	Engine *xorm.Engine
	err    error
	Cfg    *ini.File
)

/*
读取 config 默认设置
连接数据库
新建用户 User 表
添加 root 用户
*/
func init() {
	log.SetPrefix("[waf]")
	var err error
	source := "conf/conf.ini"
	Cfg, err = ini.Load(source)
	// log.Println(Cfg, err)
	if err != nil {
		log.Panicln(err)
	}
	sec := Cfg.Section("database")
	Engine, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		sec.Key("USER").String(),
		sec.Key("PASSWD").String(),
		sec.Key("HOST").String(),
		sec.Key("NAME").String()))
	if err != nil {
		log.Panicf("Faild to connect to database, err:%v", err)
	}

	Engine.Sync2(new(User))

	ret, err := Engine.IsTableEmpty(new(User))
	if err == nil && ret {
		log.Printf("create new user:%v, password:%v\n", "root", "123456")
		NewUser(User{1, "root", "123456", 0})
	}

}
