package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("wlcome to our pizza app ")
	fmt.Println("Please rate our pizza from 1 to 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for your rating", input)
	fmt.Printf("The type is %T", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("Error while converting rating to number", err)
	} else {
		fmt.Println("added 1 to rating ", numRating+1)
		fmt.Printf("The type is %T", numRating)
	}

}
