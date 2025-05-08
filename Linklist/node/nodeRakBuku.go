package node

type RakBuku struct {
	NomorRak int
	TipeBuku string
	Baris    int
}

type ListRak struct {
	Data RakBuku
	Link *ListRak
}
