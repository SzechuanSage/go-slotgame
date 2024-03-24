package server

import (
	"fmt"
	"szechuansage/slotgame"
)

var config slotgame.Config

var reels int
var reelSet [][]string

var report slotgame.Report

var indexes []int
var loopTo []int
var endOfSequence bool

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

func setNextReelView() []map[string]int {
	if !isNextSequence() {
		endOfSequence = true
	}
	_, symbolView := config.GetReelView(indexes, "reels")
	if !endOfSequence {
		advanceSequence()
	}
	return symbolView
}

func Init() {
	config = slotgame.GetConfig()

	reels = config.Reels()
	reelSet = config.ReelSet("reels")

	report = slotgame.InitReport(config.Symbols(), config.Reels())

	indexes = make([]int, config.Reels())
	loopTo = make([]int, config.Reels())
	endOfSequence = false
}

func SequenceTest() {
	for x, y := range reelSet {
		loopTo[x] = len(y)
	}

	var symbolView []map[string]int
	var times int32
	var ofAKind int
	var scatters int

	for hasNextSequence() {
		symbolView = setNextReelView()

		report.AccumulateTotal("count", 1)

		scatters = 0
		for _, symbols := range symbolView {
			scatters += symbols["S"]
		}
		report.AccumulateCombinations("S", scatters, 1)

		for symbol, symbolC := range symbolView[0] {
			if config.SymbolIsWay(symbol) {
				times = int32(symbolC)
				ofAKind = 1
				for index, reels := range symbolView[1:] {
					if (reels[symbol] == 0) && (reels["Z"] == 0) {
						break
					}
					times *= int32(reels[symbol] + reels["Z"])
					ofAKind = index + 2
				}
				report.AccumulateCombinations(symbol, ofAKind, times)
			}
		}
	}

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
