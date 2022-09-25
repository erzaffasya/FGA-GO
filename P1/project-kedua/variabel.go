package main

import (
	"fmt"
	"reflect"
)

func main()  {
	var name string = "Rohmat"
	var umur int = 20

	fmt.Println(reflect.TypeOf(name))
	fmt.Println(reflect.TypeOf(umur))
}