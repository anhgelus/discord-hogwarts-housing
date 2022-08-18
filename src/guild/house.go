package guild

import (
	"fmt"
	"github.com/anhgelus/discord-hogwarts-housing/src/database"
	"github.com/mitchellh/mapstructure"
	"strings"
)

type House struct { // house:{guild-id}:{house-id}
	Id   string
	Name string
	User string
}

func GetHouse(id string, guildId string) (House, error) {
	if !HouseExist(id, guildId) {
		return House{}, error(fmt.Errorf("house %s does not exist", id))
	}
	k := fmt.Sprintf("house:%s:%s", guildId, id)
	values, err := database.RedisHget(database.GetRedisPool(), k)
	if err != nil {
		panic(err)
	}
	house := House{}
	err = mapstructure.Decode(values, &house)
	if err != nil {
		panic(err)
	}
	if house.Name == "" || house.User == "" || house.Id == "" {
		return house, error(fmt.Errorf("house %s is not valid", id))
	}
	return house, nil
}

func GetHouses(guildId string) ([]House, error) { // house:{guild-id}:all
	k := fmt.Sprintf("house:%s:all", guildId)
	got, err := database.RedisGet(database.GetRedisPool(), k)
	if err != nil {
		return []House{}, err
	}
	splited := strings.Split(got, ",")[0 : len(got)-1]
	fmt.Println(splited)
	var houses []House
	for _, id := range splited {
		house, err := GetHouse(id, guildId)
		if err != nil {
			fmt.Println(id)
			panic(err)
		}
		houses = append(houses, house)
	}
	return houses, nil
}

func getLastHouseId(guildId string) string {
	houses, err := GetHouses(guildId)
	if err != nil {
		return "0"
	}
	lastHouse := houses[len(houses)-1]
	return lastHouse.Id
}

func HouseExist(id string, guildId string) bool {
	c := database.GetRedisPool().Get()
	value, err := c.Do("EXISTS", fmt.Sprintf("house:%s:%s", guildId, id))
	if err != nil {
		return false
	}
	if value.(int64) == 1 {
		return true
	}
	return false
}

func CreateHouse(guildId string, name string, userId string) (House, error) {
	// Save the house
	var lastId int
	lastId, err := fmt.Sscanf(getLastHouseId(guildId), "%d", &lastId)
	if err != nil {
		return House{}, err
	}
	id := fmt.Sprintf("%d", lastId+1)
	house := House{
		Name: name,
		User: userId + ",",
		Id:   id,
	}
	pool := database.GetRedisPool()
	_, err = database.RedisHset(pool, fmt.Sprintf("house:%s:%s", guildId, id), house)
	if err != nil {
		return house, err
	}
	// add the house id to the list of id's
	ids, err := database.RedisGet(pool, fmt.Sprintf("house:%s:all", guildId))
	ids = ids + id
	if err != nil {
		ids = id
	}
	_, err = database.RedisSet(pool, fmt.Sprintf("house:%s:all", guildId), fmt.Sprintf("%s,", ids))
	if err != nil {
		return house, err
	}
	// Register the user if the user do not exist
	if !UserExist(userId, guildId) {
		_, err := CreateUser(userId, guildId)
		if err != nil {
			return house, err
		}
	}
	return house, nil
}
