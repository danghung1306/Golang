package main

import "fmt"

const strInput string = "abcabacdef"

func main() {

	fmt.Println("Longest Sub String without repeating: ", findLongestSubString(strInput))

}

func findLongestSubString(strInput string) (result string) {
	var lenLongestSubString int = 0
	var subString string = ""
	for indexStrInput := 0; indexStrInput < len(strInput); indexStrInput++ {
		for indexSubString := 0; indexSubString < len(subString); indexSubString++ {
			if string(strInput[indexStrInput]) == string(subString[indexSubString]) {
				//neu có nhiều chuỗi thì chọn result = chuỗi dài nhất
				if len(subString) > lenLongestSubString {
					result = subString
					//cập nhật lại độ dài của chuỗi dài nhất
					lenLongestSubString = len(subString)
				} else {
					//do nothing
				}
				subString = ""
				break
			} else {
				//do nothing
			}
		}
		subString = subString + string(strInput[indexStrInput])
	}

	if len(subString) > lenLongestSubString {
		result = subString
	} else {
		//do nothing
	}
	return
}
