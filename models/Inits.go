package models

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"log"
)

var X *xorm.Engine

func init() {
	var err error
	fmt.Println("11111")
	X, err = xorm.NewEngine("mssql", "server=.;database=tests;user id=sa;password=111111")

	if err != nil {
		log.Fatalf("failt to carete entine")
	}
	X.SetMapper(core.SameMapper{})
	fmt.Println("-------11111")
	if err = X.Sync2(new(JYJGXX)); err != nil {
		log.Fatalf("failt JYJGXX to sync", err.Error())
	}
	if err = X.Sync2(new(TbUsers)); err != nil {
		log.Fatalf("failt TbUsers to sync", err.Error())
	}

	//初始化用户
	// var tuser = new(TbUsers).Init()
	// var d = tuser.GetUser("root")
	// if d == false {
	// 	_, _ = X.Exec("insert into [TbUsers](jyjgbh,[name],nick,pwd,[type],roles,state)values('0','root','超级管理员','1','0','1',1) ")
	// }
}
