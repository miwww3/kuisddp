package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Resep struct {
	Nama    string
	Bahan   []string
	Langkah []string
}

func main() {
	fileName := "resep.txt"
	resepMakanan := bacaResepDariFile(fileName)

	menuUtama()
	pilihan := ""
	for pilihan != "4" {
		fmt.Print("Pilih opsi (1-4): ")
		fmt.Scanln(&pilihan)
		switch pilihan {
		case "1":
			tambahResep(&resepMakanan, fileName)
		case "2":
			cariResep(resepMakanan)
		case "3":
			hapusResep(&resepMakanan, fileName)
		case "4":
			fmt.Println("Terima kasih! Sampai jumpa lagi.")
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih kembali.")
		}
	}
}

func menuUtama() {
	fmt.Println("Selamat datang di Program Resep Makanan")
	fmt.Println("1. Tambah Resep Baru")
	fmt.Println("2. Cari Resep")
	fmt.Println("3. Hapus Resep")
	fmt.Println("4. Keluar")
}

func bacaResepDariFile(fileName string) map[string]Resep {
	resepMakanan := make(map[string]Resep)

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("File resep tidak ditemukan. Membuat file baru...")
		_, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Gagal membuat file baru:", err)
		}
		return resepMakanan
	}

	lines := strings.Split(string(file), "\n")

	for i := 0; i < len(lines)-2; i += 4 {
		nama := lines[i]
		bahan := strings.Split(lines[i+1], ", ")
		langkah := strings.Split(lines[i+2], ", ")

		resepMakanan[nama] = Resep{
			Nama:    nama,
			Bahan:   bahan,
			Langkah: langkah,
		}
	}

	return resepMakanan
}

func simpanResepKeFile(resepMakanan map[string]Resep, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Terjadi kesalahan saat menyimpan data resep:", err)
		return
	}
	defer file.Close()

	for _, resep := range resepMakanan {
		_, err := file.WriteString(resep.Nama + "\n")
		if err != nil {
			fmt.Println("Gagal menulis nama resep:", err)
		}
		_, err = file.WriteString(strings.Join(resep.Bahan, ", ") + "\n")
		if err != nil {
			fmt.Println("Gagal menulis bahan resep:", err)
		}
		_, err = file.WriteString(strings.Join(resep.Langkah, ", ") + "\n\n")
		if err != nil {
			fmt.Println("Gagal menulis langkah-langkah resep:", err)
		}
	}
}

func tambahResep(resepMakanan *map[string]Resep, fileName string) {
	var namaResep string
	var bahanResep string
	var langkahResep string

	fmt.Println("Tambah Resep Baru")
	fmt.Print("Nama Resep: ")
	fmt.Scanln(&namaResep)

	bahan := []string{}
	langkah := []string{}

	for {
		fmt.Print("Bahan Resep (atau '_' untuk selesai): ")
		fmt.Scanln(&bahanResep)
		if bahanResep == "_" {
			break
		}
		bahan = append(bahan, bahanResep)
	}

	for {
		fmt.Print("Langkah Resep (atau '_' untuk selesai): ")
		fmt.Scanln(&langkahResep)
		if langkahResep == "_" {
			break
		}
		langkah = append(langkah, langkahResep)
	}

	(*resepMakanan)[namaResep] = Resep{
		Nama:    namaResep,
		Bahan:   bahan,
		Langkah: langkah,
	}

	fmt.Println("Resep berhasil ditambahkan!")
	simpanResepKeFile(*resepMakanan, fileName)
}

func cariResep(resepMakanan map[string]Resep) {
	var kataKunci string
	fmt.Print("Masukkan kata kunci untuk mencari resep: ")
	fmt.Scanln(&kataKunci)

	ditemukan := false
	for _, resep := range resepMakanan {
		if strings.Contains(resep.Nama, kataKunci) {
			fmt.Println("Nama:", resep.Nama)
			fmt.Println("Bahan:")
			for _, bahan := range resep.Bahan {
				fmt.Println("- ", bahan)
			}
			fmt.Println("Langkah-langkah:")
			for _, langkah := range resep.Langkah {
				fmt.Println(langkah)
			}
			fmt.Println()
			ditemukan = true
		}
	}

	if !ditemukan {
		fmt.Println("Resep tidak ditemukan.")
	}
}

func hapusResep(resepMakanan *map[string]Resep, fileName string) {
	var pilihan string
	for pilihan != "0" {
		fmt.Println("Daftar Resep yang Ada:")
		i := 1
		for namaResep := range *resepMakanan {
			fmt.Printf("%d. %s\n", i, namaResep)
			i++
		}
		fmt.Println("0. Kembali")
		fmt.Print("Pilih nomor resep yang ingin dihapus (0 untuk kembali): ")
		fmt.Scanln(&pilihan)

		if pilihan == "0" {
			break
		}

		num, err := strconv.Atoi(pilihan)
		if err != nil || num < 1 || num > len(*resepMakanan) {
			fmt.Println("Nomor resep tidak valid. Silakan pilih nomor yang sesuai.")
			continue
		}

		i = 1
		for namaResep := range *resepMakanan {
			if i == num {
				delete(*resepMakanan, namaResep)
				fmt.Println("Resep", namaResep, "telah dihapus!")
				simpanResepKeFile(*resepMakanan, fileName)
				return
			}
			i++
		}
	}
	fmt.Println("Kembali ke menu utama.")
}
