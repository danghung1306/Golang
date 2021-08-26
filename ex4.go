package main

import f "fmt"

var romanNum string = "MCMXCIV"
var interNum int = 0
var currentNum int = 0
var prevNum int = 0

func main() {
	romanMap := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	for i := len(romanNum) - 1; i >= 0; i-- {
		currentNum = romanMap[string(romanNum[i])]
		if currentNum >= prevNum {
			interNum = interNum + currentNum
		} else {
			interNum = interNum - currentNum
		}
		prevNum = currentNum
	}

	f.Println("The number convert: ", interNum)
}
