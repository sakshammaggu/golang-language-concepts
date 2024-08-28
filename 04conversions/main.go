package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to pizza app.");
	fmt.Println("Please rate between 1 and 5 for pizza.");

	reader := bufio.NewReader(os.Stdin);
	input, _ := reader.ReadString('\n');

	// Trim the newline character from the input
	input = strings.TrimSpace(input);
	pizzaRating, err := strconv.ParseFloat(input, 64);

	if (err != nil) {
		fmt.Println("Error:", err);
	} else {
		fmt.Println("Rating:", pizzaRating);
		fmt.Printf("Type of rating: %T \n", pizzaRating);
	}

	fmt.Println("Thanks for rating.");
}
