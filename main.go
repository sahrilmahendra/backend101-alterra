package main

import "fmt"

func main() {
	// angkaku := 10

	// fmt.Println(angkaku)

	// fmt.Println("Hello World")

	// alas := 10
	// tinggi := 15
	// L := (alas * tinggi) / 2
	// fmt.Println("Luas segitiga =", L)

	// helloWorld := "hello" + " " + "world"
	// fmt.Println(helloWorld)

	// myAge := 17
	// if dadAge := 9; dadAge < myAge {
	// 	fmt.Println(dadAge)
	// } else {
	// 	fmt.Println("I'm older")
	// }

	// angka := 10
	// if angka%2 == 0 {
	// 	fmt.Println("Angka Genap")
	// } else {
	// 	fmt.Println("Angka Ganjil")
	// }

	for i := 0; i < 5; i++ {
		fmt.Println("i ke", i)
	}

	cek := 10
	cekPrima := 0
	for i := 1; i <= cek; i++ {
		if cek%i == 0 {
			cekPrima++
		}
	}
	if cekPrima == 2 {
		fmt.Println("bilangan", cek, "adalah bilangan prima")
	} else {
		fmt.Println("bukan bilangan prima")
	}
}
