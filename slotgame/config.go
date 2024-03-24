package slotgame

import (
	"sort"
)

// Config contains the configuration for a slot server
type Config struct {
	symbols   map[string]symbol
	symbolIds []string
	scatters  []string
	reelSets  map[string]reelSet
}

type symbol struct {
	id         string
	value      string
	pays       []int
	isWild     bool
	multiplier int
	isScatter  bool
	isWay      bool
}

type reelSet struct {
	rows        []int
	reels       [][]string
	reelDisplay [][][]string
	symbolCount [][]map[string]int
}

// GetConfig returns the configuration for a slot server
func GetConfig() Config {
	var config = Config{}
	config.symbols = make(map[string]symbol)
	config.scatters = make([]string, 0)
	config.reelSets = make(map[string]reelSet)

	config.symbols["Z"] = makeSymbol("Z", "Z", []int{0, 0, 0, 0, 0})
	config.symbols["B"] = makeSymbol("B", "B", []int{0, 0, 100, 200, 300})
	config.symbols["C"] = makeSymbol("C", "C", []int{0, 0, 50, 150, 250})
	config.symbols["D"] = makeSymbol("D", "D", []int{0, 0, 30, 100, 200})
	config.symbols["E"] = makeSymbol("E", "E", []int{0, 0, 20, 50, 150})
	config.symbols["A"] = makeSymbol("A", "A", []int{0, 0, 20, 50, 150})
	config.symbols["K"] = makeSymbol("K", "K", []int{0, 0, 15, 30, 125})
	config.symbols["Q"] = makeSymbol("Q", "Q", []int{0, 0, 10, 20, 100})
	config.symbols["J"] = makeSymbol("J", "J", []int{0, 0, 10, 20, 100})
	config.symbols["T"] = makeSymbol("T", "T", []int{0, 0, 5, 15, 100})
	config.symbols["N"] = makeSymbol("N", "N", []int{0, 0, 5, 15, 100})
	config.symbols["S"] = makeSymbol("S", "S", []int{0, 0, 2, 10, 50})
	makeWild(config, "Z", 1)
	config = makeScatter(config, "S")

	config.symbolIds = make([]string, 0, len(config.symbols))
	for id := range config.symbols {
		config.symbolIds = append(config.symbolIds, id)
	}
	sort.Strings(config.symbolIds)

	var baseReelSet = reelSet{}
	baseReelSet.rows = []int{3, 5, 5, 5, 3}
	baseReelSet.reels = make([][]string, 5)
	baseReelSet.reels[0] = []string{"K", "J", "E", "T", "Q", "E", "K", "T", "E", "J", "K", "B", "J", "T", "E", "J", "K", "B", "J", "T", "J", "S", "K", "T", "J", "E", "Q", "J", "C", "K", "J", "Q", "E", "J", "K", "T", "Q", "D", "A", "T", "K", "E", "N", "K", "T", "E", "J", "K", "T", "E", "J", "T", "D", "K", "T", "E", "K", "T", "D", "B"}
	baseReelSet.reels[1] = []string{"N", "J", "Q", "Z", "N", "T", "A", "Q", "Z", "N", "J", "Q", "A", "C", "S", "A", "C", "Q", "N", "N", "A", "C", "B", "B", "B", "A", "N", "C", "Q", "Q", "K", "A", "C", "J", "Q", "J", "N", "C", "A", "K", "C", "A", "Q", "Q", "C", "D", "N", "A", "N", "C", "C", "J", "A", "J", "K", "K", "C", "N", "B", "B", "B", "Q", "C", "E", "A"}
	baseReelSet.reels[2] = []string{"N", "T", "T", "A", "B", "B", "B", "B", "B", "A", "J", "K", "J", "E", "T", "N", "E", "A", "K", "T", "D", "Q", "N", "J", "E", "A", "A", "D", "K", "Q", "J", "T", "N", "N", "K", "Q", "B", "B", "B", "B", "B", "N", "J", "T", "T", "E", "A", "A", "A", "N", "K", "K", "E", "E", "Q", "T", "A", "S", "Q", "D", "J", "K", "Q", "S", "E", "A", "N", "N", "Q", "T", "T", "C", "J", "J", "E", "E", "E", "J", "J", "D", "D", "Q", "Q", "K", "N"}
	baseReelSet.reels[3] = []string{"N", "T", "J", "Q", "Z", "N", "C", "J", "Q", "Z", "N", "T", "J", "Q", "D", "A", "A", "B", "B", "B", "B", "B", "N", "N", "C", "C", "C", "D", "Q", "J", "J", "E", "J", "S", "D", "N", "C", "C", "A", "Q", "Q", "A", "T", "K", "B", "B", "B", "B", "B", "D", "J", "A", "N", "N", "T", "D", "D", "D", "T", "T", "E", "Q", "D", "Q", "C", "J", "N", "N", "C", "C", "T", "C", "T", "Q", "T", "T", "D", "N", "S", "J", "J", "C", "Q", "D", "D"}
	baseReelSet.reels[4] = []string{"N", "Q", "B", "B", "B", "K", "N", "C", "N", "B", "B", "B", "K", "D", "A", "C", "N", "S", "K", "A", "Q", "C", "N", "E", "T", "K", "A", "C", "C", "N", "A", "A", "K", "C", "T", "N", "T", "J", "C", "C", "N", "T", "B", "B", "B", "C", "K", "N", "J", "A", "N", "C", "A", "D", "D"}
	baseReelSet.reelDisplay = makeReelDisplay(baseReelSet)
	baseReelSet.symbolCount = makeSymbolCount(baseReelSet.reelDisplay)
	config.reelSets["reels"] = baseReelSet

	var freeReelSet = reelSet{}
	freeReelSet.rows = []int{3, 5, 5, 5, 3}
	freeReelSet.reels = make([][]string, 5)
	freeReelSet.reels[0] = []string{"K", "J", "C", "D", "K", "T", "E", "Q", "D", "K", "J", "C", "T", "J", "C", "Q", "J", "S", "Q", "T", "J", "E", "Q", "T", "C", "K", "T", "C", "K", "J", "D", "C", "Q", "E", "T", "A", "D", "N", "J", "E", "T", "J", "B", "Q", "C", "T", "E", "D", "T", "E", "A", "T", "E", "K", "J", "D", "K", "Q", "E", "B"}
	freeReelSet.reels[1] = []string{"N", "T", "J", "Q", "Z", "Z", "Z", "Z", "Z", "Z", "N", "T", "A", "E", "N", "T", "Q", "A", "N", "A", "S", "A", "C", "E", "N", "C", "T", "B", "Q", "A", "N", "E", "K", "Q", "N", "D", "A", "Q", "E", "K", "N", "A", "Z", "Z", "Z", "Z", "Z", "Z", "K", "N", "N", "Q", "A", "A", "E", "D", "K", "N", "E", "K", "B", "B", "N", "B", "A"}
	freeReelSet.reels[2] = []string{"N", "B", "T", "B", "B", "A", "A", "K", "E", "J", "J", "T", "T", "N", "E", "A", "J", "K", "T", "D", "Q", "N", "J", "E", "A", "D", "K", "C", "J", "T", "N", "K", "Q", "N", "N", "K", "J", "T", "E", "K", "A", "B", "K", "B", "A", "T", "K", "A", "N", "K", "A", "Q", "D", "T", "A", "S", "Q", "N", "D", "J", "C", "Q", "S", "D", "A", "N", "T", "Q", "D", "T", "J", "C", "Q", "J", "D", "D", "Q", "Q", "E", "J", "Q", "D", "N", "J", "D"}
	freeReelSet.reels[3] = []string{"N", "T", "J", "Q", "Z", "Z", "Z", "Z", "Z", "Z", "N", "T", "N", "C", "E", "A", "B", "B", "B", "K", "B", "N", "C", "J", "T", "D", "N", "J", "E", "N", "J", "S", "K", "A", "Q", "N", "D", "A", "Q", "A", "E", "T", "K", "J", "E", "Q", "K", "D", "C", "Z", "Z", "Z", "Z", "Z", "Z", "A", "D", "A", "D", "Q", "E", "Q", "J", "E", "K", "T", "D", "N", "D", "T", "K", "C", "J", "T", "B", "A", "B", "J", "T", "E", "S", "K", "C", "Q", "E"}
	freeReelSet.reels[4] = []string{"N", "Q", "B", "B", "B", "N", "B", "B", "C", "N", "J", "D", "A", "T", "S", "Q", "J", "N", "E", "Q", "K", "E", "Q", "N", "E", "T", "Q", "B", "J", "B", "A", "Q", "E", "N", "B", "Q", "A", "N", "D", "T", "N", "Q", "D", "N", "T", "B", "Q", "J", "E", "A", "N", "B", "T", "Q", "B"}
	freeReelSet.reelDisplay = makeReelDisplay(freeReelSet)
	freeReelSet.symbolCount = makeSymbolCount(freeReelSet.reelDisplay)
	config.reelSets["freeReels"] = freeReelSet

	return config
}

