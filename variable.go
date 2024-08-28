package main

import "fmt"

func main() {
	// cara membuat variable di golang ada 3 cara
	// 1. membuat variable dengan format seperti ini var_namaVariable_tipedata
	var firstName string
	firstName = "salahudin" // inisialisasi
	fmt.Println(firstName)
	
	// 2. dengan cara menggunakan :=
	midleName := "udin" // membuat cara kedua harus wajib di inialisasi
	fmt.Println(midleName)

	// 3. membuat varible multi variable dengan menggunakan cara var(namaVaribel1, namaVaribel2 ....)
	var (
		myName = "ayyubi"
		age    = 20
	)
	fmt.Println(myName)
	fmt.Println(age)
}
