package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Arrays in Golang")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Mango"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList)) // but the length will be 4 as we have declared the array size as 4
	var vegList = [3]string{"Potato", "Tomato", "Cabbage"}
	fmt.Println(vegList)
	fmt.Println(len(vegList))
}
