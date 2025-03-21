package model

import "LIST/node"

var Daftarpegawai [10]node.Pegawai
var kuota int = 10
var index = 0

func CreatePegawai(emp node.Pegawai) bool {
	if index < kuota {
		Daftarpegawai[index] = emp
		index = index + 1
		return true
	}
	return false
}

func ReadPegawai() []node.Pegawai {
	return Daftarpegawai[0:index]
}

func UpdatePegawai(emp node.Pegawai, id int) bool {
	for i := 0; i < index; i++ {
		if Daftarpegawai[i].ID == id {
			Daftarpegawai[i] = emp
			return true
		}
	}
	return false
}

func DeletePegawai(id int) bool {
	for i := 0; i < index; i++ {
		if Daftarpegawai[i].ID == id {
			Daftarpegawai[i] = Daftarpegawai[index-1]
			index = index - 1
			return true
		}
	}
	return false
}
