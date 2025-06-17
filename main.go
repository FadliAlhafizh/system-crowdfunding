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
	var found bool
	fmt.Println("Masukkan Nama Proyek: ")
	fmt.Scan(name)

	found = false

	for !found {
		var pilihanKategori int
		fmt.Println("Pilih kategori:")
		fmt.Println("1. Kesehatan")
		fmt.Println("2. Pendidikan")
		fmt.Println("3. Pendanaan")
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&pilihanKategori)

		switch pilihanKategori {
		case 1:
			*kategori = "kesehatan"
			found = true
		case 2:
			*kategori = "pendidikan"
			found = true
		case 3:
			*kategori = "pendanaan"
			found = true
		default:
			fmt.Println("Pilihan tidak valid.")
			found = false
		}
	}

	fmt.Print("Masukkan Target Donasi: ")
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

func editProyek() {
	var name, kategori string
	var targetDonasi float64

	fmt.Print("Masukkan nama proyek yang ingin diubah: ")
	fmt.Scan(&name)

	found := false
	for i := 0; i < currentIndex; i++ {
		if listProyeks[i].Name == name {
			found = true
			fmt.Printf("Proyek ditemukan: %s\n", listProyeks[i].Name)
			fmt.Print("Masukkan nama baru: ")
			fmt.Scan(&name)

			foundKategori := false

			for !foundKategori {
				var pilihanKategori int
				fmt.Println("Pilih kategori baru:")
				fmt.Println("1. Kesehatan")
				fmt.Println("2. Pendidikan")
				fmt.Println("3. Pendanaan")
				fmt.Print("Pilihan Anda: ")
				fmt.Scan(&pilihanKategori)

				if pilihanKategori == 1 {
					kategori = "kesehatan"
					foundKategori = true
				} else if pilihanKategori == 2 {
					kategori = "pendidikan"
					foundKategori = true
				} else if pilihanKategori == 3 {
					kategori = "pendanaan"
					foundKategori = true
				} else {
					fmt.Println("Kategori tidak valid.")
				}
			}

			fmt.Print("Masukkan target donasi baru: ")
			fmt.Scan(&targetDonasi)

			listProyeks[i].Name = name
			listProyeks[i].Kategori = kategori
			listProyeks[i].TargetDonasi = targetDonasi
			fmt.Println("Data proyek berhasil diubah.")
			break
		}
	}
	if !found {
		fmt.Println("Proyek dengan ID tersebut tidak ditemukan.")
	}
}

func deleteProyek() {
	var name string
	fmt.Print("Masukkan nama proyek yang ingin dihapus: ")
	fmt.Scan(&name)

	found := false
	for i := 0; i < currentIndex; i++ {
		if listProyeks[i].Name == name {
			for j := i; j < currentIndex-1; j++ {
				listProyeks[j] = listProyeks[j+1]
			}
			currentIndex--
			fmt.Println("Proyek berhasil dihapus.")
			found = true
		}
	}
	if !found {
		fmt.Println("Proyek dengan Nama tersebut tidak ditemukan.")
	}
}

func donasiProyek(name *string, donasi float64) {
	fmt.Print("Masukkan Nama Proyek: ")
	fmt.Scan(name)

	for i := 0; i < currentIndex; i++ {
		if listProyeks[i].Name == *name {
			fmt.Print("Masukkan jumlah donasi: ")
			fmt.Scan(&donasi)

			listProyeks[i].Donasi += donasi
			listProyeks[i].TotalDonatur++
			fmt.Println("Donasi berhasil ditambahkan!")
			return
		}
	}
	fmt.Println("Proyek tidak ditemukan.")
}

