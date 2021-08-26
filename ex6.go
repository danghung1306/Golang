package main

import (
	"container/list"
	f "fmt"
	"os"
)

func main() {
	myLinkList1 := list.New().Init()
	myLinkList2 := list.New().Init()
	sumList := list.New().Init()
	sum, rem := 0, 0

	//add element for link list 1
	myLinkList1.PushBack(2)
	myLinkList1.PushBack(4)
	myLinkList1.PushBack(3)

	//add element for link list 2
	myLinkList2.PushBack(6)
	myLinkList2.PushBack(6)
	myLinkList2.PushBack(4)

	for i := 0; i < 3; i++ {
		sum = popFront(myLinkList1) + popFront(myLinkList2) + rem
		rem = 0
		if sum >= 10 {
			sum = sum - 10
			rem = 1
		}
		sumList.PushBack(sum)
	}
	if rem != 0 {
		sumList.PushBack(rem)
	} else {
		//do nothing
	}

	f.Println("Result of sum linker list:")
	for l := sumList.Front(); l != nil; l = l.Next() {
		f.Printf("%d->", l.Value)
	}

}

func popFront(ll *list.List) (num int) {
	if ll != nil {
		num = ll.Front().Value.(int)
		ll.Remove(ll.Front())
	} else {
		f.Println("Linker list is empty!")
		os.Exit(0)
	}
	return
}
