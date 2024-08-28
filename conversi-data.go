package main

import "fmt"

func main() {
	// 1. cara melakukan konversi pada data number
	/*
		note : hati hati pada conversi number pehatikan batas tipe data nya jika
		tidak terjadi interger overflow
	*/
	var angkaSatu int64 = 1000000
	var konInt32 int32 = int32(angkaSatu)

	fmt.Println(konInt32)

	// 2. cara melakukan konversi pada string
	var firstName string = "salahudin"
	var showTheChar = string(firstName[1]) // take char a and convert to word
	fmt.Println(showTheChar)
}
