package main

import "fmt"

func main() {
	fmt.Println("slices data structure");

	// basic slice declaring syntax
	var fruitList=[]string{"apple", "banana"};
	fmt.Printf("Type of slice is: %T \n", fruitList);
	fruitList = append(fruitList, "mango");
	fmt.Println("fruitList:", fruitList);
	fruitList = append(fruitList[1:]);
	fmt.Println("fruitList:", fruitList);
	fruitList = append(fruitList[0:2]);
	fmt.Println("fruitList:", fruitList);

	// priotized way of declaring slice
	highScores:=make([]int, 4);
	highScores[0]=0;
	highScores[1]=1;
	highScores[2]=2;
	highScores[3]=3;
	fmt.Println("highScores:", highScores);
	fmt.Println("highScore length:", len(highScores));
	highScores = append(highScores, 4, 5);
	fmt.Println("highScores:", highScores);
	fmt.Println("highScore length:", len(highScores));

	// how to remove a value from a slice based on index
	courses := []string{"Course1", "Course2", "Course3", "Course4"};
	fmt.Println("courses:", courses);
	index := 2; // index of the value you want to remove
	courses = append(courses[:index], courses[index+1:]...);
	fmt.Println("courses:", courses);
}