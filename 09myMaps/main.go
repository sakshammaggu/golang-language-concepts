package main

import "fmt"

func main() {
	fmt.Println("maps in golang");

	lang:=make(map[string]string);

	// adding values
	lang["Js"]="Javascript";
	lang["Go"]="Golang";
	lang["Py"]="Python";
	lang["Cpp"]="C++";

	fmt.Println("List:", lang);
	fmt.Printf("Type is: %T\n", lang);

	// deleting values
	delete(lang, "Py");

	fmt.Println("List:", lang);
	fmt.Printf("Type is: %T \n", lang);

	// accessing elements through loop
	for key, value:=range lang {
		fmt.Printf("For key %v, value is %v \n", key, value);
	}
}