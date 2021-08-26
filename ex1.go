package main

import f "fmt"

var number int
var target int

func main() {
	f.Print("Input the number element of array:")
	f.Scanf("%d", &number)
	arr := make([]int, number, number)
	for i := 0; i < number; i++ {
		f.Printf("Number[%d]: ", i)
		f.Scanf("\n%d", &arr[i])
	}
	f.Print("Input the Target: ")
	f.Scanf("\n%d", &target)

	if matches(target, number, arr) != nil {
		f.Println("Position matched: ", matches(target, number, arr))
	} else {
		f.Print("No number matched!\n")
	}
}

func matches(tar int, length int, arr []int) (result []int) {
	result = nil
	for i := 0; i < length-1; i++ {
		if arr[i] == tar {
			result = append(result, i)
			return
		} else {
			for j := i + 1; j < length; j++ {
				if (arr[i] + arr[j]) == tar {
					result = append(result, i, j)
					return
				}
			}
		}
	}
	return
}
