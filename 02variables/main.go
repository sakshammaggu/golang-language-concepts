package main

import "fmt"

func main() {
	// string
    var userName string = "saksham";
	fmt.Println(userName);
	fmt.Printf("Variable is of type: %T \n", userName);

	// bool
	var isChecked bool = false;
	fmt.Println(isChecked);
	fmt.Printf("Variable is of type: %T \n", isChecked);

	// uint8
	var smallVal uint8 = 25;
	fmt.Println(smallVal);
	fmt.Printf("Variable is of type: %T \n", smallVal);

	// float32
	var smallFloat float32 = 255.67888888888888888888888888888;
	fmt.Println(smallFloat);
	fmt.Printf("Variable is of type: %T \n", smallFloat);

	// default values and some aliases
	var intVariable int;
	fmt.Println(intVariable);
	fmt.Printf("Variable is of type: %T \n", intVariable);

	var stringVariable string;
	fmt.Println(stringVariable);
	fmt.Printf("Variable is of type: %T \n", stringVariable);

	// implicit type
	var str="abc";
	fmt.Println(str);
	// str=3; ---> gives error 

	// no var style (walrus operator) => only allowed in methods, cant be used outside method
	number:=3000;
	fmt.Println(number);
}