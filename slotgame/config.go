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
func GetConfig(gameConfig Game) Config {
	var config = Config{}
	config.symbols = make(map[string]symbol)
	config.scatters = make([]string, 0)
	config.reelSets = make(map[string]reelSet)

	for _, symbol := range gameConfig.Symbols {
		config.symbols[symbol.Symbol] = makeSymbol(symbol.Symbol, symbol.Value, symbol.Pays)
	}

	for _, wild := range gameConfig.Wilds {
		makeWild(config, wild.Symbol, wild.Multiplier)
	}

	for _, scatter := range gameConfig.Scatters {
		config = makeScatter(config, scatter.Symbol)
	}

	config.symbolIds = make([]string, 0, len(config.symbols))
	for id := range config.symbols {
		config.symbolIds = append(config.symbolIds, id)
	}
	sort.Strings(config.symbolIds)

	var baseReelSet = reelSet{}
	baseReelSet.rows = []int{3, 3, 3, 3, 3}
	baseReelSet.reels = make([][]string, 5)
	baseReelSet.reels[0] = append(baseReelSet.reels[0], gameConfig.ReelSets[0].Reels[0]...)
	baseReelSet.reels[1] = append(baseReelSet.reels[1], gameConfig.ReelSets[0].Reels[1]...)
	baseReelSet.reels[2] = append(baseReelSet.reels[2], gameConfig.ReelSets[0].Reels[2]...)
	baseReelSet.reels[3] = append(baseReelSet.reels[3], gameConfig.ReelSets[0].Reels[3]...)
	baseReelSet.reels[4] = append(baseReelSet.reels[4], gameConfig.ReelSets[0].Reels[4]...)
	baseReelSet.reelDisplay = makeReelDisplay(baseReelSet)
	baseReelSet.symbolCount = makeSymbolCount(baseReelSet.reelDisplay)
	config.reelSets[gameConfig.ReelSets[0].Name] = baseReelSet

	var freeReelSet = reelSet{}
	freeReelSet.rows = []int{3, 3, 3, 3, 3}
	freeReelSet.reels = make([][]string, 5)
	freeReelSet.reels[0] = append(freeReelSet.reels[0], gameConfig.ReelSets[1].Reels[0]...)
	freeReelSet.reels[1] = append(freeReelSet.reels[1], gameConfig.ReelSets[1].Reels[1]...)
	freeReelSet.reels[2] = append(freeReelSet.reels[2], gameConfig.ReelSets[1].Reels[2]...)
	freeReelSet.reels[3] = append(freeReelSet.reels[3], gameConfig.ReelSets[1].Reels[3]...)
	freeReelSet.reels[4] = append(freeReelSet.reels[4], gameConfig.ReelSets[1].Reels[4]...)
	freeReelSet.reelDisplay = makeReelDisplay(freeReelSet)
	freeReelSet.symbolCount = makeSymbolCount(freeReelSet.reelDisplay)
	config.reelSets[gameConfig.ReelSets[1].Name] = freeReelSet

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
			reelDisplay[index][i] = append(reelDisplay[index][i], doubleReel[i:i+r.rows[index]]...)
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
