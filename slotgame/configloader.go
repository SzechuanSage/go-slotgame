package slotgame

import (
	"encoding/json"
	"log"
	"os"
)

type GameSymbol struct {
	Symbol string
	Value  string
	Pays   []int
}

type GameWild struct {
	Symbol     string
	Multiplier int
}

type GameScatter struct {
	Symbol string
}

type GameReelSet struct {
	Name  string
	Reels [][]string
}

type Game struct {
	Name      string
	Symbols   []GameSymbol
	Wilds     []GameWild
	Scatters  []GameScatter
	ReelSets  []GameReelSet
}

func ConfigLoader(file string) (Game, error) {
	var game Game

	content, err := os.ReadFile(file)
	if err != nil {
		log.Printf("Error when opening file: %v", err)
		return game, err
	}

	err = json.Unmarshal(content, &game)
	if err != nil {
		log.Printf("Error during Unmarshal(): %v", err)
		return game, err
	}

	return game, err
}
