package main

import (
	"fmt"
)

const MAXPROYEK int = 10

type proyek struct {
	ID           int
	Name         string
	Kategori     string
	TotalDonasi  float64
	TotalDonatur int
}

var currentIndex = 0
var currentID = 1
var listProyeks [MAXPROYEK]proyek

func addProyek(name *string, kategori *string, donasi *float64, total *int) {

	fmt.Scan(name, kategori, donasi, total)

	listProyeks[currentIndex] = proyek{
		ID:           currentID,
		Name:         *name,
		Kategori:     *kategori,
		TotalDonasi:  *donasi,
		TotalDonatur: *total,
	}
	currentIndex++
	currentID++
}

func editProyek() {

}

func deleteProyek() {

}

func menu() {

}

func cetak(name string, kategori string, totalDonasi float64, totalDonatur int) {
	fmt.Print(name)
	fmt.Print(kategori)
	fmt.Print(totalDonasi)
	fmt.Print(totalDonatur)
}

func main() {
	var name string
	var kategori string
	var totalDonasi float64
	var totalDonatur int

	addProyek(&name, &kategori, &totalDonasi, &totalDonatur)
	cetak(name, kategori, totalDonasi, totalDonatur)
}
