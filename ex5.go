package main

import f "fmt"

func main() {
	array1 := []string{"aflower", "bflow", "cfloight", "dfloag", "efloy"}
	pos := findShortestWord(array1)
	lenOfShortestWord := len(array1[pos])
	f.Println("Position of Shortest Word: ", pos)
	f.Println("Length of Shortest Word  : ", lenOfShortestWord)

	num := checkCommonWord(array1, lenOfShortestWord)

	if num != 0 {
		f.Println("Common word:", array1[0][0:num-1])
	} else {
		f.Println("There is no common prefix among the input strings.")
	}

}

func findShortestWord(arr []string) (posShortestWord int) {
	posShortestWord = 0
	for i := 0; i < len(arr)-1; i++ {
		if len(arr[i+1]) < len(arr[i]) {
			posShortestWord = i + 1
		} else {
			posShortestWord = i
		}
	}
	return
}
func checkCommonWord(arr []string, lenWord int) (num int) {
	num = 0
	for i := lenWord; i > 0; i-- {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j][0:i] == arr[j+1][0:i] {
				num = i
				continue
			} else {
				num = 0
				break
			}
		}
		if num != 0 {
			break
		} else {
			//do nothing
		}
	}
	return
}
