package main

import "fmt"

func durasi(jam, menit int) int {
	var hasil int
	if jam > 0 && menit >= 10 {
		hasil = jam + 1
	} else if jam >= 0 && menit < 10 {
		hasil = jam
		if jam == 0 && menit > 0 {
			hasil = 1
		} else if jam == 0 && menit == 0 {
			hasil = 0
		}
	}
	return hasil
}

func potongan(durasi int, tarif int) float64 {
	var hasil float64
	fmt.Println(durasi, tarif)
	if durasi > 3 {
		hasil = float64(tarif) * float64(durasi) * 0.1
	} else if durasi <= 3 {
		hasil = 0
	}
	return hasil
}

func tarif(member bool) int {
	var hasil int
	if member == true {
		hasil = 3500
	} else {
		hasil = 5000
	}
	return hasil
}

func hitungSewa(jam, menit int, member bool, biaya *float64) {
	var lama, harga int
	var diskon float64
	lama = durasi(jam, menit)
	harga = tarif(member)
	diskon = potongan(lama, harga)
	*biaya = (float64(lama) * float64(harga))-diskon
}

func main() {
	var jam, menit int
	var member bool
	var total float64

	fmt.Scan(&jam, &menit, &member)
	hitungSewa(jam, menit, member, &total)
	fmt.Println("total",total)
}