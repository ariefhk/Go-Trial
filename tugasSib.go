package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("=== Aplikasi Hitung ===")
	fmt.Println("1).Hitung +-*/%")
	fmt.Println("2).Hitung Luas Bangun Datar")
	fmt.Println("3).Hitung Sin Cos Tan")
	fmt.Println("4).Hitung Gerak Lurus")
	fmt.Println("5).Hitung Energi Potensial dan Kinetik")
	fmt.Println("6).Hitung Frekuensi dan Getaran")
	fmt.Println("7).Hitung Massa Jenis")
	fmt.Println("8).Hitung Daya, Tekanan, Usaha atau Gaya")
	fmt.Println("9).Hitung Konversi Suhu")

	var operator int
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&operator)

	switch operator {
	case 1:
		kalkulator()
		break
	case 2:
		luasBangunDatar()
		break
	case 3:
		sinCosTan()
		break
	case 4:
		gerakLurus()
		break
	case 5:
		energiPotensialKinetik()
		break
	case 6:
		frekuensi_getaran()
		break
	case 7:
		masaJenis()
		break
	case 8:
		dayaTekananUsaha()
		break
	case 9:
		konversiSuhu()
		break
	}
}

func kalkulator() {

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
		fmt.Println("Hasil: ", angka1, " + ", angka2, " = ", hasil)
		break
	case 2:
		var angka1 float64
		var angka2 float64
		fmt.Print("Masukan Angka Pertama: ")
		fmt.Scan(&angka1)
		fmt.Print("Masukan Angka Kedua: ")
		fmt.Scan(&angka2)

		hasil := angka1 - angka2
		fmt.Println("Hasil: ", angka1, " - ", angka2, " = ", hasil)
		break
	case 3:
		var angka1 float64
		var angka2 float64
		fmt.Print("Masukan Angka Pertama: ")
		fmt.Scan(&angka1)
		fmt.Print("Masukan Angka Kedua: ")
		fmt.Scan(&angka2)

		hasil := angka1 * angka2
		fmt.Println("Hasil: ", angka1, " * ", angka2, " = ", hasil)
		break
	case 4:
		var angka1 float64
		var angka2 float64
		fmt.Print("Masukan Angka Pertama: ")
		fmt.Scan(&angka1)
		fmt.Print("Masukan Angka Kedua: ")
		fmt.Scan(&angka2)

		hasil := angka1 / angka2
		fmt.Println("Hasil: ", angka1, " / ", angka2, " = ", hasil)
		break
	case 5:
		var angka1 int
		var angka2 int
		fmt.Print("Masukan Angka Pertama: ")
		fmt.Scan(&angka1)
		fmt.Print("Masukan Angka Kedua: ")
		fmt.Scan(&angka2)

		hasil := angka1 % angka2
		fmt.Println("Hasil: ", angka1, " % ", angka2, " = ", hasil)
		break
	}

}

func luasBangunDatar() {
	fmt.Println("=== Hitung Luas Bangun Datar ===")
	fmt.Println("1).Persegi")
	fmt.Println("2).Persegi Panjang")
	fmt.Println("3).Segitiga")
	fmt.Println("4).Lingkaran")

	var operator int
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&operator)
	switch operator {
	case 1:
		var sisi float64
		fmt.Println("=== Hitung Luas Persegi ===")
		fmt.Print("Masukkan Sisi: ")
		fmt.Scan(&sisi)
		luas := sisi * sisi

		fmt.Println("Hasil: ", luas)
		break
	case 2:
		var panjang, lebar float64
		fmt.Println("=== Hitung Luas Persegi Panjang===")
		fmt.Print("Masukkan Panjang: ")
		fmt.Scan(&panjang)
		fmt.Print("Masukkan Lebar: ")
		fmt.Scan(&lebar)

		luas := panjang * lebar
		fmt.Println("Hasil: ", luas)
		break
	case 3:
		var alas, tinggi float64
		fmt.Println("=== Hitung Luas Segitiga===")
		fmt.Print("Masukkan alas: ")
		fmt.Scan(&alas)
		fmt.Print("Masukkan tinggi: ")
		fmt.Scan(&tinggi)

		luas := (alas * tinggi) / 2
		fmt.Println("Hasil: ", luas)
		break
	case 4:
		var r float64
		fmt.Println("=== Hitung Luas Lingkaran===")
		fmt.Print("Masukkan jari-jari: ")
		fmt.Scan(&r)

		luas := (r * r) * 3.14
		fmt.Println("Hasil: ", luas)
		break
	}
}

func sinCosTan() {
	var sudut float64
	fmt.Println("=== Hitung Sin Cos Tan ===")
	fmt.Print("Masukkan Sudut: ")
	fmt.Scan(&sudut)
	fmt.Println("=== Hasil ===")
	fmt.Println("Sinus: ", math.Sin(sudut))
	fmt.Println("Cos: ", math.Cos(sudut))
	fmt.Println("Tan: ", math.Tan(sudut))
}

