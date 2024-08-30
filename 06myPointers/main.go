package main

import "fmt"

func main() {
	fmt.Println("My pointers");

	var ptr *int;	// pointer of integer type
	fmt.Println("value of ptr:", ptr);

	myNumber:=23;
	var myPtr=&myNumber;	// pointer with inital value, used & (ampersand)
	fmt.Println("value of ptr:", myPtr);
	fmt.Println("value of ptr:", *myPtr);

}