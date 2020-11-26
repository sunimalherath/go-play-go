package main

import (
	"awesomeProject/oop"
	"fmt"
)

func main() {
	//greetings.SayHello()
	//datatypes.FiddleSlice()

	//datatypes.PointerTest()

	johndoe := oop.Person{Name: "John Doe"}
	johndoe.PrintInfo()
	fmt.Println("Changing name to John Smith...")
	johndoe.ChangeName("John Smith")
	johndoe.PrintInfo()

	janedoe := oop.Person{Name: "Jane Doe"}
	janedoe.PrintInfo()
	fmt.Println("Changing name to Jane Smith")
	janedoe.ChangeNameNoPointer("Jane Smith")
	janedoe.PrintInfo()
}
