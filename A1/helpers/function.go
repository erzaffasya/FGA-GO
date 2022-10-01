package helpers

import (
	"fmt"
	"unicode"
)

func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}

	return true
}

func Print(biodata []Biodata, absen int) {
	fmt.Println("\nNama\t\t\t\t: ", biodata[absen].Nama)
	fmt.Println("Alamat\t\t\t\t: ", biodata[absen].Alamat)
	fmt.Println("Pekerjaan\t\t\t: ", biodata[absen].Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang\t: ", biodata[absen].Alasan)
}
