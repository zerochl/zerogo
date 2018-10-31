package RedisManager

import (
	"zerogo/utils/redis"
	"encoding/json"
	"log"
	"time"
	"zerogo/entity"
)

func SaveVideos(key string, videos []entity.Video) {
	videosJson, err := json.Marshal(videos)
	if (nil != err) {
		log.Println("in redis manager save videos error:", err.Error())
		return
	}
	redisIO.Set(key, string(videosJson), int64(time.Hour * 24 / time.Second))
}

func SaveVideoWithHash(key,hashKey string,video *entity.Video){
	videosJson, err := json.Marshal(video)
	if (nil != err) {
		log.Println("in redis manager save video error:", err.Error())
		return
	}
	redisIO.Hset(key,hashKey,string(videosJson),int64(time.Hour * 24 / time.Second))
}

//func SaveEpisodeWithHash(key,hashKey string,episodeList []entity.Episode){
//	videosJson, err := json.Marshal(episodeList)
//	if (nil != err) {
//		log.Println("in redis manager save video error:", err.Error())
//		return
//	}
//	redisIO.Hset(key,hashKey,string(videosJson),int64(time.Hour * 24 / time.Second))
//}

func GetVideosByKeyAndTag(key, tag string) []entity.Video {
	realKey := key + "_" + tag;
	cacheValue, err := redisIO.Get(realKey)
	if (nil != err) {
		log.Println("in redis get videos error:", err.Error())
		return nil
	}
	//log.Println("json:",string(cacheValue))
	videoList := make([]entity.Video, 5)
	err = json.Unmarshal(cacheValue, &videoList)
	if (nil != err) {
		log.Println("in redis json to object error:", err.Error())
		return nil
	}
	return videoList
}

func GetHashValueByKey(key,hashKey string) string {
	cacheValue,err := redisIO.Hget(key,hashKey)
	if(nil != err){
		log.Println("get cache from redis false:",err.Error())
		return ""
	}
	return string(cacheValue)
}
