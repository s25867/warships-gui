package gui

type Spatial interface {
	Size() (int, int)
	Position() (int, int)
}
