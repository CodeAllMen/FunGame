package models

import (
	"github.com/astaxie/beego/orm"
)

type Games struct {
	Id      int64  `orm:"pk;auto"`
	Url     string `orm:"size(100)"`
	Img     string `orm:"size(100)"`
	Name    string `orm:"size(100)"`
	PlayNum int
}

func init() {
	orm.RegisterModel(new(Games))
}
