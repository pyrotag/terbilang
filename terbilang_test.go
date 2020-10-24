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
