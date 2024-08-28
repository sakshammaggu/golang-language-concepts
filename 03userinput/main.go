package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Take User input..");
	userInput:=bufio.NewReader(os.Stdin);

	// commam ok, error err
	input, _ :=userInput.ReadString('\n');
	fmt.Println("user input:", input);
	fmt.Printf("Type of input is: %T", input);
}