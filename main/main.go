package main

import (
	"log"
	"szechuansage/slotgame"
	"szechuansage/server"
)

func LoadConfig() {
	config, err := slotgame.ConfigLoader("./config.json")
	if err != nil {
		log.Fatal("Error when loading config: ", err)
	}
	log.Printf("Config user: %v", config["user"])
}

func RunSequenceTest() {
	server.Init()
	server.SequenceTest()
}

func main() {
	// LoadConfig()
	RunSequenceTest()
}
