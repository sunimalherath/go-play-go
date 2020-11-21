package datatypes

import (
	"fmt"
	"reflect"
)

func PointerTest() {
	x := 43
	iPointer := &x // iPointer holds the memory location for x (0xc00001c050)

	fmt.Println(reflect.TypeOf(iPointer))
	fmt.Println(iPointer) 	// 0xc00001c050
	fmt.Println(*iPointer)	// 43 - gets the value from where the iPointer is pointing to

	// changing the value of x
	x = 76
	fmt.Println(*iPointer)	// 76
}
