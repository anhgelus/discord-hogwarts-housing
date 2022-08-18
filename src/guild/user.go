package guild

import (
	"fmt"
	"github.com/anhgelus/discord-hogwarts-housing/src/database"
	"github.com/mitchellh/mapstructure"
)

type User struct { // user:{guild-id}:{user-id}
	Id      string
	HouseId string
}

// GetUser
// id - User id
// guildId - Guild id
func GetUser(id string, guildId string) User {
	k := fmt.Sprintf("user:%s:%s", guildId, id)
	values, err := database.RedisHget(database.GetRedisPool(), k)
	if err != nil {
		panic(err)
	}
	user := User{}
	err = mapstructure.Decode(values, &user)
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	return user
}

// CreateUser
// id - User id
// guildId - Guild id
func CreateUser(id string, guildId string) (User, error) {
	user := User{
		Id:      id,
		HouseId: "0",
	}
	_, err := database.RedisHset(database.GetRedisPool(), fmt.Sprintf("user:%s:%s", guildId, id), user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func UserExist(id string, guildId string) bool {
	c := database.GetRedisPool().Get()
	value, err := c.Do("EXISTS", fmt.Sprintf("user:%s:%s", guildId, id))
	if err != nil {
		return false
	}
	if value.(int64) == 1 {
		return true
	}
	return false
}
