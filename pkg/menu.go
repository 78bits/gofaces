package gofaces

type MenuItem struct {
	Group   string
	Label   string
	Iconset Iconset
	Icon    string
	Name    string
}

type Menu struct {
	Items        []MenuItem
	ItemSelected int
}
