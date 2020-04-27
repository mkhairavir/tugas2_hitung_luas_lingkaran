package main

import "fmt"

func main()  {
	var pi float32 = 3.14
	var jariLingkaran float32 = 5
	var luasLingkaran float32 = pi * jariLingkaran * jariLingkaran

	fmt.Println("Selamat Datang di Program Penghitung Luas Lingkaran")
	fmt.Printf("Hasil hitung dari lingkaran dengan jari - jari %f adalah: %f", jariLingkaran, luasLingkaran)
}