func makeSymbol(id string, value string, pays []int) symbol {
	var newSymbol = symbol{id: id, value: value, pays: pays}
	newSymbol.isWild = false
	newSymbol.multiplier = 1
	newSymbol.isScatter = false
	newSymbol.isWay = true
	return newSymbol
}

func makeWild(c Config, i string, m int) {
	var s = c.symbols[i]
	s.isWild = true
	s.multiplier = m
	s.isWay = false
	c.symbols[i] = s
}

func makeScatter(c Config, i string) Config {
	var s = c.symbols[i]
	s.isScatter = true
	s.isWay = false
	c.symbols[i] = s
	c.scatters = append(c.scatters, i)
	return c
}

func makeReelDisplay(r reelSet) [][][]string {
	var reelDisplay = make([][][]string, len(r.rows))
	for index, reel := range r.reels {
		reelDisplay[index] = make([][]string, len(reel))
		doubleReel := append(reel, reel...)
		for i := 0; i < len(reel); i++ {
			reelDisplay[index][i] = append(doubleReel[i : i+r.rows[index]])
		}
	}
	return reelDisplay
}

func makeSymbolCount(r [][][]string) [][]map[string]int {
	var symbolCount = make([][]map[string]int, len(r))
	for rIndex, row := range r {
		symbolCount[rIndex] = make([]map[string]int, len(r[rIndex]))
		for sIndex, symbols := range row {
			symbolCount[rIndex][sIndex] = make(map[string]int)
			for _, s := range symbols {
				symbolCount[rIndex][sIndex][s]++
			}
		}
	}
	return symbolCount
}

