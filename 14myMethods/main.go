package main

import "fmt"

type User struct {
	Name	string
	Email	string
	Status	bool
	Age		int
}

// method to get status
func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status);
}

// method to set new mail
func (u User) SetNewMail(mail string) {
	u.Email=mail;
	fmt.Println("new email:", u.Email);
}

func main() {
	fmt.Println("structs in go");

	// no inheritance in golang; no concept of parent 

	user1:=User{"abc", "abc@gmail.com", true, 10};
	fmt.Println("user1:", user1);
	fmt.Printf("%+v\n", user1);
	user1.GetStatus();

	user1.SetNewMail("def@gmail.com");
}