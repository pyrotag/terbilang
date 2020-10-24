package terbilang

import "math"

func toWord(someNumber float64) string {
	var response string = ""
	var printUnit bool = true
	var isBelasan bool = false

	return response
}

func digitToUnit(digit float64) string {
	var returnUnit string = ""
	units := []string{"", "ribu", "juta", "milyar", "triliun", "quadriliun", "quintiliun", "sextiliun", "septiliun", "oktiliun", "noniliun", "desiliun", "undesiliun", "duodesiliun", "tredesiliun", "quattuordesiliun", "quindesiliun", "sexdesiliun", "septendesiliun", "oktodesiliun", "novemdesiliun", "vigintiliun"}

	curIndex := math.Floor(digit / 3)
	maxIndex := len(units) - 1

	returnUnit = units[int(curIndex)]

	if curIndex > float64(maxIndex) {
		returnUnit = units[int(maxIndex)]
	}

	return returnUnit

}
