package terbilang

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func convertNumber(someNumber int) string {
	var response string = ""
	var printUnit bool = true
	var isBelasan bool = false

	return response
}

func convertNumberafterComma(number int) string {

	letterOfNumber := strconv.Itoa(number)
	arrayOfLetter := strings.Split(letterOfNumber, "")
	arrayOfWord := []string{}

	for _, i := range arrayOfLetter {
		j, _ := strconv.Atoi(i)
		arrayOfWord = append(arrayOfWord, numberToText(j))
	}

	return strings.Join(arrayOfWord, " ")
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

func numberToText(index int) string {
	numbers := []string{"satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan"}
	return numbers[index-1]
}

type Setting struct {
	decimal string
}

func ToWord(targetNumber float64, setting Setting) string {
	var result string = ""

	if setting == (Setting{}) {
		setting.decimal = "."
	}

	strOfTargetNumber := strconv.FormatFloat(targetNumber, 'f', 6, 64)
	numberComponent := strings.Split(strOfTargetNumber, setting.decimal)

	majorSegment, errStrConvert := strconv.Atoi(numberComponent[0])
	if errStrConvert != nil {
		panic(errStrConvert)
	}

	result = convertNumber(majorSegment)

	if len(numberComponent) == 2 {
		afterCommaSegment, errStrConvert := strconv.Atoi(numberComponent[1])
		if errStrConvert != nil {
			panic(errStrConvert)
		}
		result = fmt.Sprintf("%s koma %s", result, convertNumberafterComma(afterCommaSegment))
	}

	return result
}
