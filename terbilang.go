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

	fmt.Println("ARR LENGTH", len(arrStrNum))

	for i := 0; i < len(arrStrNum); i++ {
		length := len(arrStrNum) - 1 - i
		fmt.Println("LENGTH", length)
		if (length % 3) == 0 {
			fmt.Println("CONDITION A")
			num := ""
			if arrStrNum[i] == "1" && (isBelasan || (digitToUnit(length) == "ribu" && ((i-2 >= 0 && arrStrNum[i-2] == "" || arrStrNum[i-2] == "0") && arrStrNum[i-1] == "" || arrStrNum[i-1] == "0"))) {
				num = "se"
			} else {
				targetNum, _ := strconv.Atoi(arrStrNum[i])
				num = numberToText(targetNum)
			}

			fmt.Println("CONDITION A -", result, num)
			result = fmt.Sprintf("%s %s", result, num)
			fmt.Println("CONDITION A - RESULT", result)

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
			fmt.Println("CONDITION B")
			num := ""
			if arrStrNum[i] == "1" {
				num = "seratus"
			} else {
				targetNum, _ := strconv.Atoi(arrStrNum[i])
				num = fmt.Sprintf("%s ratus", numberToText(targetNum))
			}
			result = fmt.Sprintf("%s %s", result, num)
			fmt.Println("CONDITION B - RESULT", result)
		} else if (length%3) == 1 && arrStrNum[i] != "0" {
			fmt.Println("CONDITION C")
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
			fmt.Println("CONDITION C - RESULT", result)
		}
	}

	result = strings.Trim(result, " ")
	reg := regexp.MustCompile(`\s+`)
	result = reg.ReplaceAllString(result, " ")

	return result
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

type Setting struct {
	decimal string
}

func ToWord(targetNumber float64) string {
	var result string = ""

	strOfTargetNumber := strconv.FormatFloat(targetNumber, 'f', 6, 64)
	numberComponent := strings.Split(strOfTargetNumber, ".")

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
