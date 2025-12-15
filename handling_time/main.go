package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("welcome to time study of go lang ")

	presentTime := time.Now()
	fmt.Println("The present time is ", presentTime)
		fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))  // standard in doc 

		createdDate := time.Date(2020, time.January,26,12,10,2,1,time.UTC)
		fmt.Println("The created date is ", createdDate)
		fmt.Println(createdDate.Format("01-02-2006 Monday"))

}