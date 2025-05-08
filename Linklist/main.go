package main

import (
	"LINKLIST/model"
	"LINKLIST/node"
	"LINKLIST/view"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MenuBuku() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(model.ReadRak())

	for {
		fmt.Println("\n=== MENU BUKU ===")
		fmt.Println("1. Tambah Data Buku")
		fmt.Println("2. Tampilkan Data Buku")
		fmt.Println("3. Update Data Buku")
		fmt.Println("4. Hapus Data Buku")
		fmt.Println("5. Logout")
		fmt.Print("Pilih menu: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Input tidak valid, silakan masukkan angka.")
			continue
		}

		switch choice {
		case 1:
			fmt.Println("Anda memilih: Tambah Data Buku")
			view.Insert()
		case 2:
			fmt.Println("Anda memilih: Tampilkan Data Buku")
			view.Views()
		case 3:
			fmt.Println("Anda memilih: Update Data Buku")
			view.Update()
		case 4:
			fmt.Println("Anda memilih: Hapus Data Buku")
			view.Delete()
		case 5:
			fmt.Println("Logout berhasil.")
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}

func MenuUtama() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== MENU UTAMA ===")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih menu: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Input tidak valid, silakan masukkan angka.")
			continue
		}

		switch choice {
		case 1:
			fmt.Println("Anda memilih: Register")
			view.Register()
		case 2:
			fmt.Println("Anda memilih: Login")
			if view.Login() {
				MenuBuku()
			} else {
				fmt.Println("Login gagal, coba lagi.")
			}
		case 3:
			fmt.Println("Terima kasih! Program selesai.")
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}

func main() {

	rak1 := node.RakBuku{
		NomorRak: 1,
		TipeBuku: "Self Development",
		Baris:    2,
	}

	rak2 := node.RakBuku{
		NomorRak: 2,
		TipeBuku: "Pendidikan",
		Baris:    3,
	}

	rak3 := node.RakBuku{
		NomorRak: 3,
		TipeBuku: "Kesehatan",
		Baris:    4,
	}

	rak4 := node.RakBuku{
		NomorRak: 4,
		TipeBuku: "Sejarah",
		Baris:    4,
	}

	model.CreateRak(rak1)
	model.CreateRak(rak2)
	model.CreateRak(rak3)
	model.CreateRak(rak4)

	MenuUtama()
}
