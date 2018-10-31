package initial

import (
	"github.com/astaxie/beego/orm"
	"zerogo/entity"
)

func InitEntity() {
	//orm.RegisterModel(new(entity.Video))
	orm.RegisterModelWithPrefix("any_",new(entity.User))
}