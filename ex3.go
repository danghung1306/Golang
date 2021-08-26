package main

import (
	f "fmt"
	"os"
)

var num int = 123211687556456

func main() {
	if num < 0 {
		f.Println("The number is not Palindrome!")
		os.Exit(0)
	} else {
		//do nothing
	}
	mySlice := []int{}
	for num > 0 {
		tmp := num % 10
		num = num / 10
		mySlice = append(mySlice, tmp)
	}

	if checkPalindrome(mySlice) {
		f.Println("The number is Palindrome!")
	} else {
		f.Println("The number is not Palindrome!")
	}

	f.Println(mySlice)

}

func checkPalindrome(msl []int) (flag bool) {
	flag = true
	for i := 0; i < (len(msl) / 2); i++ {
		if msl[i] == msl[len(msl)-i-1] {
			continue
		} else {
			flag = false
		}

	}
	return
}
