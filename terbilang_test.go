package terbilang

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDigitToUnit(t *testing.T) {
	result := digitToUnit(4)
	fmt.Println(result)
	assert.True(t, result == "ribu")
}

func TestConvertNumber(t *testing.T) {
	result := convertNumber(12344321)
	fmt.Println(result)
	assert.True(t, result == "duabelas juta tiga ratus empat puluh empat ribu tiga ratus dua puluh satu")
}

func TestFromInt(t *testing.T) {
	result := FromInt{Value: 2000}.ToWord()
	fmt.Println(result)
	assert.True(t, result == "dua ribu")
}

func TestFromFloat(t *testing.T) {
	result := FromFloat{Value: 2150000.00001}.ToWord()
	fmt.Println(result)
	assert.True(t, result == "dua juta seratus lima puluh ribu koma nol nol nol nol satu nol")
}

func TestFromString1(t *testing.T) {
	result := FromString{"15150"}.ToWord()
	fmt.Println(result)
	assert.True(t, result == "limabelas ribu seratus lima puluh")
}

func TestFromString2(t *testing.T) {
	result := FromString{"31689"}.ToWord()
	fmt.Println(result)
	assert.True(t, result == "tiga puluh satu ribu enam ratus delapan puluh sembilan")
}
