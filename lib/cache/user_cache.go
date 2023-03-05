package cache

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"gowebsocket/lib/redislib"
	"gowebsocket/models"
	"runtime/debug"
)

const (
	userOnlinePrefix    = "online:" // 用户在线状态
	userOnlineCacheTime = 12 * 60 * 60
)

/*********************  查询用户是否在线  ************************/
func GetUserOnlineKey(userKey string) (key string) {
	key = fmt.Sprintf("%s%s", userOnlinePrefix, userKey)

	return
}

func GetUserOnlineInfo(userKey string) (userOnline *models.UserOnline, err error) {
	redisClient := redislib.GetClient()

	key := GetUserOnlineKey(userKey)

	data, err := redisClient.Get(key).Bytes()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("GetUserOnlineInfo", userKey, err)

			return
		}

		fmt.Println("GetUserOnlineInfo", userKey, err)

		return
	}

	userOnline = &models.UserOnline{}
	err = json.Unmarshal(data, userOnline)
	if err != nil {
		fmt.Println("获取用户在线数据 json Unmarshal", userKey, err)

		return
	}

	fmt.Println("获取用户在线数据", userKey, "time", userOnline.LoginTime, userOnline.HeartbeatTime, "AccIp", userOnline.AccIp, userOnline.IsLogoff)

	return
}

// 设置用户在线数据
func SetUserOnlineInfo(userKey string, userOnline *models.UserOnline) (err error) {

	redisClient := redislib.GetClient()
	// redis的key："acc:user:online:userKey"
	key := GetUserOnlineKey(userKey)

	valueByte, err := json.Marshal(userOnline)
	if err != nil {
		fmt.Println("设置用户在线数据 json Marshal", key, err)

		return
	}

	_, err = redisClient.Do("setEx", key, userOnlineCacheTime, string(valueByte)).Result()
	if err != nil {
		fmt.Println("设置用户在线数据 ", key, err)

		return
	}

	return
}

// DelOvertimeUsersInRedis 心跳超时后使用该函数删除缓存中超时用户数据
func DelOvertimeUsersInRedis(overtimeUsers []string) (err error) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("心跳超时后 删除部分用户数据失败", string(debug.Stack()), r)
		}
	}()

	redisClient := redislib.GetClient()
	for _, client := range overtimeUsers {
		val, err := redisClient.Do("Del", client).Result()
		if err != nil {
			fmt.Println("Key ", val, " Deleted Error")
			panic(err)
		}
		fmt.Println("Key ", val, " has been deleted")
	}

	return nil
}
