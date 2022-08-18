package guild

import (
	"fmt"
	"github.com/anhgelus/discord-hogwarts-housing/src/database"
	"github.com/gomodule/redigo/redis"
)

type Config struct { // config:{guild-id}
	Houses string // Houses is a string of house ids separated by commas.
}

type House struct { // house:{guild-id}:{house-id}
	Id   int
	Name string
	User string
}

func GetHouse(id int, guildId int) House {
	k := fmt.Sprintf("house:%v:%v", guildId, id)
	values, err := database.RedisHgetAll(database.GetRedisPool(), k)
	if err != nil {
		panic(err)
	}
	house := House{}
	err = redis.ScanStruct(values, &house)
	if err != nil {
		panic(err)
	}
	return house
}
