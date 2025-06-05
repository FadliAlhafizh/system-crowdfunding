package main

import (
	"fmt"
	"strings"
)

const MAXPROYEK int = 10

type proyek struct {
	ID           int
	Name         string
	Kategori     string
	Donasi       float64
	TargetDonasi float64
	TotalDonatur int
}

var currentIndex = 0
var currentID = 1
var listProyeks [MAXPROYEK]proyek

func isKategori(kategori string) bool {
	return kategori == "kesehatan" || kategori == "pendidikan" || kategori == "pendanaan"
}

func addProyek(name *string, kategori *string, targetDonasi *float64) {

	fmt.Println("Masukan Nama Proyek: ")
	fmt.Scan(name)

	for {
		fmt.Println("Pilih kategori berikut: 1. Kesehatan/2. Pendidikan/3. Pendanaan")
		fmt.Scan(kategori)
		if isKategori(*kategori) {
			break
		}
	}
	fmt.Print("Masukan Target Donasi: ")
	fmt.Scan(targetDonasi)
	fmt.Println()

	listProyeks[currentIndex] = proyek{
		ID:           currentID,
		Name:         *name,
		Kategori:     *kategori,
		Donasi:       0,
		TargetDonasi: *targetDonasi,
		TotalDonatur: 0,
	}
	currentIndex++
	currentID++
}

func editProyek(id *int, name *string, kategori *string, targetDonasi *float64, total *int) {
	for i := 0; i < currentIndex; i++ {
		if listProyeks[i].ID == *id {
			listProyeks[i].Name = *name
			listProyeks[i].Kategori = *kategori
			listProyeks[i].TargetDonasi = *targetDonasi
			listProyeks[i].TotalDonatur = *total
			return
		}
	}
}

func deleteProyek(id *int) {
	for i := 0; i < currentIndex; i++ {
		if listProyeks[i].ID == *id {
			for j := i; j < currentIndex-1; j++ {
				listProyeks[j] = listProyeks[j+1]
			}
			currentIndex--
			return
		}
	}
}

func searchProject() {

	var metode int
	var keyword string

	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search")
	fmt.Println("Pilih metode pencarian: ")
	fmt.Scan(&metode)

	fmt.Print("Masukan kata kunci (nama/kategori)")
	fmt.Scan(&keyword)

	switch metode {
	case 1:
		sequentialSearch(keyword)
	case 2:
		InsertionSortByName(listProyeks[:currentIndex], currentIndex)
		binarySearch(keyword)
	default:
		fmt.Println("Pilihan tidak valid")
	}
}

func sequentialSearch(keyword string) {
	found := false
	for i := 0; i < currentIndex; i++ {
		if strings.EqualFold(listProyeks[i].Name, keyword) || strings.EqualFold(listProyeks[i].Kategori, keyword) {
			fmt.Println(listProyeks[i])
			found = true
		}
	}
	if !found {
		fmt.Println("Proyek tidak ditemukan")
	}
}

func binarySearch(keyword string) int {
	var left, right int
	left = 0
	right = 0

	for left <= right {
		var mid int
		mid = (left + right) / 2
		if listProyeks[mid].Name == keyword || listProyeks[mid].Kategori == keyword {
			return mid
		} else if listProyeks[mid].Name < keyword || listProyeks[mid].Kategori < keyword {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func SelectionSortByRaised(projects []proyek) {
	n := len(projects)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if projects[j].TargetDonasi > projects[maxIdx].TargetDonasi {
				maxIdx = j
			}
		}
		projects[i], projects[maxIdx] = projects[maxIdx], projects[i]
	}
}

func SelectionSortByDonors(projects []proyek) {
	n := len(projects)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if projects[j].TotalDonatur > projects[maxIdx].TotalDonatur {
				maxIdx = j
			}
		}
		projects[i], projects[maxIdx] = projects[maxIdx], projects[i]
	}
}

func InsertionSortByName(projects []proyek, n int) {
	for i := 1; i < n; i++ {
		key := projects[i]
		j := i - 1
		for j >= 0 && projects[j].Name > key.Name {
			projects[j+1] = projects[j]
			j--
		}
		projects[j+1] = key
	}
}

func InsertionSortByRaised(proyek []proyek, n int) {
	for i := 1; i < n; i++ {
		idx := proyek[i]
		j := i - 1
		for j >= 0 && proyek[j].TargetDonasi < idx.TargetDonasi {
			proyek[j+1] = proyek[j]
			j--
		}
		proyek[j+1] = idx
	}
}

func InsertionSortByDonors(proyek []proyek, n int) {
	for i := 1; i < n; i++ {
		idx := proyek[i]
		j := i - 1
		for j >= 0 && float64(proyek[j].TotalDonatur) < float64(idx.TotalDonatur) {
			proyek[j+1] = proyek[j]
			j--
		}
		proyek[j+1] = idx
	}
}

func sortingMenu() {
	var pilih int
	fmt.Println("1. Urutkan berdasarkan Target Donasi ")
	fmt.Println("2. Urutkan berdasarkan Total Donatur ")
	fmt.Scan(&pilih)

	switch pilih {
	case 1:
		SelectionSortByRaised(listProyeks[:currentIndex])
	case 2:
		SelectionSortByDonors(listProyeks[:currentIndex])
	case 3:
		InsertionSortByRaised(listProyeks[:currentIndex], currentIndex)
	case 4:
		InsertionSortByDonors(listProyeks[:currentIndex], currentIndex)
	}
}

func menu() {

	fmt.Println("=== Sistem Proyek Crowdfunding ===")
	fmt.Println("1. Tambah Proyek")
	fmt.Println("2. Ubah Proyek")
	fmt.Println("3. Hapus Proyek")
	fmt.Println("4. Cari Proyek")
	fmt.Println("5. Urutkan Proyek")
	fmt.Println("6. Tampilkan Proyek dengan Target Tercapai")
	fmt.Println("7. Tampilkan Semua Proyek")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih menu: ")
}

func cetak(name string, kategori string, totalDonasi float64, totalDonatur int) {
	fmt.Print(name)
	fmt.Print(kategori)
	fmt.Print(totalDonasi)
	fmt.Print(totalDonatur)
}

func tampilkanSemuaProyek() {
	if currentIndex == 0 {
		fmt.Println("Belum ada proyek yang terdaftar.")
		return
	}
	for i := 0; i < currentIndex; i++ {
		p := listProyeks[i]
		cetak(p.Name, p.Kategori, p.Donasi, p.TotalDonatur)
	}
}

func tampilkanProyekSukses() {
	found := false
	fmt.Println("=== Proyek yang Mencapai Target Donasi ===")
	for i := 0; i < currentIndex; i++ {
		if listProyeks[i].Donasi >= listProyeks[i].TargetDonasi {
			cetak(listProyeks[i].Name, listProyeks[i].Kategori, listProyeks[i].Donasi, listProyeks[i].TotalDonatur)
			found = true
		}
	}
	if !found {
		fmt.Println("Belum ada proyek yang mencapai target.")
	}
}

func main() {

	for {
		menu()
		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			var name, kategori string
			var targetDonasi float64
			addProyek(&name, &kategori, &targetDonasi)
		case 2:
			// editProyek()
		case 3:
			// deleteProyek()
		case 4:
			searchProject()
		case 5:
			sortingMenu()
		case 6:
			tampilkanProyekSukses()
		case 7:
			tampilkanSemuaProyek()
		case 0:
			fmt.Println("Terima kasih telah menggunakan sistem crowdfunding")
			return
		default:
			fmt.Println("Menu tidak valid")
		}
	}
}
