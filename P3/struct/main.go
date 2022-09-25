package main

import "fmt"

type Person struct {
	name string
	age  int
}

// embeded
type Employee struct {
	name     string
	age      int
	division string
	alamat   Alamat
}

type Alamat struct {
	street   string
	city     string
	province string
}

func main() {
	// penulisan satu
	// var employee Employee
	// employee.name = "Airell"
	// employee.age = 23
	// employee.division = "Curiculu Developer"
	// employee.alamat.street = "OKE"

	// penulisan kedua
	// var employee2 = Employee{}
	// employee2.name = "Jaini"
	// employee2.age = 23
	// employee2.division = "Curiculu Developer"
	// employee2.alamat.street = "Cendana"
	// employee2.alamat.city = "SBY"
	// employee2.alamat.province = "JATIM"

	// // penulisan ketiga
	// var employee3 = Employee{
	// 	name:     "Oke",
	// 	age:      21,
	// 	division: "Y",
	// 	alamat: Alamat{
	// 		street:   "Cendana",
	// 		city:     "sby",
	// 		province: "Jatim",
	// 	},
	// }

	// slice of struct
	var people = []Person{
		{name: "Airell", age: 23},
		{name: "Ananda", age: 23},
		{name: "Mailo", age: 23},
	}
	for _, v := range people {
		fmt.Printf("%+v\n",v)
	}

}