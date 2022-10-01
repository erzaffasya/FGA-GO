package main

import (
	"assignment-1/helpers"
	"fmt"
	"os"
	"strconv"
)

func main() {
	biodata := helpers.Init()

	if len(os.Args) != 1 {
		for i := 1; i < len(os.Args); i++ {
			if helpers.IsLetter(os.Args[i]) {
				fmt.Printf("\nMasukan nomor absen %v tidak valid!\n", os.Args[i])
				continue
			}

			absen, _ := strconv.Atoi(os.Args[i])

			if absen <= 0 {
				fmt.Printf("\nMasukan nomor absen %v tidak valid!\n", absen)
				continue
			}

			if absen >= len(biodata) {
				fmt.Printf("\nBiodata nomor absen %v tidak ditemukan!\n", absen)
				continue
			}

			helpers.Print(biodata, absen)
		}
	} else {
		fmt.Println("\nMasukan nomor absen setelah biodata.go!")
	}
}
