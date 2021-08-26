package main

import f "fmt"

var num1 int32 = 120
var num2 int32 = 0

func main() {
	if num1 > 0 {
		for num1 > 0 {
			num2 = num2*10 + num1%10
			num1 = num1 / 10
		}
	} else {
		for num1 < 0 {
			num2 = num2*10 + num1%10
			num1 = num1 / 10
		}
	}
	f.Println(num2)
}
