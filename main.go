package main

import (
	"fembasics/imports"
	"fmt"
)

func main() {
	newTicket := imports.Ticket{
		ID:    123,
		Event: "theater",
	}


	fmt.Println(newTicket)
}