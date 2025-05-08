package model

import "LINKLIST/node"

var DaftarRakBuku node.ListRak

func CreateRak(rak node.RakBuku) bool {
	tempLL := node.ListRak{
		Data: rak,
		Link: nil,
	}
	if DaftarRakBuku.Link == nil {
		DaftarRakBuku.Link = &tempLL
		return true
	} else {
		temp := &DaftarRakBuku
		for temp.Link != nil {
			temp = temp.Link
		}
		temp.Link = &tempLL
		return true
	}
	return false
}

func ReadRak() []node.RakBuku {
	daftarRak := []node.RakBuku{}
	temp := &DaftarRakBuku
	for temp.Link != nil {
		daftarRak = append(daftarRak, temp.Link.Data)
		temp = temp.Link
	}
	return daftarRak
}

func DeleteRakBuku(id int) bool {
	temp := &DaftarRakBuku
	for temp.Link != nil {
		if temp.Link.Data.NomorRak == id {
			temp.Link = temp.Link.Link
			return true
		}
		temp = temp.Link
	}
	return false
}

func UpdateRakBuku(rak node.RakBuku, id int) bool {
	temp := DaftarRakBuku.Link
	for temp != nil {
		if temp.Data.NomorRak == id {
			temp.Data = rak
			return true
		}
		temp = temp.Link
	}
	return false
}

func SearchRakBuku(id int) bool {
	temp := &DaftarRakBuku
	for temp.Link != nil {
		if temp.Link.Data.NomorRak == id {
			return true
		}
		temp = temp.Link
	}
	return false
}

func GetRakBuku(id int) string {
	temp := &DaftarRakBuku
	for temp.Link != nil {
		if temp.Link.Data.NomorRak == id {
			return temp.Link.Data.TipeBuku
		}
		temp = temp.Link
	}
	return ""
}
