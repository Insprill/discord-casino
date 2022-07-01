package casino

import (
	"encoding/json"
	"errors"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"os"
)

var players = map[string]*Player{}

//region IO

func Save() {
	marshal, err := json.Marshal(players)
	if err != nil {
		panic(err)
		return
	}

	err = ioutil.WriteFile("data.json", marshal, 666)
	if err != nil {
		panic(err)
		return
	}

	println("Successfully saved user data")
}

func Load() {
	_, err := os.Stat("data.json")
	if errors.Is(err, os.ErrNotExist) {
		return
	}

	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		panic(err)
		return
	}

	err = json.Unmarshal(data, &players)
	if err != nil {
		panic(err)
		return
	}

	println("Successfully loaded user data")
}

//endregion

func GetPlayerById(userId string) *Player {
	if _, ok := players[userId]; ok {
		return players[userId]
	} else {
		players[userId] = &Player{
			ID: userId,
		}
		return players[userId]
	}
}

func GetPlayer(user *discordgo.User) *Player {
	return GetPlayerById(user.ID)
}
