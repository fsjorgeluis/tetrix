package _interface

type InputAction int

const (
	MoveLeft InputAction = iota
	MoveRight
	RotateCW
	RotateCCW
	SoftDrop
	HardDrop
	Quit
	NoAction
	Restart
)

type InputHandler interface {
	Poll() InputAction
}
