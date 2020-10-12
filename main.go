package main

import ("fmt")

func main() {
	var angka = 5
	if angka % 2 == 0 {
		fmt.Println("Inputan harus ganjil")
	} else {
		for ulangSatu := 0; ulangSatu < angka; ulangSatu++ {
			for ulangDua := 0; ulangDua < angka; ulangDua++ {
				if ulangDua == 0 || ulangSatu == angka/2 || ulangDua == angka-1 {
					fmt.Print("*  ")
				} else {
					fmt.Print("=  ")
				}
			}
			fmt.Println("\n")
		}
	}
}