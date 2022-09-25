package main

import (
	"fmt"
	"log"
	"strconv"
)	

func main()  {
	log.Println("Oke")
	// greeting := agreeNameInFile("Rohmat",20)
	// log.Println(greeting)

	greeting,err := agreeNameInFile("Rohmat","20")
	if err != nil {
		log.Println(err)
	}
	log.Println(greeting)
}

// func agreeNameInFile(name string,age uint) string  {	
// 	return fmt.Sprintf("Hello My name is %s and my age = %d",name,age)
// }

// str to int
func agreeNameInFile(name string,age string) (string, error)  {	
	intAge, err := strconv.ParseInt(age,0,8)
	if err != nil {
		return "",err
	}
	return fmt.Sprintf("Hello My name is %s and my age = %d",name,intAge),nil
}