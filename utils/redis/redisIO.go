package redisIO
import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"AnyVideo-Go/caches"
	"errors"
	"AnyVideo-Go/utils"
	"log"
)

var redis *caches.MyRedisCache

func InitRedis(redisConfig string) error{
	cacheConfig := beego.AppConfig.String("cache")

	log.Printf("cacheConfig:%v \n",cacheConfig)

	var cc cache.Cache
	if "zeroredis" == cacheConfig {
		var err error

		defer utils.Recover("redis init falure")
		cc, err = cache.NewCache("zeroredis", redisConfig)

		if err != nil {
			log.Printf("%v", err)
			return err
		}
		cache, ok := cc.(*caches.MyRedisCache)
		if ok {
			redis = cache
		}else {
			log.Println("parse cache to MyRedisCache failure !")
		}
	}
	return nil
}

func Set(key,val string, expire int64) error {
	var err error
	data, err := utils.Encode(val)

	if err != nil {
		return err
	}

	if data == nil {
		return errors.New("cc is nil")
	}

	defer utils.Recover("redis set falure")

	err = redis.Set(key, data, expire)
	if err != nil {
		log.Printf("%v", err)
	}

	return err;
}

func Hset(key string,field,val string, expire int64) error {
	var err error
	data, err := utils.Encode(val)

	if err != nil {
		return err
	}

	if data == nil {
		return errors.New("cc is nil")
	}

	defer utils.Recover("redis set falure")

	err = redis.Hset(key, field,data, expire)
	if err != nil {
		log.Printf("%v", err)
	}

	return err;
}


func Get(key string) ([]byte,error) {
	defer utils.Recover("redis get falure")
	data := redis.Get(key)

	if data == nil {
		return nil,errors.New("key point value is nil ")
	}
	result,err := utils.Decode(data.([]byte))
	if err != nil {
		log.Println("decode failure", err)
	}
	return result,err
}

func IncrBy(key string, incr int) int {
	var err error
	var val int
	defer utils.Recover("redis get falure")
	val,err = redis.IncrBy(key,incr)
	if err != nil {
		log.Println("decode failure", err)
	}
	return val
}


func IncrByWithTimeOut(key string, incr int,timeOut int64) int {
	var err error
	var val int
	defer utils.Recover("redis get falure")
	val,err = redis.IncrByWithTimeOut(key,incr,timeOut)
	if err != nil {
		log.Println("decode failure", err)
	}
	return val
}

func Hget(key string,field string) ([]byte,error) {
	defer utils.Recover("redis get falure")
	data := redis.Hget(key,field)

	if data == nil {
		return nil,errors.New("key point value is nil ")
	}
	result,err := utils.Decode(data.([]byte))
	if err != nil {
		log.Println("decode failure", err)
	}
	return result,err
}

func Delete(key string) error {
	var err error
	defer utils.Recover("redis get falure")
	err = redis.Delete(key)

	if err != nil {
		log.Println("decode failure", err)
	}
	return err
}
