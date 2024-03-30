package main

import (
	"encoding/json"
	"log"
	"os"

	"szechuansage/server"
	"szechuansage/slotgame"
)

func LoadGame() {
	content, err := os.ReadFile("198.json")
	if err != nil {
		log.Fatal("Error when loading config: ", err)
	}
	var payload slotgame.Game
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	log.Printf("Config name: %v", payload.Name)
	log.Printf("Config symbols: %v", payload.Symbols[0].Symbol)
	log.Printf("Config wilds: %v", payload.Wilds[0])
	log.Printf("Config scatters: %v", payload.Scatters[0])
	log.Printf("Config base: %v", payload.ReelSets[0])
	log.Printf("Config freeReels: %v", payload.ReelSets[1])
}

func RunSequenceTest() {
	server.Init("198.json")
	server.SequenceTest("base")
	// server.Init("198.json")
	// server.SequenceTest("free")
}

func RunRandomTest() {
	server.Init("198.json")
	server.RandomTest("base", 1e8)
	server.Init("198.json")
	server.RandomTest("free", 1e8)
}

func main() {
	// LoadGame()
	// RunSequenceTest()
	RunRandomTest()
}
