package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	//"strconv"
)

// Struktur Recipe untuk menyimpan detail resep
type Recipe struct {
	Nama        string
	Bahan       []Bahan
	Langkah     []string
}

// Struktur Bahan untuk menyimpan detail bahan resep
type Bahan struct {
	Nama     string
	Satuan   string
}

var resep []Recipe

func main() {
	loadData() // Memuat resep yang ada dari file saat aplikasi dimulai

	for {
		tampilkanMenu()
		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tambahResep()
		case 2:
			cariResep()
		case 3:
			hapusResep()
		case 4:
			simpanData() // Menyimpan resep ke file sebelum keluar
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func tampilkanMenu() {
	fmt.Println("\nProgram Resep Makanan")
	fmt.Println("=====================")
	fmt.Println("1. Tambah Resep")
	fmt.Println("2. Cari Resep")
	fmt.Println("3. Hapus Resep")
	fmt.Println("4. Keluar")
	fmt.Print("Pilih Menu [1-4]: ")
}

func tambahResep() {
	var namaResep string
	var bahan []Bahan
	var langkah []string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Nama Resep: ")
	scanner.Scan()
	namaResep = scanner.Text()

	fmt.Println("Masukkan bahan-bahan (tekan garis bawah untuk menghentikan penambahan bahan):")
	for {
		var namaBahan, satuan string

		fmt.Print("Bahan: ")
		scanner.Scan()
		namaBahan = scanner.Text()
		if namaBahan == "_" {
			break
		}

		fmt.Print("Satuan: ")
		scanner.Scan()
		satuan = scanner.Text()

		bahanBaru := Bahan{Nama: namaBahan, Satuan: satuan}
		bahan = append(bahan, bahanBaru)
	}

	fmt.Println("Masukkan langkah-langkah (tekan garis bawah untuk menghentikan penambahan langkah):")
	for {
		var langkahBaru string
		fmt.Print("Langkah: ")
		scanner.Scan()
		langkahBaru = scanner.Text()
		if langkahBaru == "_" {
			break
		}
		langkah = append(langkah, langkahBaru)
	}

	resepBaru := Recipe{Nama: namaResep, Bahan: bahan, Langkah: langkah}
	resep = append(resep, resepBaru)
	fmt.Println("Resep berhasil ditambahkan!")
}


func cariResep() {
	var cariNama string
	fmt.Print("Masukkan nama resep: ")
	fmt.Scanln(&cariNama)

	ditemukan := false
	for _, resep := range resep {
		if resep.Nama == cariNama {
			ditemukan = true
			fmt.Println("Resep ditemukan!")
			fmt.Println("Nama Resep:", resep.Nama)
			fmt.Println("Bahan-bahan:")
			for i, bahan := range resep.Bahan {
				fmt.Printf("%d. %s: - %s\n", i+1, bahan.Nama, bahan.Satuan)
			}
			fmt.Println("Langkah-langkah:")
			for i, langkah := range resep.Langkah {
				fmt.Printf("%d. %s\n", i+1, langkah)
			}
			break
		}
	}
	if !ditemukan {
		fmt.Println("Resep tidak ditemukan.")
	}
}


func hapusResep() {
    fmt.Println("Pilih Resep yang akan dihapus:")
    for i, resep := range resep {
        fmt.Printf("%d. %s\n", i+1, resep.Nama)
    }

    var pilihanHapus int
    fmt.Print("Pilih Resep [1-", len(resep), "]: ")
    fmt.Scanln(&pilihanHapus)
    if pilihanHapus < 1 || pilihanHapus > len(resep) {
        fmt.Println("Pilihan tidak valid")
        return
    }

    resep = append(resep[:pilihanHapus-1], resep[pilihanHapus:]...)
    fmt.Println("Resep berhasil dihapus!")
}

func simpanData() {
    file, err := os.Create("resep.txt")
    if err != nil {
        fmt.Println("Error menyimpan data:", err)
        return
    }
    defer file.Close()

    for _, resep := range resep {
        fmt.Fprintf(file, "Nama Resep: %s\nBahan: ", resep.Nama)
        for i, bahan := range resep.Bahan {
            fmt.Fprintf(file, "%s, %s", bahan.Nama, bahan.Satuan)
            if i != len(resep.Bahan)-1 {
                fmt.Fprint(file, " ")
            }
        }
        fmt.Fprint(file, "\nLangkah: ")
        for i, langkah := range resep.Langkah {
            fmt.Fprintf(file, "%s", langkah)
            if i != len(resep.Langkah)-1 {
                fmt.Fprint(file, ". ")
            }
        }
        fmt.Fprintln(file, "\n") // Baris kosong antara setiap resep
    }
}

func loadData() {
    file, err := os.Open("resep.txt")
    if err != nil {
        fmt.Println("File tidak ditemukan. Memulai dengan daftar resep kosong.")
        return
    }
    defer file.Close()

    var resepTemp Recipe
    var bahanTemp Bahan
    var langkahTemp string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.HasPrefix(line, "Nama Resep:") {
            resepTemp.Nama = strings.TrimSpace(strings.TrimPrefix(line, "Nama Resep:"))
        } else if strings.HasPrefix(line, "Bahan:") {
            parts := strings.Fields(line)[1:]
            for i := 0; i < len(parts); i += 3 {
                bahanTemp = Bahan{} // Buat objek bahanTemp yang baru

                bahanTemp.Nama = parts[i]
                bahanTemp.Satuan = parts[i+1]
                resepTemp.Bahan = append(resepTemp.Bahan, bahanTemp)
            }
        } else if strings.HasPrefix(line, "Langkah:") {
            langkahTemp = strings.TrimSpace(strings.TrimPrefix(line, "Langkah:"))
            resepTemp.Langkah = strings.Split(langkahTemp, ". ")
        } else if line == "" {
            resep = append(resep, resepTemp)
            resepTemp = Recipe{}
        }
    }
    if err := scanner.Err(); err != nil {
        fmt.Println("Error membaca file:", err)
        return
    }
}

