package node

type Admin struct {
	Username string
	Password string
}

type ListAdmin struct {
	Data Admin
	Link *ListAdmin
}
