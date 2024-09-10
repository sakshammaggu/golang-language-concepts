package main

import "fmt"

type User struct {
	Name	string
	Email	string
	Status	bool
	Age		int
}

func main() {
	fmt.Println("structs in go");

	// no inheritance in golang; no concept of parent 

	user1:=User{"abc", "abc@gmail.com", true, 10};
	fmt.Println("user1:", user1);
	fmt.Printf("%+v\n", user1);
}