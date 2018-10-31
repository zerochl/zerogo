package initial

import (
	"zerogo/utils/redis"
	"encoding/json"
	"log"
	"github.com/astaxie/beego"
)

func InitRedis()  {
	createRedis()
}

type RedisConfig struct {
	CONN string `json:"conn"`
	KEY  string `json:"key"`
	PASSWORD string `json:"password"`
	DBNUM string `json:"dbNum"`
}

func createRedis() {
	//redisConfig := &RedisConfig{Constant.REDIS_HOST + ":" + strconv.Itoa(Constant.REDIS_PORT),"",Constant.REDIS_PSW,Constant.REDIS_DB_NUM}
	redisConfig := &RedisConfig{beego.AppConfig.DefaultString("redis_host", ""),
		"",
		beego.AppConfig.DefaultString("redis_psw", ""),
		beego.AppConfig.DefaultString("redis_db_num", "0")}

	configJson,err := json.Marshal(redisConfig)
	if(err != nil){
		log.Println(err)
		return
	}
	log.Println("redis config json str:",string(configJson))
	err = redisIO.InitRedis(string(configJson))
	//singleton, err = cache.NewCache("zeroredis", string(configJson))
	if err != nil {
		log.Println("redis connect error:",err)
		//singleton = nil
		return
	}
	//singleton.Put("test2",123321,10 * time.Second)

	//err = Hset("beegoRedis","test3","12345678",int64(time.Hour * 24 /time.Second))
	//if(err != nil){
	//	log.Println("put error:",err)
	//}else{
	//	//log.Println("put success:",result)
	//	log.Println("put success:")
	//}
	//var result []byte
	//result,err = Hget("beegoRedis","test3")
	//if(err != nil){
	//	log.Println("put error:",err)
	//}else{
	//	log.Println("put success:",string(result))
	//	//log.Println("put success:")
	//}

	//encode,_:= utils.Encode("12345678")
	//log.Println("encode success:",string(encode))
	//var decode []byte
	//decode,err = utils.Decode(encode)
	//log.Println("decode success:",string(decode))
	//encode,_ := utils.Encode(12345678)
	//var decode int
	//utils.Decode(encode,&decode)
	//log.Println("decode success:",decode)
}
