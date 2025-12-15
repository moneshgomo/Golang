package main

import (
	"fmt"
)

/*
Pointers are the actual memory addresses of variables.
why pointes are used?
1. To share memory between different parts of a program without copying data.
2. To modify variables in functions without returning them.
3. To work with large data structures efficiently by avoiding unnecessary copies.
*/
func main() {

	fmt.Println("Welcome to Pointers")
	// var pointer_1 *int; // declaring a pointer variable;
	
	// fmt.Println("Pointer value before assignment:", pointer_1); // prints: <nil>
	

	myNumber := 23
	var pointer_2 = &myNumber // & ( reference operator & ) is used to get the address of a variable

	// fmt.Println("Address of myNumber variable:", &myNumber)
	fmt.Println("Pointer value after assignment:", pointer_2)
	fmt.Println("Value at pointer_2 address:", *pointer_2) // * ( dereference operator ) is used to get the value at the address stored in the pointer

	*pointer_2 = *pointer_2 * 2 
	fmt.Println("Value of myNumber after modifying through pointer:", myNumber)
}