// Reels returns the number of reels in the slot game
func (c *Config) Reels() int {
	return 5
}

// Rows returns the number of rows for each reel in the slot game
func (c *Config) Rows(key string) []int {
	return c.reelSets[key].rows
}

// MinimumBet returns the minimum bet per spin in the slot game
func (c *Config) MinimumBet() int {
	return 10
}

// ReelSet returns a set of reels in the slot game
// key is the identifier of the reel set
func (c *Config) ReelSet(key string) [][]string {
	return c.reelSets[key].reels
}

// SymbolID returns the symbol id of a given symbol in the slot game
// key is the identifier of the symbol
func (c *Config) SymbolID(key string) string {
	return c.symbols[key].id
}

// SymbolValue returns the symbol value of a given symbol in the slot game
// key is the identifier of the symbol
func (c *Config) SymbolValue(key string) string {
	return c.symbols[key].value
}

// SymbolPays returns the payout of a given symbol in the slot game
// key is the identifier of the symbol
func (c *Config) SymbolPays(key string) []int {
	return c.symbols[key].pays
}

// SymbolIsWild returns whether or not the given symbol is a wild (substitute)
// key is the identifier of the symbol
func (c *Config) SymbolIsWild(key string) bool {
	return c.symbols[key].isWild
}

// SymbolMultiplier returns the multiplier of a given symbol in the slot game
// key is the identifier of the symbol
func (c *Config) SymbolMultiplier(key string) int {
	return c.symbols[key].multiplier
}

// SymbolIsScatter returns whether or not the given symbol is a scatter
// key is the identifier of the symbol
func (c *Config) SymbolIsScatter(key string) bool {
	return c.symbols[key].isScatter
}

// SymbolIsWay returns whether or not the given symbol is a way symbol (pays by ways)
// key is the identifier of the symbol
func (c *Config) SymbolIsWay(key string) bool {
	return c.symbols[key].isWay
}

// Scatters returns the identifer of all scatter symbols in the slot game
func (c *Config) Scatters() []string {
	return c.scatters
}

// Symbols returns the identifers of all symbols in the slot game
func (c *Config) Symbols() []string {
	return c.symbolIds
}

// SymbolCounter returns a map to count symbols
func (c *Config) SymbolCounter() map[string]int {
	var counter = make(map[string]int)
	for _, symbol := range c.symbolIds {
		counter[symbol] = 0
	}
	return counter
}

// GetReelView returns the reel view for a given set of reel stops and a given reel set
func (c *Config) GetReelView(stops []int, reelSet string) ([][]string, []map[string]int) {
	var view = make([][]string, c.Reels())
	var count = make([]map[string]int, c.Reels())
	for x, y := range stops {
		view[x] = c.reelSets[reelSet].reelDisplay[x][y]
		count[x] = c.reelSets[reelSet].symbolCount[x][y]
	}
	return view, count
}
