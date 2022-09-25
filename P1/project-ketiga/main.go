package main

import (
	"log"
)

func main(){
	// nil
	// var nama *string
	// fmt.Printf("%T \n",nama)
	// log.Println(nama)
	// name := "Rohmat" //ambil kode alamat
	// nama = &name
	// log.Println(*nama)

	// const wajib di isi dan tidak boleh pointer
	const nama string = "Rohmat"	
	log.Printf("%T",nama)

	// aritmatika (+,-,*,/,%,++,--)
	// var ar_value1 = 2+2
	// var ar_value2 = 2-2
	// var ar_value3 = 2*2
	// var ar_value4 = 2/2
	// var ar_value5 = 10%3 //modulus atau sisa bagi
	// var ar_value6 = ar_value6++
	// var ar_value7 = ar_value7--
	// log.Println(ar_value1)
	// log.Println(ar_value2)
	// log.Println(ar_value3)
	// log.Println(ar_value4)
	// log.Println(ar_value5)

	// perbandingan (==,!=,>,<,>=,<=) return true false
	// log.Println(2==2)
	// log.Println(2!=2)
	// log.Println(2>2)
	// log.Println(2<2)
	// log.Println(2>=2)
	// log.Println(2<=2)

	// logika (!,||,&&)
	log.Println((2==2) && (2<=2))
}