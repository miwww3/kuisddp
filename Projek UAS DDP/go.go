package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Check if there are any arguments passed to the program
	// * os.Args is an array of string
	// the index 0 is the file name, so we need to start from index 1
	var args = os.Args[1:]
	var argsLength = len(args)

	if argsLength > 4 {
		fmt.Print("Argumen yang anda masukkan terlalu banyak\n")
		return
	}

	if argsLength != 0 {
		switch strings.ToLower(args[0]) {
		case "tambah":
			insertData(args[1], args[2], args[3])
			return
		case "lihat":
			lihatBarang()
			return
		case "cari":
			searchBarang(args[1])
			return
		case "tentang":
			tentangAplikasi()
			return
		default:
			fmt.Print("\n\n!! Argumen yang anda masukkan tidak tersedia !!\n\n")
		}
	}

	displayMenu()
}


// --------------------------------------------------------------------------------------------

func displayMenu() {

	scanner := bufio.NewScanner(os.Stdin)
	
	listMenu()

	scanner.Scan()

	switch scanner.Text() {
		case "1":
			inputBarang()
		case "2":
			lihatBarang()
		case "3":
			cariBarang()
		case "4":
			tentangAplikasi()
		default:
			fmt.Println("Pilihan tidak tersedia")
	}
}

func listMenu() {
	fmt.Println("┌────────────────────────────────────────────────┐")

    // Print the header
    fmt.Println("│                                                │")
    fmt.Println("│   Repository Barang PT Maju Kena Mundur Kena   │")
    fmt.Println("│                                                │")
    fmt.Println("│                                                │")

    // Print the menu options
    fmt.Println("│   Masukan pilihan Anda:                        │")
    fmt.Println("│                                                │")
    fmt.Println("│   1. Input Barang                              │")
    fmt.Println("│   2. Lihat Daftar Barang                       │")
    fmt.Println("│   3. Cari Barang                               │")
    fmt.Println("│   4. Tentang Aplikasi                          │")
    fmt.Println("│                                                │")
    fmt.Println("└────────────────────────────────────────────────┘")
    fmt.Print("> Pilih menu: ")
}

// --------------------------------------------------------------------------------------------
func cariBarang() {
	clearScreen()
	
	fmt.Println("┌───────────────────────────────────────────────┐")
	fmt.Println("│                                               │")
	fmt.Println("│   Cari Barang                                 │")
	fmt.Println("│                                               │")
	fmt.Println("└───────────────────────────────────────────────┘")
	fmt.Println("")

	scanner := bufio.NewScanner(os.Stdin)	
	fmt.Print("> Masukkan nama barang: ")
	scanner.Scan()
	query := scanner.Text()

	searchBarang(query)
}

func searchBarang(query string) {
	// Read existing data from the file
	fileContent, err := os.ReadFile("barang.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Unmarshal existing data
	var existingData []map[string]string
	err = json.Unmarshal(fileContent, &existingData)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Proses pencarian
	results := searchResult(existingData, query)

	if len(results) == 0 {
		fmt.Println("No results found.")
		return
	}

	fmt.Println("\nHasil Pencarian : ")

	for i, result := range results {
		// Convert the (int) index to string
		iteration := strconv.FormatInt(int64(i + 1), 10)
		fmt.Printf("%s. %s | %s | %s", iteration, result["namaBarang"], result["hargaBarang"], result["stokBarang"] + "\n")
	}
}

func searchResult(data []map[string]string, input string) []map[string]string {
	var results []map[string]string

	for _, item := range data {
		// Case-insensitive search using strings.Contains
		if strings.Contains(strings.ToLower(item["namaBarang"]), strings.ToLower(input)) ||
			strings.Contains(strings.ToLower(item["hargaBarang"]), strings.ToLower(input)) ||
			strings.Contains(strings.ToLower(item["stokBarang"]), strings.ToLower(input)) {
			results = append(results, item)
		}
	}

	return results
}


// --------------------------------------------------------------------------------------------

func tentangAplikasi() {
	clearScreen()

    fmt.Println("┌───────────────────────────────────────────────┐")
    fmt.Println("│                                               │")
    fmt.Println("│   Tentang Aplikasi                            │")
    fmt.Println("│                                               │")
    fmt.Println("│   Penulis : Muhammad Nauval Faiq Khilmi       │")
    fmt.Println("│   NIM     : **********                        │")
    fmt.Println("│   Kelas   : *****                             │")
    fmt.Println("└───────────────────────────────────────────────┘")
}

//  --------------------------------------------------------------------------------------------

func lihatBarang() {
	clearScreen()

	fmt.Println("┌───────────────────────────────────────────────┐")
    fmt.Println("│                                               │")
    fmt.Println("│   Lihat Daftar Barang                         │")
    fmt.Println("│                                               │")
	fmt.Println("└───────────────────────────────────────────────┘")
	fmt.Println("")

	// Read existing data from the file
	fileContent, err := os.ReadFile("barang.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Unmarshal existing data (decode JSON to array)
	var existingData []map[string]string
	err = json.Unmarshal(fileContent, &existingData)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	for i, data := range existingData {
		iteration := strconv.FormatInt(int64(i + 1), 10)
		fmt.Printf("%s. %s | %s | %s", iteration, data["namaBarang"], data["hargaBarang"], data["stokBarang"] + "\n")
	}
}

// --------------------------------------------------------------------------------------------

func inputBarang() {
	clearScreen()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("┌───────────────────────────────────────────────┐")
    fmt.Println("│                                               │")
    fmt.Println("│   Input Barang                                │")
    fmt.Println("│                                               │")
	fmt.Println("└───────────────────────────────────────────────┘")
	fmt.Println("")

	fmt.Print("> Nama Barang: ")
	scanner.Scan()
	namaBarang := scanner.Text()

	fmt.Print("> Harga Barang: ")
	scanner.Scan()
	hargaBarang := scanner.Text()

	fmt.Print("> Stok Barang: ")
	scanner.Scan()
	stokBarang := scanner.Text()

	insertData(namaBarang, hargaBarang, stokBarang)
}


func insertData(namaBarang string, hargaBarang string, stokBarang string) {
	var data = []map[string]string{
		{
			"namaBarang": namaBarang,
			"hargaBarang": hargaBarang,
			"stokBarang": stokBarang,
		},
	}

	checkFile(data[0])
}

func checkFile(newData map[string]string) {
	if _, err := os.Stat("barang.txt"); os.IsNotExist(err) {
		// If the file doesn't exist, create it with the new data
		createNewFile(newData)
	} else {
		// If the file exists, read the existing data, append the new data, and write it back
		appendToFile(newData)
	}
}

func createNewFile(data map[string]string) {
	convertAndWrite([]map[string]string{data})
	fmt.Println("File created with new data.")
}

func appendToFile(newData map[string]string) {
	// Read existing data from the file
	fileContent, err := os.ReadFile("barang.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// check if file is empty
	if len(fileContent) == 0 {
		convertAndWrite([]map[string]string{newData})
		fmt.Println("File is empty. New data written.")
		return
	}

	// Unmarshal existing data
	var existingData []map[string]string
	err = json.Unmarshal(fileContent, &existingData)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		createNewFile(newData)
		return
	}

	// Append new data
	existingData = append(existingData, newData)

	convertAndWrite(existingData)

	fmt.Println("\nData berhasil disimpan.")
}

func convertAndWrite(data []map[string]string) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Write JSON data back to the file
	err = os.WriteFile("barang.txt", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}