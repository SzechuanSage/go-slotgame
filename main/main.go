package main

import (
	"log"
	"szechuansage/slotgame"
)

func main() {
	config, err := slotgame.ConfigLoader("./config.json")
	if err != nil {
		log.Fatal("Error when loading config: ", err)
	}
	log.Printf("Config user: %v", config["user"])
}
