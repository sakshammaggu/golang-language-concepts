package main

import "fmt"

func main() {
	defer fmt.Println("hello world");
	fmt.Println("defer in go");

	// defer statements are printed in reverse order -> LIFO
	defer fmt.Println("one");
	defer fmt.Println("two");
	defer fmt.Println("three");
}