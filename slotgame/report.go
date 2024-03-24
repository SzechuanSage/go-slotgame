package slotgame

import (
	"fmt"
)

// Report contains the fields for a slot game test
type Report struct {
	symbols      []string
	totals       map[string]int32
	combinations map[string][]int32
	pays         map[string][]int64
	baseWins     map[string]int32
}

// InitReport returns a report that has been initialised with symbols and reels
func InitReport(symbols []string, reels int) Report {
	var report = Report{}

	report.totals = make(map[string]int32)
	report.totals["count"] = 0
	report.totals["hits"] = 0
	report.totals["bet"] = 0
	report.totals["win"] = 0

	report.combinations = make(map[string][]int32)
	report.pays = make(map[string][]int64)
	report.symbols = append(report.symbols, symbols...)

	for _, symbol := range symbols {
		report.combinations[symbol] = make([]int32, reels+1)
		report.pays[symbol] = make([]int64, reels+1)
	}

	report.baseWins = make(map[string]int32)
	return report
}

// AccumulateTotal accumulates one of the totals in the report
func (r *Report) AccumulateTotal(key string, amount int32) {
	r.totals[key] += amount
}

// AccumulateCombinations accumulates combinations for a particuar symbol
func (r *Report) AccumulateCombinations(key string, count int, amount int32) {
	r.combinations[key][count] += amount
}

// GetCombinations returns combinations for a particuar symbol
func (r *Report) GetCombinations(key string, count int) int32 {
	return r.combinations[key][count]
}

// AccumulatePays accumulates pays for a particuar combination of symbols
func (r *Report) AccumulatePays(key string, count int, amount int64) {
	r.pays[key][count] += amount
}

// PrintTotals prints the totals section of the report
func (r *Report) PrintTotals() {
	var line = fmt.Sprintf("%-12s   %12d\n", "Count", r.totals["count"])
	fmt.Println(line)
}

// PrintCombinations prints the combinations section of the report
func (r *Report) PrintCombinations() {
	var line = ""
	for _, s := range r.symbols {
		line = fmt.Sprintf("%-4s", s)
		for _, c := range r.combinations[s][1:] {
			line += fmt.Sprintf("%16d", c)
		}
		fmt.Println(line)
	}
}

// PrintPays prints the pays section of the report
func (r *Report) PrintPays() {
	var line = ""
	for _, s := range r.symbols {
		line = fmt.Sprintf("%-4s", s)
		for _, p := range r.pays[s][1:] {
			line += fmt.Sprintf("%16d", p)
		}
		fmt.Println(line)
	}
}
