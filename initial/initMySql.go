package initial

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"time"
	//"github.com/gogather/com/log"
	//"PhotoBBS/entity"
	"log"
	"zerogo/entity"
)

func InitMySql()  {
	userName := beego.AppConfig.String("mysql_user_name")
	userPass := beego.AppConfig.String("mysql_user_psw")
	ip := beego.AppConfig.String("mysql_ip")
	port := beego.AppConfig.String("mysql_port")
	dbName := beego.AppConfig.String("mysql_db_name")
	maxIdle,_ := beego.AppConfig.Int("mysql_max_idle")
	maxConn,_ := beego.AppConfig.Int("mysql_max_conn")
	orm.Debug = true
	driver_url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local", userName, userPass, ip, port, dbName)
	beego.Info("driver_url:", driver_url)
	orm.RegisterDataBase("default", "mysql", driver_url,maxIdle,maxConn)
	// 设置为 UTC 时间
	orm.DefaultTimeLoc = time.UTC

	//db, err := orm.GetDB()
	//if err != nil {
	//	fmt.Println("get default DataBase error")
	//}

	o := orm.NewOrm()
	var qs orm.QuerySeter
	qs = o.QueryTable("any_user")
	count,_ := qs.Count()
	log.Println("count:",count)

	//qsId := qs.Filter("Id",10)
	var users []*entity.User
	num,_ := qs.All(&users)
	log.Println("num:",num,";user one name:",users[0].NickName)//,users[0].NickName)

	//qs.Filter("id", 1) // WHERE id = 1
	//qs.Filter("profile__age", 18) // WHERE profile.age = 18
	//qs.Filter("Profile__Age", 18) // 使用字段名和 Field 名都是允许的
	//qs.Filter("profile__age", 18) // WHERE profile.age = 18
	//qs.Filter("profile__age__gt", 18) // WHERE profile.age > 18
	//qs.Filter("profile__age__gte", 18) // WHERE profile.age >= 18
	//qs.Filter("profile__age__in", 18, 20) // WHERE profile.age IN (18, 20)
	//
	//qs.Filter("profile__age__in", 18, 20).Exclude("profile__lt", 1000)
	//// WHERE profile.age IN (18, 20) AND NOT profile_id < 1000

	//当前支持的操作符号：
	//
	//exact / iexact 等于
	//contains / icontains 包含
	//gt / gte 大于 / 大于等于
	//lt / lte 小于 / 小于等于
	//startswith / istartswith 以…起始
	//endswith / iendswith 以…结束
	//in
	//isnull
	//后面以 i 开头的表示：大小写不敏感

	//user := new(entity.User)
	//user.Id = 1
	//o.Read(user)
	//log.Println("user name:",user.Nickname)
}
