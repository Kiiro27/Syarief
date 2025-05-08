package view

import (
	"LINKLIST/model"
	"LINKLIST/node"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Insert() {
	var kota, penerbit, harga, halaman, namabuku string
	var idbuku, nomor_rak, tahun int

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan ID Buku : ")
	fmt.Scan(&idbuku)
	reader.ReadString('\n')

	fmt.Print("Masukkan Nama Buku : ")
	namabuku, _ = reader.ReadString('\n')
	namabuku = strings.TrimSpace(namabuku)

	fmt.Print("Masukkan Nama Penerbit : ")
	penerbit, _ = reader.ReadString('\n')
	penerbit = strings.TrimSpace(penerbit)

	fmt.Print("Masukkan Tahun Buku : ")
	fmt.Scan(&tahun)
	reader.ReadString('\n')

	fmt.Print("Masukkan Kota Terbit : ")
	kota, _ = reader.ReadString('\n')
	kota = strings.TrimSpace(kota)

	fmt.Print("Masukkan Halaman Buku : ")
	halaman, _ = reader.ReadString('\n')
	halaman = strings.TrimSpace(halaman)

	fmt.Print("Masukkan Harga Buku : ")
	harga, _ = reader.ReadString('\n')
	harga = strings.TrimSpace(harga)

	fmt.Print("Masukkan Nomor Rak : ")
	fmt.Scan(&nomor_rak)
	reader.ReadString('\n')

	buku := node.Buku{
		ID:       idbuku,
		NamaBuku: namabuku,
		Penerbit: node.Terbit{
			NamaPenerbit: penerbit,
			Tahun:        tahun,
			Kota:         kota,
		},
		Halaman:    halaman,
		Harga:      harga,
		Rak_Number: nomor_rak,
	}

	cekRakBuku := model.SearchRakBuku(nomor_rak)
	if cekRakBuku {
		model.CreateBuku(buku)
		fmt.Println("==== Buku Berhasil Ditambahkan ====")
	} else {
		fmt.Println("== Nomor Rak tidak ada Di Database ==")
	}

	fmt.Println()
}

func Views() {
	fmt.Println("Daftar Buku")
	for i, emp := range model.ReadBuku() {
		fmt.Println("===== Buku ke - ", i+1, " =====")
		fmt.Println("ID Buku\t\t: ", emp.ID)
		fmt.Println("Nama Buku\t: ", emp.NamaBuku)
		fmt.Println("============================================")
		fmt.Println("Penerbit\t\t: ")
		fmt.Println("Nama Penerbit\t: ", emp.Penerbit.NamaPenerbit)
		fmt.Println("Kota Terbit\t: ", emp.Penerbit.Kota)
		fmt.Println("Tahun Terbit\t: ", emp.Penerbit.Tahun)
		fmt.Println("============================================")
		fmt.Println("Halaman Buku\t: ", emp.Halaman)
		fmt.Println("Harga buku\t: ", emp.Harga)
		fmt.Println("Rak \t\t: ", model.GetRakBuku(emp.Rak_Number))
		fmt.Println("============================================")
		fmt.Println()
	}
}

func Update() {
	var idbuku, tahun int
	var namabuku, kota, halaman, harga, penerbit string

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan ID Buku yang akan diUpdate : ")
	fmt.Scan(&idbuku)
	reader.ReadString('\n')

	fmt.Print("Masukkan Nama Buku : ")
	namabuku, _ = reader.ReadString('\n')
	namabuku = strings.TrimSpace(namabuku)

	fmt.Print("Masukkan Tahun Terbit : ")
	fmt.Scan(&tahun)
	reader.ReadString('\n')

	fmt.Print("Masukkan Nama Penerbit : ")
	penerbit, _ = reader.ReadString('\n')
	penerbit = strings.TrimSpace(penerbit)

	fmt.Print("Masukkan Kota Terbit : ")
	fmt.Scan(&kota)
	reader.ReadString('\n')

	fmt.Print("Masukkan Halaman Buku : ")
	fmt.Scan(&halaman)
	reader.ReadString('\n')

	fmt.Print("Masukkan Harga Buku : ")
	fmt.Scan(&harga)
	reader.ReadString('\n')

	buku := node.Buku{
		ID:       idbuku,
		NamaBuku: namabuku,
		Penerbit: node.Terbit{
			NamaPenerbit: penerbit,
			Tahun:        tahun,
			Kota:         kota,
		},
		Halaman: halaman,
		Harga:   harga,
	}
	cek := model.UpdateBuku(buku, idbuku)
	if cek {
		fmt.Println("=== Buku Berhasil Update ===")
	} else {
		fmt.Println("=== Pegawai Gagal DiUpdate ===")
	}
	fmt.Println()
}

func Delete() {
	var id int
	fmt.Println("Masukkan ID Buku yang ingin dihapus : ")
	fmt.Scan(&id)

	cek := model.DeleteBuku(id)

	if cek {
		fmt.Println("=== Buku Berhasil DiHapus ===")
	} else {
		fmt.Println("== Buku Gagal Dihapus ==")
	}
	fmt.Println()
}
