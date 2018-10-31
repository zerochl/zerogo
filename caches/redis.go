package caches
import (
	"github.com/astaxie/beego/cache"
	"github.com/garyburd/redigo/redis"
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

var (
// the collection name of redis for cache adapter.
	DefaultKey string = "zeroredis"
)

type MyRedisCache struct {
	p        *redis.Pool // redis connection pool
	conninfo string
	dbNum    int
	key      string
	password string
}

// create new redis cache with default collection name.
func NewRedisCache() cache.Cache {
	return &MyRedisCache{key: DefaultKey}
}

// actually do the redis cmds
func (rc *MyRedisCache) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	c := rc.p.Get()
	defer c.Close()

	return c.Do(commandName, args...)
}

// Get cache from redis.
func (rc *MyRedisCache) Get(key string) interface{} {
	if v, err := rc.Do("GET", key); err == nil {
		return v
	}
	return nil
}

// HGet cache from redis.
func (rc *MyRedisCache) Hget(key string,field string) interface{} {
	if v, err := rc.Do("HGET", key,field); err == nil {
		return v
	}
	return nil
}


// GetMulti get cache from redis.
func (rc *MyRedisCache) GetMulti(keys []string) []interface{} {
	size := len(keys)
	var rv []interface{}
	c := rc.p.Get()
	defer c.Close()
	var err error
	for _, key := range keys {
		err = c.Send("GET", key)
		if err != nil {
			goto ERROR
		}
	}
	if err = c.Flush(); err != nil {
		goto ERROR
	}
	for i := 0; i < size; i++ {
		if v, err := c.Receive(); err == nil {
			rv = append(rv, v.([]byte))
		} else {
			rv = append(rv, err)
		}
	}
	return rv
	ERROR:
	rv = rv[0:0]
	for i := 0; i < size; i++ {
		rv = append(rv, nil)
	}

	return rv
}

// put cache to redis.
func (rc *MyRedisCache) Put(key string, val interface{}, timeout time.Duration) error {
	var err error
	if _, err = rc.Do("SETEX", key, timeout/time.Second, val); err != nil {
		return err
	}
	return err
}

func (rc *MyRedisCache) Set(key string, val interface{}, timeout int64) error {
	var err error
	_, err = rc.Do("SETEX", key, timeout, val)
	return err;
}

func (rc *MyRedisCache) Hset(key string,field string, val interface{}, timeout int64) error {
	var err error
	_, err = rc.Do("HSET", key,field ,val)
	rc.Do("EXPIRE",key,timeout)
	return err;
}



// delete cache in redis.
func (rc *MyRedisCache) Delete(key string) error {
	var err error
	_, err = rc.Do("DEL", key)
	return err
}


// check cache's existence in redis.
func (rc *MyRedisCache) IsExist(key string) bool {
	v, err := redis.Bool(rc.Do("EXISTS", key))
	if err != nil {
		return false
	}
	return v
}

// increase counter in redis.
func (rc *MyRedisCache) Incr(key string) error {
	_, err := redis.Bool(rc.Do("INCRBY", key, 1))
	return err
}

// increase counter in redis.
func (rc *MyRedisCache) IncrBy(key string,incr int) (int,error) {
	number, err := redis.Int(rc.Do("INCRBY", key, incr))
	return number,err
}

// increase counter in redis.
func (rc *MyRedisCache) IncrByWithTimeOut(key string,incr int,timeout int64) (int,error) {
	number, err := redis.Int(rc.Do("INCRBY", key, incr))
	rc.Do("EXPIRE",key,timeout)
	return number,err
}

// decrease counter in redis.
func (rc *MyRedisCache) Decr(key string) error {
	_, err := redis.Bool(rc.Do("INCRBY", key, -1))
	return err
}

// clean all cache in redis. delete this redis collection.
func (rc *MyRedisCache) ClearAll() error {
	cachedKeys, err := redis.Strings(rc.Do("HKEYS", rc.key))
	if err != nil {
		return err
	}
	for _, str := range cachedKeys {
		if _, err = rc.Do("DEL", str); err != nil {
			return err
		}
	}
	_, err = rc.Do("DEL", rc.key)
	return err
}

// start redis cache adapter.
// config is like {"key":"collection key","conn":"connection info","dbNum":"0"}
// the cache item in redis are stored forever,
// so no gc operation.
func (rc *MyRedisCache) StartAndGC(config string) error {
	var cf map[string]string
	json.Unmarshal([]byte(config), &cf)

	if _, ok := cf["key"]; !ok {
		cf["key"] = DefaultKey
	}
	if _, ok := cf["conn"]; !ok {
		return errors.New("config has no conn key")
	}
	if _, ok := cf["dbNum"]; !ok {
		cf["dbNum"] = "0"
	}
	if _, ok := cf["password"]; !ok {
		cf["password"] = ""
	}
	rc.key = cf["key"]
	rc.conninfo = cf["conn"]
	rc.dbNum, _ = strconv.Atoi(cf["dbNum"])
	rc.password = cf["password"]

	rc.connectInit()

	c := rc.p.Get()
	defer c.Close()

	return c.Err()
}

// connect to redis.
func (rc *MyRedisCache) connectInit() {
	dialFunc := func() (c redis.Conn, err error) {
		c, err = redis.Dial("tcp", rc.conninfo)
		if err != nil {
			return nil, err
		}

		if rc.password != "" {
			if _, err := c.Do("AUTH", rc.password); err != nil {
				c.Close()
				return nil, err
			}
		}

		_, selecterr := c.Do("SELECT", rc.dbNum)
		if selecterr != nil {
			c.Close()
			return nil, selecterr
		}
		return
	}
	// initialize a new pool
	rc.p = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 180 * time.Second,
		Dial:        dialFunc,
	}
}

func init() {
	cache.Register("zeroredis", NewRedisCache)
}

