package main

import "fmt"

type Skripsi struct {
	ID         int
	Judul      string
	Penulis    string
	Tahun      int
	Pembimbing string
	Topik      string
	Status     string
}

var arsip []Skripsi
var nextID = 4

func main() {
	arsip = append(arsip, Skripsi{1, "Aplikasi_Chatbot", "Caca", 2024, "Pak_Budi", "AI", "Lulus"})
	arsip = append(arsip, Skripsi{2, "Sistem_Keamanan", "Andi", 2023, "Bu_Siti", "Keamanan", "Lulus"})
	arsip = append(arsip, Skripsi{3, "Web_Kampus", "Budi", 2024, "Pak_Anton", "Web", "Revisi"})

	var pilihan int

	for {
		fmt.Println("\n=== MENU SKRIPSIN ===")
		fmt.Println("1. Tampil Data")
		fmt.Println("2. Tambah Data ")
		fmt.Println("3. Ubah Data  ")
		fmt.Println("4. Hapus Data")
		fmt.Println("5. Cari Judul")
		fmt.Println("6. Cari Penulis")
		fmt.Println("7. Urut Nama")
		fmt.Println("8. Urut Tahun")
		fmt.Println("9. Statistik")
		fmt.Println("0. Keluar")
		fmt.Println("==========================================")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		if pilihan == 1 {
			tampilData()
		} else if pilihan == 2 {
			tambahData()
		} else if pilihan == 3 {
			ubahData()
		} else if pilihan == 4 {
			hapusData()
		} else if pilihan == 5 {
			cariSequential()
		} else if pilihan == 6 {
			cariBinary()
		} else if pilihan == 7 {
			urutSelection()
			fmt.Println("Data diurutkan berdasarkan Nama (A-Z)!")
			tampilData()
		} else if pilihan == 8 {
			urutInsertion()
			fmt.Println("Data diurutkan berdasarkan Tahun (Lama-Baru)!")
			tampilData()
		} else if pilihan == 9 {
			statistik()
		} else if pilihan == 0 {
			fmt.Println("Program Selesai.")
			break
		} else {
			fmt.Println("Pilihan salah!")
		}
	}
}

func tampilData() {
	fmt.Println("\n--- Data Skripsi ---")
	for i := 0; i < len(arsip); i++ {
		fmt.Printf(" %d | Judul: %s | Penulis: %s (%d) | Pembimbing: %s | Topik: %s | Status: %s\n", arsip[i].ID, arsip[i].Judul, arsip[i].Penulis, arsip[i].Tahun, arsip[i].Pembimbing, arsip[i].Topik, arsip[i].Status)
	}
}

func tambahData() {
	var judul, penulis, pembimbing, topik, status string
	var tahun int

	fmt.Println("\n--- Tambah Data (Gunakan '_' untuk spasi) ---")
	fmt.Print("Judul      : ")
	fmt.Scanln(&judul)
	fmt.Print("Penulis    : ")
	fmt.Scanln(&penulis)
	fmt.Print("Tahun      : ")
	fmt.Scanln(&tahun)
	fmt.Print("Pembimbing : ")
	fmt.Scanln(&pembimbing)
	fmt.Print("Topik      : ")
	fmt.Scanln(&topik)
	fmt.Print("Status     : ")
	fmt.Scanln(&status)

	arsip = append(arsip, Skripsi{nextID, judul, penulis, tahun, pembimbing, topik, status})
	nextID++
	fmt.Println("Data berhasil ditambah!")
}

func ubahData() {
	var idTarget int
	fmt.Print("Masukkan ID Skripsi yang mau diubah: ")
	fmt.Scanln(&idTarget)

	for i := 0; i < len(arsip); i++ {
		if arsip[i].ID == idTarget {
			fmt.Println("Data ditemukan! Masukkan data baru (Gunakan '_' untuk spasi):")
			fmt.Print("Judul Baru      : ")
			fmt.Scanln(&arsip[i].Judul)
			fmt.Print("Penulis Baru    : ")
			fmt.Scanln(&arsip[i].Penulis)
			fmt.Print("Tahun Baru      : ")
			fmt.Scanln(&arsip[i].Tahun)
			fmt.Print("Pembimbing Baru : ")
			fmt.Scanln(&arsip[i].Pembimbing)
			fmt.Print("Topik Baru      : ")
			fmt.Scanln(&arsip[i].Topik)
			fmt.Print("Status Baru     : ")
			fmt.Scanln(&arsip[i].Status)
			fmt.Println("Data berhasil diubah!")
			return
		}
	}
	fmt.Println("ID tidak ditemukan.")
}

func hapusData() {
	var idTarget int
	fmt.Print("Masukkan ID Skripsi yang mau dihapus: ")
	fmt.Scanln(&idTarget)

	for i := 0; i < len(arsip); i++ {
		if arsip[i].ID == idTarget {
	
			arsip = append(arsip[:i], arsip[i+1:]...)
			fmt.Println("Data berhasil dihapus!")
			return
		}
	}
	fmt.Println("ID tidak ditemukan.")
}

func cariSequential() {
	var target string
	fmt.Print("Masukkan Judul yang dicari (sama persis): ")
	fmt.Scanln(&target)

	ketemu := false
	for i := 0; i < len(arsip); i++ {
		if arsip[i].Judul == target {
			fmt.Println("=> Ditemukan! Judul:", arsip[i].Judul, "oleh", arsip[i].Penulis)
			ketemu = true
		}
	}
	if !ketemu {
		fmt.Println("Judul tidak ditemukan.")
	}
}

func cariBinary() {
	urutSelection() 

	var target string
	fmt.Print("Masukkan Nama Penulis yang dicari (sama persis): ")
	fmt.Scanln(&target)

	low := 0
	high := len(arsip) - 1
	ketemu := false

	for low <= high {
		mid := (low + high) / 2
		if arsip[mid].Penulis == target {
			fmt.Println("=> Ditemukan! Penulis:", arsip[mid].Penulis, "| Judul:", arsip[mid].Judul)
			ketemu = true
			break
		} else if arsip[mid].Penulis < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if !ketemu {
		fmt.Println("Penulis tidak ditemukan.")
	}
}

func urutSelection() {
	n := len(arsip)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arsip[j].Penulis < arsip[minIdx].Penulis {
				minIdx = j
			}
		}
		temp := arsip[i]
		arsip[i] = arsip[minIdx]
		arsip[minIdx] = temp
	}
}

func urutInsertion() {
	n := len(arsip)
	for i := 1; i < n; i++ {
		key := arsip[i]
		j := i - 1

		for j >= 0 && arsip[j].Tahun > key.Tahun {
			arsip[j+1] = arsip[j]
			j = j - 1
		}
		arsip[j+1] = key
	}
}

func statistik() {
	fmt.Println("\n--- Statistik Skripsi ---")
	fmt.Println("Total Dokumen Tersimpan:", len(arsip))

	if len(arsip) == 0 {
		return
	}

	urutInsertion()

	tahunSekarang := arsip[0].Tahun
	jumlah := 0

	fmt.Println("Jumlah per Tahun Lulus:")
	for i := 0; i < len(arsip); i++ {
		if arsip[i].Tahun == tahunSekarang {
			jumlah++
		} else {
			fmt.Println("- Tahun", tahunSekarang, ":", jumlah, "dokumen")
			tahunSekarang = arsip[i].Tahun
			jumlah = 1
		}
	}
	fmt.Println("- Tahun", tahunSekarang, ":", jumlah, "dokumen")
}