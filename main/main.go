package main

import (
	"encoding/json"
	// "fmt"
	"log"
	"os"

	"szechuansage/bonsai"
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

func BonsaiTest() {
	var myReel bonsai.BonsaiReel
	log.Printf("Bonsai Reel: %v", myReel)

	var newSymbols = [3]int{0, 0, 1}
	result := myReel.SetSymbols(newSymbols)
	if result {
		log.Printf("Bonsai Reel: %v %v", myReel, myReel.GetValue())
	} else {
		log.Printf("There was an error setting symbols to %v", newSymbols)
	}

	var bonsai1, bonsai2, bonsai3 bonsai.BonsaiReel
	bonsai1.SetSymbols([3]int{2, 0, 1})
	bonsai2.SetSymbols([3]int{0, 1, 1})
	bonsai3.AddSymbols(bonsai1, bonsai2)
	log.Printf("Bonsai Reel: %v %v", bonsai3, bonsai3.GetValue())

	// mymap := bonsai.NewBonsais();
	// for _, v := range mymap {
	// 	fmt.Println(v)
	// }
	bonsai.FreeGames()
	bonsai.WeightedDraw()
	bonsai.ReelDisplays()
}

func RunBonsaiTest() {
	server.Init("bonsai.json")
	server.BonsaiTest("base", 1)
}

func main() {
	// LoadGame()
	// RunSequenceTest()
	// RunRandomTest()
	// BonsaiTest()
	RunBonsaiTest()
}
