# Terbilang
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pyrotag/terbilang)](https://pkg.go.dev/github.com/pyrotag/terbilang)

Golang library to convert Number to Word in Bahasa/Indonesia.

This repository inspiration is based from https://github.com/develoka/angka-terbilang-js  

## Install

```
go get -u github.com/pyrotag/terbilang
```

## Usage 
```
import (
    terbilang "github.com/pyrotag/terbilang"
)

wordResultFromString := terbilang.FromString{"1999"}.ToWord()
// Result 
// Seribu sembilan ratus sembilan puluh sembilan

wordResultFromInt := terbilang.FromInt{1525000}.ToWord()
// Result 
// Satu juta lima ratus dua puluh lima ribu

wordResultFromInt := terbilang.FromFloat{153.192311}.ToWord()
// Result 
// Seratus lima puluh tiga koma satu sembilan dua tiga satu satu


```
