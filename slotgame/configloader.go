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

type Game struct {
	Name      string
	Symbols   []GameSymbol
	Wilds     []GameWild
	Scatters  []GameScatter
	Base      [][]string
	FreeReels [][]string
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
