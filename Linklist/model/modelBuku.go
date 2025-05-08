package model

import "LINKLIST/node"

var DaftarBuku node.ListBuku

func CreateBuku(emp node.Buku) bool {
	tempLL := node.ListBuku{
		Data: emp,
		Link: nil,
	}
	if DaftarBuku.Link == nil {
		DaftarBuku.Link = &tempLL
		return true
	} else {
		temp := &DaftarBuku
		for temp.Link != nil {
			temp = temp.Link
		}
		temp.Link = &tempLL
		return true
	}
	return false
}

func ReadBuku() []node.Buku {
	daftarbuku := []node.Buku{}
	temp := &DaftarBuku
	for temp.Link != nil {
		daftarbuku = append(daftarbuku, temp.Link.Data)
		temp = temp.Link
	}
	return daftarbuku
}

func UpdateBuku(emp node.Buku, id int) bool {
	temp := DaftarBuku.Link
	for temp != nil {
		if temp.Data.ID == id {
			temp.Data = emp
			return true
		}
		temp = temp.Link
	}
	return false
}

func DeleteBuku(id int) bool {
	temp := DaftarBuku.Link
	for temp != nil {
		if temp.Data.ID == id {
			temp.Link = temp.Link.Link
			return true
		}
		temp = temp.Link
	}
	return false
}
