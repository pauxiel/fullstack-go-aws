package main

import "fmt"

func main() {
	myName := "Paul"
	myInt := 10
	myFloat := 10.0


	fmt.Printf("Hello my name is %s my int is %d my float is %f\n", myName, myInt, myFloat)

	var myFriendsName string
	var myBool bool
	var myOtherInt int

	myFriendsName = "Prime"

	fmt.Printf("my friends name is %s my bool is %t and my other int is %d\n", myFriendsName, myBool, myOtherInt)
}