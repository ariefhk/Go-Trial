package main

import (
	"fmt"
)

func main()  {
	fmt.Println("=== Aplikasi Hitung ===")
	fmt.Println("1).Hitung +-*/%")
	fmt.Println("2).Hitung Luas Bangun Datar")
	fmt.Println("3).Hitung Sin Cos Tan")
	fmt.Println("4).Hitung Gerak Lurus")
	fmt.Println("5).Hitung Energi Potensial dan Kinetik")
	fmt.Println("6).Hitung Frekuensi dan Getaran")
	fmt.Println("7).Hitung Massa Jenis")
	fmt.Println("8).Hitung Konversi Suhu")

	var operator int
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&operator)

	switch operator {
	case 1:
		kalkulator()
		break
	}
}

func kalkulator()  {
	
	fmt.Println("=== Hitung +-*/% ===")
	fmt.Println("1).Penjumlahan (+)")
	fmt.Println("2).Pengurangan (-)")
	fmt.Println("3).Perkalian (*)")
	fmt.Println("4).Pembagian (/)")
	fmt.Println("5).Modulus (%)")

	var operator int
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&operator)
	switch operator {
	case 1:
		var angka1 float64 
		var angka2 float64 
		fmt.Print("Masukan Angka Pertama: ")
		fmt.Scan(&angka1)
		fmt.Print("Masukan Angka Kedua: ")
		fmt.Scan(&angka2)

		hasil := angka1 + angka2
		fmt.Println("Hasil: ",angka1," + ",angka2," = ", hasil)
		break
	case 2:
		var angka1 float64 
		var angka2 float64 
		fmt.Print("Masukan Angka Pertama: ")
		fmt.Scan(&angka1)
		fmt.Print("Masukan Angka Kedua: ")
		fmt.Scan(&angka2)

		hasil := angka1 - angka2
		fmt.Println("Hasil: ",angka1," - ",angka2," = ", hasil)
		break
	case 3:
		var angka1 float64 
		var angka2 float64 
		fmt.Print("Masukan Angka Pertama: ")
		fmt.Scan(&angka1)
		fmt.Print("Masukan Angka Kedua: ")
		fmt.Scan(&angka2)

		hasil := angka1 * angka2
		fmt.Println("Hasil: ",angka1," * ",angka2," = ", hasil)
		break
	case 4:
		var angka1 float64 
		var angka2 float64 
		fmt.Print("Masukan Angka Pertama: ")
		fmt.Scan(&angka1)
		fmt.Print("Masukan Angka Kedua: ")
		fmt.Scan(&angka2)

		hasil := angka1 / angka2
		fmt.Println("Hasil: ",angka1," / ",angka2," = ", hasil)
		break
	case 5:
		var angka1 int 
		var angka2 int 
		fmt.Print("Masukan Angka Pertama: ")
		fmt.Scan(&angka1)
		fmt.Print("Masukan Angka Kedua: ")
		fmt.Scan(&angka2)

		hasil := angka1 % angka2
		fmt.Println("Hasil: ",angka1," % ",angka2," = ", hasil)
		break
	}
	
}