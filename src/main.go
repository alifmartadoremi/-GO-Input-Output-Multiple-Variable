package main

import(
	"fmt"
)

func main() {
	var firstname, lastname string
	fmt.Print("Insert your first name : ")
	fmt.Scanln(&firstname)
	
	fmt.Print("Insert your last name : ")
	fmt.Scanln(&lastname)

	fmt.Println("Your full name is "+firstname+" "+lastname)
}