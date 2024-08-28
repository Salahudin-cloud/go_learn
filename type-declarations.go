package main

import "fmt"

func main() {
	// type declaration = membuat alias untuk tipe data
	// cara membuat type declaration yaitu : type nama_declaration tipe_data

	type NIM string
	type isGraduated bool

	var mahasiswaSalahudin NIM = "21220004"
	var mahasiswaStatus isGraduated = false

	fmt.Println(mahasiswaSalahudin)
	fmt.Println(mahasiswaStatus)
}