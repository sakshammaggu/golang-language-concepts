package main

import "fmt"

func main() {
	fmt.Println("if else in go");

	score := 85

    if score >= 90 {
        fmt.Println("Grade: A")
    } else if score >= 80 { 
        fmt.Println("Grade: B")
    } else if score >= 70 {
        fmt.Println("Grade: C")
    } else {
        fmt.Println("Grade: D")
    }
}