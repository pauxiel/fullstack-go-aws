package main

import (
	"fmt"
	"slices"
)

// func main() {
// 	animals := [2]string{}
//      animals[0] = "dog"
// 	   animals[1] = "cat"

// 		 fmt.Println(animals)

// }

func main() {
	animals := []string{
		"dog",
		"cat",
	}

	animals = append(animals, "moose")
		 fmt.Println(animals)

		 animals = slices.Delete(animals, 0, 1)

		 for i:= 0; i < len(animals); i++ {
			fmt.Printf("This is animal %s\n", animals[i])
		 }

	
}