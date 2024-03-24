package server

import (
	"fmt"
	"szechuansage/slotgame"
)

var c = slotgame.GetConfig()

var reels = c.Reels()
var reelSet = c.ReelSet("reels")

var r = slotgame.InitReport(c.Symbols(), c.Reels())

var indexes = make([]int, c.Reels())
var loopTo = make([]int, c.Reels())
var endOfSequence = false

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
	_, symbolView := c.GetReelView(indexes, "reels")
	if !endOfSequence {
		advanceSequence()
	}
	return symbolView
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

		r.AccumulateTotal("count", 1)

		scatters = 0
		for _, symbols := range symbolView {
			scatters += symbols["S"]
		}
		r.AccumulateCombinations("S", scatters, 1)

		for symbol, symbolC := range symbolView[0] {
			if c.SymbolIsWay(symbol) {
				times = int32(symbolC)
				ofAKind = 1
				for index, reels := range symbolView[1:] {
					if (reels[symbol] == 0) && (reels["Z"] == 0) {
						break
					}
					times *= int32(reels[symbol] + reels["Z"])
					ofAKind = index + 2
				}
				r.AccumulateCombinations(symbol, ofAKind, times)
			}
		}
	}

	for _, symbol := range c.Symbols() {
		var payTable = c.SymbolPays(symbol)
		for count, pay := range payTable {
			var c = r.GetCombinations(symbol, count+1)
			r.AccumulatePays(symbol, count+1, int64(c) * int64(pay))
		}
	}

	fmt.Println(symbolView)
	r.PrintTotals()
	fmt.Println("Combinations")
	r.PrintCombinations()
	fmt.Println("Pays")
	r.PrintPays()
}
