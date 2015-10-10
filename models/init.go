package models

import (
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/go-xorm/xorm"
	//_ "github.com/mattn/go-sqlite3"
	"fmt"
	//"log"
)

//var X *xorm.Engine
func init() {
	fmt.Println("--------A")
	//var err error
	//X,err=xorm.NewEngine("mysql","server=.;database=test;user id=sa;password=111111")
	//多个数据并且需要分区的应用，可以创建多个 X
	/* for i:=0;i<5;i++{
	   	X[i],err=xorm.NewEngine("mysql","server=.;database=test;user id=sa;password=111111")
	   }
	    X.Close()//关闭数据连接

	*/
	/* if err!=nil {

	   	log.Fatalf("")
	   }
	   if err=X.Sync(new(User));err!=nil {

	   	log.Fatalf("",err.Error())

	   }
	   log.Printin("hello")
	   fmt.Printin("hello...")*/
}

//自定义时间类型
/*type JsonTime time.Time
func (j JsonTime) MarshalJson() ([]byte,error) {
	return []byte('"'+time.Time(j).Format("2015-01-01 15:04:05")+'"'),nil
}
type User struct{
	Id int64
	Name string 'xorm:"autoincr"'
	CreatedAt int64 'xorm:"created"'
}
 tbMapper:=core.NewPrefixMapper(core.SnakeMapper{},"prefix_")
 Engine.SetTableMapper(tbMapper)

 Engine.CreateTables()//创建表
 Engine.DropTables()//删除表
 Engine.ShowWarn=true
 err :=Engine.Sync(new(User),new(Group))
 Engine.DumpAll(w io.Writer)
 Engine.Import(r io.Reader)
 //获取单条数据
 user :=new(User)
 has,err:=Engine.Id(id).Get(user)
 has,err:=Engine.Id(xorm.PK{1,2}.Get(user))
 //Get()方法
 //根据where来获取单挑数据
 user :=new(User)
 has,err:=Engine.Where("name=?","xlw").Get(user)
 //根据user结构体中已有的非空数据获得单条数据
 user:=&User{Id:1}
 has,err:=Engine.Get(user)
 //用其它条件
 user:=&User(Name:"xlw")
 has,err:=Engine.Get(user)
 //Find()方法
 everyone:=make([]Userinfo,0)
 err:=Engine.Find(&EveryOne)
 //map用户返回数据
 user:=make(map[int64]Userinfo)
 err:=Engine.Find(&Users)
 //加入条件
 Users:=make([Userinfo,0])
 err:=Engine.Where("age>? or name-=?",30,"xlw").Limit(20,10).Find(&Users)*/
