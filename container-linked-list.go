package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack(1)
	data.PushBack(2)
	data.PushBack(3)
	data.PushFront(0)

	//data.Front().Next().Next()
	//data.Front().Prev()

	//for element := data.Front(); element != nil; element = element.Next() {
	//	fmt.Println(element.Value)
	//}

	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)
}
