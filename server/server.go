package server

import (
	"fmt"
	"log"
	"math/rand"
	"szechuansage/slotgame"
)

var config slotgame.Config

var reels int
var reelSet [][]string

var report slotgame.Report

var indexes []int
var loopTo []int
var endOfSequence bool

var symbolView []map[string]int

func hasNextSequence() bool {
	return !endOfSequence
}

func isNextSequence() bool {
	for i := reels - 1; i >= 0; i-- {
		if indexes[i] != loopTo[i]-1 {
			return true
		}
	}
	return false
}

func atEndOfLoop(index int) bool {
	return (indexes[index]+1 >= loopTo[index])
}

func advanceSequence() {
	for i := reels - 1; i >= 0; i-- {
		if atEndOfLoop(i) {
			indexes[i] = 0
		} else {
			indexes[i]++
			return
		}
	}
}

func setNextReelView(reelSetName string) []map[string]int {
	if !isNextSequence() {
		endOfSequence = true
	}
	_, symbolView := config.GetReelView(indexes, reelSetName)
	if !endOfSequence {
		advanceSequence()
	}
	return symbolView
}

func setRandomReelView(reelSetName string) []map[string]int {
	for x, y := range loopTo {
		indexes[x] = rand.Intn(y)
	}
	_, symbolView := config.GetReelView(indexes, reelSetName)
	return symbolView
}

func evaluateReelView() {
	var times int64
	var ofAKind int

	for symbol, symbolC := range symbolView[0] {
		if config.SymbolIsWay(symbol) {
			times = int64(symbolC)
			ofAKind = 1
			for index, reels := range symbolView[1:] {
				if (reels[symbol] == 0) && (reels["Z"] == 0) {
					break
				}
				times *= int64(reels[symbol] + reels["Z"])
				ofAKind = index + 2
			}
			report.AccumulateCombinations(symbol, ofAKind, times)
		}
	}
}

func produceReport() {
	for _, symbol := range config.Symbols() {
		var payTable = config.SymbolPays(symbol)
		for count, pay := range payTable {
			var c = report.GetCombinations(symbol, count+1)
			report.AccumulatePays(symbol, count+1, int64(c) * int64(pay))
		}
	}

	fmt.Println(symbolView)
	report.PrintTotals()
	fmt.Println("Combinations")
	report.PrintCombinations()
	fmt.Println("Pays")
	report.PrintPays()
}

func Init(game string) {
	gameConfig, err := slotgame.ConfigLoader(game)
	if err != nil {
		log.Fatal("Error when loading config: ", err)
	}

	config = slotgame.GetConfig(gameConfig)

	reels = config.Reels()

	report = slotgame.InitReport(config.Symbols(), config.Reels())

	indexes = make([]int, config.Reels())
	loopTo = make([]int, config.Reels())
	endOfSequence = false
}

func SequenceTest(reelSetName string) {
	var scatters int

	reelSet = config.ReelSet(reelSetName)
	for x, y := range reelSet {
		loopTo[x] = len(y)
	}

	for hasNextSequence() {
		symbolView = setNextReelView(reelSetName)

		report.AccumulateTotal("count", 1)

		scatters = 0
		for _, symbols := range symbolView {
			scatters += symbols["S"]
		}
		report.AccumulateCombinations("S", scatters, 1)

		evaluateReelView()
	}

	produceReport()
}

// RandomTest performs testSpins spins of a slot game
func RandomTest(reelSetName string, testSpins uint32) {
	var spin uint32
	var scatters int

	reelSet = config.ReelSet(reelSetName)
	for x, y := range reelSet {
		loopTo[x] = len(y)
	}

	for spin = 0; spin < testSpins; spin += 1 {
		symbolView = setRandomReelView(reelSetName)

		report.AccumulateTotal("count", 1)

		scatters = 0
		for _, symbols := range symbolView {
			scatters += symbols["S"]
		}
		report.AccumulateCombinations("S", scatters, 1)

		evaluateReelView()
	}

	produceReport()
}