func searchProject() {
	var metode int
	var keyword string

	fmt.Println("Pilih metode pencarian:")
	fmt.Println("1. Cari berdasarkan NAMA (Sequential Search)")
	fmt.Println("2. Cari berdasarkan KATEGORI (Sequential Search)")
	fmt.Println("3. Cari berdasarkan NAMA (Binary Search)")
	fmt.Println("3. Cari berdasarkan KATEGORI (KATEGORI Search)")
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&metode)

	switch metode {
	case 1:
		fmt.Print("Masukkan nama proyek: ")
		fmt.Scan(&keyword)
		sequentialSearchByName(keyword)
	case 2:
		fmt.Print("Masukkan kategori (kesehatan/pendidikan/pendanaan): ")
		fmt.Scan(&keyword)
		if isKategori(keyword) {
			sequentialSearchByKategori(keyword)
		} else {
			fmt.Println("Kategori tidak valid.")
		}
	case 3:
		fmt.Print("Masukkan nama proyek: ")
		fmt.Scan(&keyword)
		InsertionSortByName(listProyeks[:currentIndex], currentIndex)
		idx := binarySearchByName(keyword)
		if idx != -1 {
			cetak(listProyeks[idx].Name, listProyeks[idx].Kategori, listProyeks[idx].Donasi, listProyeks[idx].TotalDonatur, listProyeks[idx].TargetDonasi)
		} else {
			fmt.Println("Proyek tidak ditemukan.")
		}
	case 4:
		fmt.Print("Masukkan nama proyek: ")
		fmt.Scan(&keyword)
		InsertionSortByKategori(listProyeks[:currentIndex], currentIndex)
		idx := binarySearchByKategori(keyword)
		if idx != -1 {
			cetak(listProyeks[idx].Name, listProyeks[idx].Kategori, listProyeks[idx].Donasi, listProyeks[idx].TotalDonatur, listProyeks[idx].TargetDonasi)
		} else {
			fmt.Println("Proyek tidak ditemukan.")
		}
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func sequentialSearchByName(keyword string) {
	found := false
	keyword = strings.ToLower(keyword)
	for i := 0; i < currentIndex; i++ {
		nama := strings.ToLower(listProyeks[i].Name)
		if nama == keyword {
			cetak(listProyeks[i].Name, listProyeks[i].Kategori, listProyeks[i].Donasi, listProyeks[i].TotalDonatur, listProyeks[i].TargetDonasi)
			found = true
		}
	}
	if !found {
		fmt.Println("Proyek tidak ditemukan.")
	}
}

func sequentialSearchByKategori(kategori string) {
	found := false
	kategori = strings.ToLower(kategori)
	for i := 0; i < currentIndex; i++ {
		kat := strings.ToLower(listProyeks[i].Kategori)
		if kat == kategori {
			cetak(listProyeks[i].Name, listProyeks[i].Kategori, listProyeks[i].Donasi, listProyeks[i].TotalDonatur, listProyeks[i].TargetDonasi)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada proyek dengan kategori tersebut.")
	}
}

func binarySearchByName(keyword string) int {
	var left, right int
	left = 0
	right = currentIndex - 1

	for left <= right {
		var mid int
		mid = (left + right) / 2
		if listProyeks[mid].Name == keyword {
			return mid
		} else if listProyeks[mid].Name < keyword {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func binarySearchByKategori(keyword string) int {
	var left, right int
	left = 0
	right = currentIndex - 1

	for left <= right {
		var mid int
		mid = (left + right) / 2
		if listProyeks[mid].Kategori == keyword {
			return mid
		} else if listProyeks[mid].Kategori < keyword {
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

func InsertionSortByKategori(projects []proyek, n int) {
	for i := 1; i < n; i++ {
		key := projects[i]
		j := i - 1
		for j >= 0 && projects[j].Kategori > key.Kategori {
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

	tampilkanSemuaProyek()
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
	fmt.Println("8. Donasi ke Proyek")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih menu: ")
}

func cetak(name string, kategori string, totalDonasi float64, totalDonatur int, targetDonasi float64) {
	fmt.Println("===================================")
	fmt.Println("Nama Proyek   :", name)
	fmt.Println("Kategori      :", strings.Title(kategori))
	fmt.Printf("Total Donasi  : Rp%.0f\n", totalDonasi)
	fmt.Printf("Target Donasi : Rp%.0f\n", targetDonasi)
	fmt.Println("Jumlah Donatur:", totalDonatur, "orang")
	fmt.Println("===================================")
}

func tampilkanSemuaProyek() {
	if currentIndex == 0 {
		fmt.Println("Belum ada proyek yang terdaftar.")
		return
	}
	for i := 0; i < currentIndex; i++ {
		p := listProyeks[i]
		cetak(p.Name, p.Kategori, p.Donasi, p.TotalDonatur, p.TargetDonasi)
	}
}

func tampilkanProyekSukses() {
	found := false
	fmt.Println("=== Proyek yang Mencapai Target Donasi ===")
	for i := 0; i < currentIndex; i++ {
		if listProyeks[i].Donasi >= listProyeks[i].TargetDonasi {
			cetak(listProyeks[i].Name, listProyeks[i].Kategori, listProyeks[i].Donasi, listProyeks[i].TotalDonatur, listProyeks[i].TargetDonasi)
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
			editProyek()
		case 3:
			deleteProyek()
		case 4:
			searchProject()
		case 5:
			sortingMenu()
		case 6:
			tampilkanProyekSukses()
		case 7:
			tampilkanSemuaProyek()
		case 8:
			var name string
			var donasi float64
			donasiProyek(&name, donasi)
		case 0:
			fmt.Println("Terima kasih telah menggunakan sistem crowdfunding")
			return
		default:
			fmt.Println("Menu tidak valid")
		}
	}
}
