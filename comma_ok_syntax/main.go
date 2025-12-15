// https://pkg.go.dev

package main

import (
	"fmt" 
	"bufio"
	"os"   )

func main(){
	
	welcome := "welcome to the user input "
	fmt.Print(welcome)
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the rating for the service ")


	// OK COMMA SYNTAX 

			input, _ := reader.ReadString('\n')  // here the _ is used to ignore the error value returned by ReadString function
			fmt.Println("Thanks for rating : ", input)
			
			fmt.Printf("Type of the input is  %T",input)

}