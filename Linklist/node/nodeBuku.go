package node

type Terbit struct {
	NamaPenerbit string
	Tahun        int
	Kota         string
}

type Buku struct {
	ID         int
	NamaBuku   string
	Penerbit   Terbit
	Halaman    string
	Harga      string
	Rak_Number int
}

type ListBuku struct {
	Data Buku
	Link *ListBuku
}
