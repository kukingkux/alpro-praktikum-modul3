package main

import "fmt"

func main() {
	var alas, tinggi, luas float64
	fmt.Print("Masukkan Alas : ")
	fmt.Scan(&alas)
	fmt.Print("Tinggi : ")
	fmt.Scan(&tinggi)
	luas = (alas * tinggi / 2)
	fmt.Print(luas)
}
