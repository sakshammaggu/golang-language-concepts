package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to myTime.");

	presentTime:=time.Now();
	fmt.Println("Present time:", presentTime);

	formatPresentTime:=presentTime.Format("01-02-2006");
	fmt.Println("Format Present time:", formatPresentTime);

	customCreatedDate:=time.Date(2022, time.January, 01, 23, 23, 0, 0, time.UTC);
	fmt.Println("Custom Present Date:", customCreatedDate.Format("01-02-2006 Monday"));
}