func gerakLurus() {
	fmt.Println("=== Hitung Gerak Lurus ===")
	fmt.Println("1).Gerak Lurus Beraturan")
	fmt.Println("2).Gerak Lurus Berubah Beraturan")

	var operator int
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&operator)
	switch operator {
	case 1:
		var v, t float64
		fmt.Println("=== Hitung Gerak Lurus Beraturan ===")
		fmt.Print("Masukkan Jarak (meter/v): ")
		fmt.Scan(&v)
		fmt.Print("Masukkan Waktu (sekon/t): ")
		fmt.Scan(&t)

		rumus := v * t
		fmt.Println("Hasil: ", rumus)
		break
	case 2:
		fmt.Println("=== Hitung Gerak Lurus Berubah Beraturan ===")
		fmt.Println("1).Mencari Kecepatan Akhir jika Waktu Diketahui")
		fmt.Println("2).Mencari Kecepatan Akhir jika Jarak Diketahui")
		fmt.Println("3).Mencari Jarak Tempuh")

		var operator int
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&operator)
		switch operator {
		case 1:
			var v0, t, a float64
			fmt.Println("=== Mencari Kecepatan Akhir jika Waktu Diketahui ===")
			fmt.Print("Masukkan Kecepatan Awal (m/s): ")
			fmt.Scan(&v0)
			fmt.Print("Masukkan Waktu (sekon/t): ")
			fmt.Scan(&t)
			fmt.Print("Masukkan Percepatan (m/s2): ")
			fmt.Scan(&a)

			rumus := v0 + (t * a)
			fmt.Println("Hasil: ", rumus)
			break
		case 2:
			var v0, s, a float64
			fmt.Println("=== Mencari Kecepatan Akhir jika Jarak Diketahui ===")
			fmt.Print("Masukkan Kecepatan Awal (m/s): ")
			fmt.Scan(&v0)
			fmt.Print("Masukkan Jarak (meter/s): ")
			fmt.Scan(&s)
			fmt.Print("Masukkan Percepatan (m/s2): ")
			fmt.Scan(&a)

			rumus := (v0 * v0) + (2 * a * s)
			fmt.Println("Hasil: ", rumus)
			break
		case 3:
			var v0, t, a float64
			fmt.Println("=== Mencari Jarak Tempuh ===")
			fmt.Print("Masukkan Kecepatan Awal (m/s): ")
			fmt.Scan(&v0)
			fmt.Print("Masukkan Waktu (sekon/t): ")
			fmt.Scan(&t)
			fmt.Print("Masukkan Percepatan (m/s2): ")
			fmt.Scan(&a)

			rumus := (v0 * t) + (0.5 * a * t * t)
			fmt.Println("Hasil: ", rumus)
			break
		}
	}
}

func energiPotensialKinetik() {
	fmt.Println("=== Hitung Energi Potensial & Kinetik ===")
	fmt.Println("1).Hitung Energi Potensial")
	fmt.Println("2).Hitung Energi Kinetik")
	var operator int
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&operator)
	switch operator {
	case 1:
		var m, g, h float64
		fmt.Println("=== Hitung Energi Potensial ===")
		fmt.Print("Masukkan Massa Benda (kg): ")
		fmt.Scan(&m)
		fmt.Print("Masukkan Gravitasi (m/s2): ")
		fmt.Scan(&g)
		fmt.Print("Masukkan Tinggi Benda dari permukaan tanah (meter/h): ")
		fmt.Scan(&h)

		rumus := m * g * h
		fmt.Println("Hasil: ", rumus)
		break
	case 2:
		var m, v float64
		fmt.Println("=== Hitung Energi Kinetik ===")
		fmt.Print("Masukkan Massa Benda (kg): ")
		fmt.Scan(&m)
		fmt.Print("Masukkan Kecepatan Benda (m/s): ")
		fmt.Scan(&v)

		rumus := 0.5 * m * v * v
		fmt.Println("Hasil: ", rumus)
		break
	}
}

func frekuensi_getaran() {

	fmt.Println("=== Hitung Frekuensi atau Getaran ===")
	fmt.Println("1).Hitung Frekuensi")
	fmt.Println("2).Hitung Getaran")

	var operator int
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&operator)
	switch operator {
	case 1:
		var n, t float64
		fmt.Println("=== Hitung Frekuensi ===")
		fmt.Print("Masukkan Jumlah Getaran: ")
		fmt.Scan(&n)
		fmt.Print("Masukkan Waktu (s): ")
		fmt.Scan(&t)

		rumus := n / t
		fmt.Println("Hasil: ", rumus)
		break
	case 2:
		var f float64
		fmt.Println("=== Hitung Getaran ===")
		fmt.Print("Masukkan frekuensi: ")
		fmt.Scan(&f)

		rumus := 1 / f
		fmt.Println("Hasil: ", rumus)
		break
	}

}

