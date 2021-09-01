package main

import (
	"bufio"
	f "fmt"
	"os"
	"regexp"
	"strings"
	"sync"
)

const numberPattern string = "^(\\d+)$"
const stringPattern string = "^(\"[a-zA-Z]+\")$"

func main() {
	var wg = sync.WaitGroup{}
	wg.Add(1)
	buffer := inputString()

	ch := make(chan string, len(buffer))

	go func(ch <-chan string) {
		for i := range ch {
			f.Printf("\nWarning – could not determine data type of: %s", i)
		}
		wg.Done()
	}(ch)

	for _, j := range buffer {
		if flag, _ := regexp.MatchString(numberPattern, j); flag {
			f.Printf("\ntype of: %s is int", j)
		} else if strFlag, _ := regexp.MatchString(stringPattern, j); strFlag {
			f.Printf("\ntype of: %s is string", j)
		} else {
			ch <- j
		}
	}
	wg.Done()
	wg.Wait()
}

func inputString() (str []string) {
	f.Print("\nInput your string: ")
	reader := bufio.NewReader(os.Stdin)

	buffer, _ := reader.ReadString('\n')
	buffer = strings.Replace(buffer, "\n", "", -1)
	str = strings.Fields(buffer)
	return
}
