package model

import (
	"LINKLIST/node"
)

var DaftarAdmin node.ListAdmin

func CreateAdmin(admin node.Admin) bool {
	tempLL := node.ListAdmin{
		Data: admin,
		Link: nil,
	}

	if DaftarAdmin.Link == nil {
		DaftarAdmin.Link = &tempLL
		return true
	} else {
		temp := &DaftarAdmin
		for temp.Link != nil {
			temp = temp.Link
		}
		temp.Link = &tempLL
		return true
	}
	return false
}

func Login(username, password string) bool {
	temp := DaftarAdmin.Link
	for temp != nil {
		if temp.Data.Username == username && temp.Data.Password == password {
			return true
		}
		temp = temp.Link
	}
	return false
}

func UserExists(username string) bool {
	temp := DaftarAdmin.Link
	for temp != nil {
		if temp.Data.Username == username {
			return true
		}
		temp = temp.Link
	}
	return false
}