func masaJenis() {
	var m, v float64
	fmt.Println("=== Hitung Masa Jenis ===")
	fmt.Print("Masukkan nilai m: ")
	fmt.Scan(&m)
	fmt.Print("Masukkan nilai v: ")
	fmt.Scan(&v)

	rumus := m / v
	fmt.Println("Hasil: ", rumus)
}

func dayaTekananUsaha() {
	fmt.Println("=== Hitung Daya, Tekanan, Usaha, atau Gaya ===")
	fmt.Println("1).Hitung Daya")
	fmt.Println("2).Hitung Tekanan")
	fmt.Println("3).Hitung Usaha")
	fmt.Println("4).Hitung Gaya")
	var operator int
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&operator)
	switch operator {
	case 1:
		var w, t float64
		fmt.Println("=== Hitung Daya ===")
		fmt.Print("Masukkan usaha (joule): ")
		fmt.Scan(&w)
		fmt.Print("Masukkan Waktu (sekon): ")
		fmt.Scan(&t)

		rumus := w / t
		fmt.Println("Hasil: ", rumus)
		break
	case 2:
		var f, a float64
		fmt.Println("=== Hitung Tekanan ===")
		fmt.Print("Masukkan gaya (newton): ")
		fmt.Scan(&f)
		fmt.Print("Masukkan luas alas (m2): ")
		fmt.Scan(&a)

		rumus := f / a
		fmt.Println("Hasil: ", rumus)
		break
	case 3:
		var f, s float64
		fmt.Println("=== Hitung Usaha ===")
		fmt.Print("Masukkan gaya (newton): ")
		fmt.Scan(&f)
		fmt.Print("Masukkan jarak (m): ")
		fmt.Scan(&s)

		rumus := f * s
		fmt.Println("Hasil: ", rumus)
		break
	case 4:
		var m, a float64
		fmt.Println("=== Hitung Gaya ===")
		fmt.Print("Masukkan massa (kg): ")
		fmt.Scan(&m)
		fmt.Print("Masukkan percepatan (m/s2): ")
		fmt.Scan(&a)

		rumus := m * a
		fmt.Println("Hasil: ", rumus)
		break
	}

}

func konversiSuhu() {
	fmt.Println("=== Hitung Konversi Suhu ===")
	fmt.Println("1).Konversi Celcius ")
	fmt.Println("2).Konversi Reamur")
	fmt.Println("3).Hitung Fahrenreit")
	fmt.Println("4).Hitung Kelvin")
	var operator int
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&operator)
	switch operator {
	case 1:
		var c float64
		fmt.Println("=== Konversi Celcius ===")
		fmt.Print("Masukkan Suhu Celcius: ")
		fmt.Scan(&c)

		reamur := c * (4 / 5)
		kelvin := c + 273
		fahrenreit := (c * (9 / 5)) + 32

		fmt.Println("Reamur: ", reamur)
		fmt.Println("Kelvin: ", kelvin)
		fmt.Println("Fahrenreit: ", fahrenreit)
		break
	case 2:
		var r float64
		fmt.Println("=== Konversi Reamur ===")
		fmt.Print("Masukkan Suhu Reamur: ")
		fmt.Scan(&r)

		celcius := r * (5 / 4)
		kelvin := (r * (5 / 4)) + 273
		fahrenreit := (r * (9 / 4)) + 32

		fmt.Println("Reamur: ", celcius)
		fmt.Println("Kelvin: ", kelvin)
		fmt.Println("Fahrenreit: ", fahrenreit)
		break
	case 3:
		var f float64
		fmt.Println("=== Konversi Fahrenreit ===")
		fmt.Print("Masukkan Suhu Fahrenreit: ")
		fmt.Scan(&f)

		celcius := (5 / 9) * (f - 32)
		reamur := (4 / 9) * (f - 32)
		kelvin := (5/9)*(f-32) + 273

		fmt.Println("Celcius: ", celcius)
		fmt.Println("Reamur: ", reamur)
		fmt.Println("Kelvin: ", kelvin)
		break
	case 4:
		var k float64
		fmt.Println("=== Konversi Kelvin ===")
		fmt.Print("Masukkan Suhu Kelvin: ")
		fmt.Scan(&k)

		celcius := k - 273
		reamur := (4 / 5) * (k - 273)
		fahrenreit := (9/5)*(k-273) + 32

		fmt.Println("Celcius: ", celcius)
		fmt.Println("Reamur: ", reamur)
		fmt.Println("Fahrenreit: ", fahrenreit)
		break
	}
}
