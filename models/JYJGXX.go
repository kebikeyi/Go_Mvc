package models

type JYJGXX struct {
	ID       int64  //
	ProvID   string `xorm:"varchar(6)"`
	CityID   string `xorm:"varchar(6)"`
	CountyID string `xorm:"varchar(6)"`
	JYJGLX   string `xorm:"varchar(20)"` //检验机构类型 0系统 1运管站 2监测站
	JYJGBH   string `xorm:"varchar(20)"`
	JYJGMC   string `xorm:"varchar(20)"`
	JKXLH    string `xorm:"varchar(200)"`
	Phone    string `xorm:"varchar(50)"`
	Contact  string `xorm:"varchar(50)"`
	State    int
	_godb    `xorm:"-"`
}

//初始化函数
func (this *JYJGXX) Init() *JYJGXX {
	this._godb._bean = this
	this._godb._beans = make([]JYJGXX, 0)
	return this
}
