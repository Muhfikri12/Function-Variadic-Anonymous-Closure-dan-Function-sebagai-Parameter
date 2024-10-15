package main

import (
	"fmt"
)

// No.1 Buatkan Kode Golang untuk counter angka jika mencapai counter 200 maka counternya berhenti
// No.2 Buatkan Kode untuk mengkonversi dari angka menjadi text hari untuk angka terhitung dari 1 sampai 7 senin sampai minggu
// No.3
// No.4 Tentukan Gradi nilai seorang siswa dari nilai tertentu grade E itu 0-50 grade d 50-60 grade c 60-70 grade b 71-80 a 81-100
// No.5 Buatkan Function untuk menghitung luas dan keliling banguanan datar
// No.6 Buatkan Function untuk menampilkan harga jual berdasarkan diskon tertentu
// No.7 Buatkan Function jika membeli 1 barang tidak mendapatkan diskon, Jika beli 2 diskon 10 persen, jika beli 4 diskon 50 persen lebih dari 4 diskon 75
// No.8 Hitung Total Gaji Karyawan jika 1 jam nya 50rb kemudiannya lemburnya 60rb
/* No.9 buatkan function untuk studi kasus berikut :

- diketahui ada product sepatu  adidas , puma, kappa

- harga sepatu adidas 200.000

- harga sepatu puma 150.000

- harga sepatu kappa 600.000

- harga diskon

- jika membeli sepatu adidas dan puma potongan 50.000

- jika membeli sepatu puma dan kappa potongan 150.000

- jika membeli sepatu adidas dan kappa potongan 75.000

*/

// Jawaban Soal 9
const (
	hargaAdidas = 200000
	hargaPuma = 150000
	hargaKeppa = 600000

)
func totalHarga(adidas, puma, keppa int) int {
	total := 0
	diskon := 0

	total += adidas * hargaAdidas
	total += puma * hargaPuma
	total += keppa * hargaKeppa

	if adidas > 0 && puma > 0 {
		diskon += 50000
	}
	if puma > 0 && keppa > 0 {
		diskon += 150000
	}
	if adidas > 0 && keppa > 0 {
		diskon += 75000
	}

	total -= diskon

	return total
}

// Jawaban Soal 1

func days(hari int) string {

	var tanggal string

	switch hari{
	case 1: 
		tanggal = "senin"
	case 2:
		tanggal = "selasa"
	case 3: 
		tanggal = "Rabu"
	case 4:
		tanggal = "Kamis"
	case 5:
		tanggal = "Jumat"
	case 6:
		tanggal = "Sabtu"
	case 7:
		tanggal = "Minggu"
	}

	return tanggal
	
}

// Jawaban Soal 2
func grade(nilai int) string {

	if nilai < 0 || nilai > 100 {	
		return "Nilai Tidak Valid"
	}

	switch {
	case nilai < 50:
		return "E"
	case nilai > 50 && nilai <= 60:
		return "D"
	case nilai > 60 && nilai <= 70:
		return "C"
	case nilai > 70 && nilai <= 80:
		return "B"
	case nilai > 80 :
		return "A"
	default:
		return "Nilai Tidak Valid"
	}
}

// Jawaban Soal 4
func persegi(sisi int) (float64, float64)  {
	
	luas := sisi * sisi
	keliling := 4 * sisi

	return float64(luas), float64(keliling)
}

// Jawaban Soal 5
func hargaJualArray(hargaAsli []float64, diskon float64) []float64 {

	hargaSetelahDiskon := make([]float64, len(hargaAsli))

	for i, harga := range hargaAsli {
		nilaiDiskon := harga * (diskon / 100)
		hargaSetelahDiskon[i] = harga - nilaiDiskon
	}

	return hargaSetelahDiskon
}

// Jawaban Soal 7
func diakonBarang(hargaBarang int, jumlahBarang int) int {
	
	var diskon int

	switch {
	case jumlahBarang == 1:
		diskon = 0
	case jumlahBarang > 2 && jumlahBarang <= 3:
		diskon = 20
	case jumlahBarang == 4:
		diskon = 50
	case jumlahBarang > 4:
		diskon = 75
	default:
		diskon = 0
	}

	totalHarga := hargaBarang * jumlahBarang
	totalDiskon := totalHarga * (diskon/100)
	hargaSetelahDiskonBarang := totalHarga - totalDiskon

	return hargaSetelahDiskonBarang
}

// Jawaban Soal 8
func gaji(jamKerja int, jamLembur int) int {
	
	totalGajiJamKerja := jamKerja * 50000
	totalLemburJamKerja := jamLembur * 60000

	totalGaji := totalGajiJamKerja + totalLemburJamKerja

	return totalGaji
}

func main()  {

	// Soal No 1
	for counter := 0; counter <= 200; counter++ {
		fmt.Println("Counter Ke", counter)
	}

	// Soal No 2
	fmt.Println(days(5))

	// Soal No 4
	fmt.Println(grade(1000))

	// Soal No 5
	luas, keliling := persegi(5)
	fmt.Printf("Persegi (Sisi = 5): Luas = %.2f, Keliling = %.2f\n", luas, keliling)

	// Soal No 6

	hargaBarang := []float64{100000.0, 150000.0, 200000.0, 50000.0}
	diskon := 20.0 // Diskon 20%
	hargaSetelahDiskon := hargaJualArray(hargaBarang, diskon)

	fmt.Printf("Harga Asli: %v\n", hargaBarang)
	fmt.Printf("Diskon: %.2f%%\n", diskon)
	fmt.Printf("Harga Setelah Diskon: %v\n", hargaSetelahDiskon)

	// Soal No 7

	hargaBarang1 := 100000
	jumlahBarang := 2

	hargaSetelahDiskonbaru := diakonBarang(hargaBarang1, jumlahBarang)

	fmt.Println(hargaSetelahDiskonbaru)

	// Soal No 8 

	jamKerja := 40
	jamLembur := 5

	totalGaji := gaji(jamKerja, jamLembur)

	fmt.Println(totalGaji)

	// Soal No 9

	totalHarga := totalHarga(1, 1, 0)
	fmt.Printf("Total harga setelah diskon: %d\n", totalHarga)
}