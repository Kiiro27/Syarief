package view

import (
	"LIST/model"
	"LIST/node"
	"bufio"
	"fmt"
	"os"
)

func Insert() {
	var City, name, notelp, Gmail string
	var id, Nomer int
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan ID Pegawai : ")
	fmt.Scan(&id)

	fmt.Print("Masukkan Nama Pegawai : ")
	name, _ = reader.ReadString('\n')
	name = name[:len(name)-1]

	fmt.Print("Masukkan Jalan : ")
	jalan, _ := reader.ReadString('\n')
	jalan = jalan[:len(jalan)-1]

	fmt.Print("Masukkan Nomor Rumah : ")
	fmt.Scan(&Nomer)

	fmt.Print("Masukkan Kota : ")
	fmt.Scan(&City)

	fmt.Print("Masukkan Nomor Telepon : ")
	fmt.Scan(notelp)

	fmt.Print("Masukkan Email : ")
	fmt.Scan(&Gmail)

	pegawai := node.Pegawai{
		ID:     id,
		Nama:   name,
		Alamat: node.Address{Jalan: jalan, Kota: City, Nomor: Nomer},
		Notelp: notelp,
		Email:  Gmail,
	}

	cek := model.CreatePegawai(pegawai)
	if cek {
		fmt.Println("=== Pegawai Berhasil Diatambahkan ===")
	} else {
		fmt.Println("Pegawai gagal ditambahkan")
	}
	fmt.Println()
}

func Views() {
	fmt.Println("Daftar Pegawai")
	for i, emp := range model.ReadPegawai() {
		fmt.Println("--- Pegawai ke - ", i+1, " ---")
		fmt.Println("ID Pegawai\t: ", emp.ID)
		fmt.Println("Nama Pegawai\t: ", emp.Nama)
		fmt.Println("Alamat\t\t: ", emp.Alamat.Jalan, emp.Alamat.Nomor, emp.Alamat.Kota)
		fmt.Println("Email\t\t: ", emp.Email)
		fmt.Println()
	}
}

func Update() {
	var Kota, nama, notelp, gmail string
	var id, Nomer int
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan ID Pegawai : ")
	fmt.Scan(&id)

	fmt.Print("Masukkan Nama Pegawai : ")
	nama, _ = reader.ReadString('\n')
	nama = nama[:len(nama)-1]

	fmt.Print("Masukkan Jalan : ")
	jalan, _ := reader.ReadString('\n')
	jalan = jalan[:len(jalan)-1]

	fmt.Print("Masukkan Nomor Rumah : ")
	fmt.Scan(&Nomer)

	fmt.Print("Masukkan Kota : ")
	fmt.Scan(&Kota)

	fmt.Print("Masukkan Nomor Telepon : ")
	fmt.Scan(notelp)

	fmt.Print("Masukkan Email : ")
	fmt.Scan(&gmail)

	pegawai := node.Pegawai{
		ID:     id,
		Nama:   nama,
		Alamat: node.Address{Jalan: jalan, Kota: Kota, Nomor: Nomer},
		Notelp: notelp,
		Email:  gmail,
	}

	cek := model.UpdatePegawai(pegawai, id)
	if cek {
		fmt.Println("=== Pegawai Berhasil diUpdate ===")
	} else {
		fmt.Println("Pegawai gagal diUpdate")
	}
	fmt.Println()
}

func Delete() {
	var id int
	fmt.Print("Masukkan ID Pegawai yang akan dihapus : ")
	fmt.Scan(&id)

	cek := model.DeletePegawai(id)
	if cek {
		fmt.Println("=== Pegawai berhasil dihapus ===")
	} else {
		fmt.Println("Pegawai gagal Dihapus")
	}
	fmt.Println()
}
