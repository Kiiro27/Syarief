package view

import (
	"LINKLIST/model"
	"LINKLIST/node"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Register() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("=== Registrasi User Baru ===")
	for {
		fmt.Print("Masukkan Username: ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)
		if username == "" {
			fmt.Println("Username tidak boleh kosong, coba lagi.")
			continue
		}
		if model.UserExists(username) {
			fmt.Println("Username sudah ada, coba yang lain.")
			continue
		}

		fmt.Print("Masukkan Password: ")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(password)
		if password == "" {
			fmt.Println("Password tidak boleh kosong, coba lagi.")
			continue
		}

		admin := node.Admin{
			Username: username,
			Password: password,
		}
		if model.CreateAdmin(admin) {
			fmt.Println("Registrasi berhasil!")
		} else {
			fmt.Println("Registrasi gagal, coba lagi.")
		}
		break
	}
}

func Login() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("=== Login ===")
	fmt.Print("Masukkan Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Masukkan Password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	if model.Login(username, password) {
		fmt.Println("Login berhasil!")
		return true
	} else {
		fmt.Println("Username atau password salah!")
		return false
	}
}
