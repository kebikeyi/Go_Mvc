package models

type TbUsers struct {
	ID        int64  //
	JYJGBH    string `xorm:"varchar(20)"`  //
	Name      string `xorm:"varchar(32)"`  //
	Nick      string `xorm:"varchar(32)"`  //
	Pwd       string `xorm:"varchar(100)"` //
	Type      int    //
	AddTime   string `xorm:"varchar(8)"`    //
	Rights    string `xorm:"varchar(2000)"` //
	Roles     string `xorm:"varchar(500)"`  //
	LoginTime string `xorm:"varchar(8)"`    //
	LoginIP   string `xorm:"varchar(50)"`   //
	Memo      string `xorm:"varchar(200)"`  //
	State     int    //	4)"` //
	_godb     `xorm:"-"`
}

//初始化函数
func (this *TbUsers) Init() *TbUsers {
	this._godb._bean = this
	this._godb._beans = make([]TbUsers, 0)
	return this
}
