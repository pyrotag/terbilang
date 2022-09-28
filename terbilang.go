package terbilang

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func convertNumber(someNumber int) string {
	var result string = ""
	var printUnit bool = true
	var isBelasan bool = false

	stringOfNumber := strconv.Itoa(someNumber)
	arrStrNum := strings.Split(stringOfNumber, "")

	for i := 0; i < len(arrStrNum); i++ {
		length := len(arrStrNum) - 1 - i

		if (length % 3) == 0 {
			num := ""

			isMinus2 := false
			isMinus1 := false

			if i-2 >= 0 {
				if arrStrNum[i-2] == "" || arrStrNum[i-2] == "0" {
					isMinus2 = true
				}
			}

			if i-1 >= 0 {
				if arrStrNum[i-1] == "" || arrStrNum[i-1] == "0" {
					isMinus1 = true
				}
			}

			if arrStrNum[i] == "1" && (isBelasan || (digitToUnit(length) == "ribu" && (isMinus2 && isMinus1))) {
				num = "se"
			} else {
				targetNum, _ := strconv.Atoi(arrStrNum[i])
				num = numberToText(targetNum)
			}

			result = fmt.Sprintf("%s %s", result, num)

			if (i-1 >= 0) && (i-2 >= 0) {
				if arrStrNum[i-2] != "0" || arrStrNum[i-1] != "0" || arrStrNum[i] != "0" {
					printUnit = true
				}
			}

			if printUnit == true {
				printUnit = false
				addBelas := ""
				if isBelasan == true {
					addBelas = "belas"
					isBelasan = false
				}
				result = fmt.Sprintf("%s%s %s", result, addBelas, digitToUnit(length))
			}

		} else if (length%3) == 2 && arrStrNum[i] != "0" {
			num := ""
			if arrStrNum[i] == "1" {
				num = "seratus"
			} else {
				targetNum, _ := strconv.Atoi(arrStrNum[i])
				num = fmt.Sprintf("%s ratus", numberToText(targetNum))
			}
			result = fmt.Sprintf("%s %s", result, num)
		} else if (length%3) == 1 && arrStrNum[i] != "0" {
			if arrStrNum[i] == "1" {
				if arrStrNum[i+1] == "0" {
					result = fmt.Sprintf("%s %s", result, "sepuluh")
				} else {
					isBelasan = true
				}
			} else {
				targetNum, _ := strconv.Atoi(arrStrNum[i])
				result = fmt.Sprintf("%s %s puluh", result, numberToText(targetNum))
			}
		}
	}

	result = strings.Trim(result, " ")
	reg := regexp.MustCompile(`\s+`)
	result = reg.ReplaceAllString(result, " ")

	return result
}

func convertNumberafterComma(strNumber string) string {

	arrayOfLetter := strings.Split(strNumber, "")
	arrayOfWord := []string{}

	if len(arrayOfLetter) >= 1 {
		arrayOfWord = append(arrayOfWord, "koma")
		for _, i := range arrayOfLetter {
			j, _ := strconv.Atoi(i)
			if j == 0 {
				arrayOfWord = append(arrayOfWord, "nol")
			} else {
				arrayOfWord = append(arrayOfWord, numberToText(j))
			}
		}
	}

	return strings.Join(arrayOfWord, " ")
}

func digitToUnit(digit int) string {
	var returnUnit string = ""
	units := []string{"", "ribu", "juta", "milyar", "triliun", "quadriliun", "quintiliun", "sextiliun", "septiliun", "oktiliun", "noniliun", "desiliun", "undesiliun", "duodesiliun", "tredesiliun", "quattuordesiliun", "quindesiliun", "sexdesiliun", "septendesiliun", "oktodesiliun", "novemdesiliun", "vigintiliun"}

	curIndex := math.Floor(float64(digit) / 3)
	maxIndex := len(units) - 1

	returnUnit = units[int(curIndex)]

	if curIndex > float64(maxIndex) {
		returnUnit = units[int(maxIndex)]
	}

	return returnUnit
}

func numberToText(index int) string {
	numbers := []string{"satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan"}
	if index-1 < 0 {
		return ""
	}
	return numbers[index-1]
}

func stringNumToWord(stringNum string) string {
	var result string = ""

	numberComponent := strings.Split(stringNum, ".")
	majorSegment, errStrConvert := strconv.Atoi(numberComponent[0])
	if errStrConvert != nil {
		panic(errStrConvert)
	}

	result = convertNumber(majorSegment)

	if len(numberComponent) == 2 {
		result = fmt.Sprintf("%s %s", result, convertNumberafterComma(numberComponent[1]))
	}

	return result
}

type FromFloat struct {
	Value float64
}

func (val FromFloat) ToWord() string {
	strOfTargetNumber := strconv.FormatFloat(val.Value, 'f', 6, 64)
	return stringNumToWord(strOfTargetNumber)
}

type FromString struct {
	Value string
}

func (val FromString) ToWord() string {
	return stringNumToWord(val.Value)
}

type FromInt struct {
	Value int
}

func (val FromInt) ToWord() string {
	strOfTargetNumber := strconv.Itoa(val.Value)
	return stringNumToWord(strOfTargetNumber)
}
