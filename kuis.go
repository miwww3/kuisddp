// ----------START PROGRAM------------
package main 
import (
	"fmt"
	"os"
	"bufio"
)

//------------STRUCT-----------------
type kuis struct{ // pada tipe data kuis terdapat dua properti yaitu pertanyaan dan jawaban yang bertipe data string 
	pertanyaan string
	jawaban string
}

func main (){

	//----------------------ARRAY------------------------
	// pada pemodelan soal kuis ini, saya hanya menggunakan empat model soal dengan pilihan opsi jawaban yang sama.
	varKuis := [4]kuis{ //tipe data kuis diberi nama varKuis (variabelnya kuis) supaya mudah dibedakan, yang mana format data untuk kuis nya adalah {"pertanyaan", "jawaban"}
		{"Yang merupakan member JKT 48?", "Freya"},
		{"Yang merupakan Presiden Indonesia?", "Jokowi"},
		{"Yang merupakan penyanyi muda di Indonesia?", "Nadin Amizah"},
		{"Yang merupakan kalimat tren di tiktok?", "Agus Takut Mah"},
	}

	//memasukkan kode membaca spasi string atau kita sebut burung beo
	reader := bufio.NewReader(os.Stdin)

	//mendeklarasikan variabel nama, score, jawabanBenar,dan jawabanSalah serta memasukkan nilai dari variabel score, jawabanBenar , dan jawabanSalah
	var nama string 
	var score, jawabanBenar, jawabanSalah int
	score = 0
	jawabanBenar = 0
	jawabanSalah = 0

	// meminta untuk menginputkan nama, yang mana hasil inputan akan disimpan di variabel "nama"
	fmt.Print("Input Nama : ")
	nama, _ = reader.ReadString('\n')
	fmt.Println("") //Baris kosong sebagai pemisah


	//----------------------SOAL-------------------------
	for i := 0; i<4; i++{ // diulang sebanyak jumlah pertanyaan yang ada di kuis yaitu 4
		var js int //variabel js adalah variabel yang akan menyimpan jawaban yang dimasukkann oleh user 
		fmt.Println(varKuis[i].pertanyaan)
		for j := 0; j<4; j++{ /// diulang sebanyak jumlah opsi pilihan jawaban yang ada yaitu 4
			fmt.Println(j,". ", varKuis[j].jawaban) // akan mengeluarkan opsi pilihan jawaban yang ada(ex: 0. Freya)
		}
		fmt.Print("Jawaban (0,1,2,3) : ") // opsi pilihan angka yang dapat dipilih untuk memilih jawaban 
		fmt.Scanln(&js) //jawaban dimasukkan dan disimpan

		if varKuis[i].jawaban == varKuis[js].jawaban{ //jika nilai variabel kuis ke-i sama dengan nilai variabel kuis ke-js (opsi pilihan angka jawaban yang kita pilih) 
			// maka nilai dari variable jawabanBenar dan Score akan bertambah satu
			jawabanBenar++
			score++
		} else if 0 <= js && js<=3 { // jika tidak sama maka nilai dari variable jawabanSalah bertambah satu 
			jawabanSalah++
		} else {//memastikan user menjawab sesuai dengan pilihan yang ada saja
			fmt.Println("") //Baris kosong sebagai pemisah
			fmt.Println("Dimohon untuk menjawa soalnya terlebih dahulu")
			fmt.Println("") //Baris kosong sebagai pemisah
			i--
		}
		fmt.Println("") //Baris kosong sebagai pemisah antar pertanyaan
	}


	//---------------KUIS SELESAI----------------
	fmt.Println("Kuis Selesai") //setelah selesai ngeloop maka akan ada pernyataan kuis telah selesai
	fmt.Println("") //enter
	fmt.Println("") //enter


	//----------------HASIL KUIS------------------
	fmt.Println("Statistic Kuis") //stastik Kuis
	fmt.Print("Nama 			: ", nama) // output nama yang tadi di inputkan
	fmt.Println("Score 			: ", score) // output hais score 
	fmt.Println("Jawaban Benar 		: ", jawabanBenar) // output jumlah jawaban yang telah dijawab dengan benar
	fmt.Println("Jawaban Salah 		: ", jawabanSalah) // ouput jumlh jawaban yang telah dijawan dengan salah

}