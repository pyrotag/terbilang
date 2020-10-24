package terbilang

import (
	"fmt"
	"math"
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
			if arrStrNum[i] == "1" && (isBelasan || (digitToUnit(length) == "ribu" && ((arrStrNum[i-2] == "" || arrStrNum[i-2] == "0") && arrStrNum[i-1] == "" || arrStrNum[i-1] == "0"))) {
				num = "se"
			} else {
				targetNum, _ := strconv.Atoi(arrStrNum[i])
				num = numberToText(targetNum)
			}
			result = fmt.Sprintf("%s %s", result, num)

			if arrStrNum[i-2] != "0" || arrStrNum[i-1] != "0" || arrStrNum[i] != "0" {
				printUnit = true
			}

			if printUnit == true {
				printUnit = false
				addBelas := ""
				if isBelasan == true {
					addBelas = "belas"
					isBelasan = false
				}
				result = fmt.Sprintf("%s %s%s", result, addBelas, digitToUnit(length))
			}

		} else if (length%3) == 2 && arrStrNum[i] != "0" {

		} else if (length%3) == 1 && arrStrNum[i] != "0" {

		}
	}

	//for (let i = 0; i < angka.length; i++) {
	// 	const length = angka.length - 1 - i
	// 	if (length % 3 == 0) {
	// 	  const num = (angka[i] == 1 && (isBelasan || (digitToUnit(length) == 'ribu' && ((angka[i - 2] == undefined || angka[i - 2] == 0) && (angka[i - 1] == undefined || angka[i - 1] == 0))))) ? 'se' : `${numberToText(angka[i])} `
	// 	  result += ` ${num}`

	// 	  if ((angka[i - 2] && angka[i - 2] != 0) || (angka[i - 1] && angka[i - 1] != 0) || angka[i] != 0) {
	// 		printUnit = true
	// 	  }
	// 	  if (printUnit) {
	// 		printUnit = false
	// 		result += ((isBelasan) ? 'belas ' : '') + digitToUnit(length)
	// 		if (isBelasan) {
	// 		  isBelasan = false
	// 		}
	// 	  }
	// 	} else if (length % 3 == 2 && angka[i] != 0) {
	// 	  result += ` ${(angka[i] == 1) ? 'se' : numberToText(angka[i]) + ' '}ratus`
	// 	} else if (length % 3 == 1 && angka[i] != 0) {
	// 	  if (angka[i] == 1) {
	// 		if (angka[i + 1] == 0) {
	// 		  result += ' sepuluh'
	// 		} else {
	// 		  isBelasan = true
	// 		}
	// 	  } else {
	// 		result += ` ${numberToText(angka[i])} puluh`
	// 	  }
	// 	}
	//   }

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
