package main 

import "fmt"

// unparamatrised function
func greeter() {
	fmt.Println("hello functions");
}

// paramaterised function
func adder(x int, y int) int {
	sum:=x+y;
	return sum;
}

// pro functions
func proAdder(values ...int) int {
	sum:=0;
	for _, val:=range values {
		sum+=val;
	}
	return sum;
}

func main() {
	fmt.Println("functions in go");

	greeter();

	sum:=adder(1, 2);
	fmt.Println("Sum:", sum);

	proSum:=proAdder(1, 2, 3);
	fmt.Println("Pro sum:", proSum)
}	