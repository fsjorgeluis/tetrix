package infrastructure

import (
	"time"

	_interface "github.com/fsjorgeluis/tetrix/interface"
	"github.com/gdamore/tcell/v2"
)

type TCellInput struct {
	screen  tcell.Screen
	inputCh chan _interface.InputAction
	stopCh  chan struct{}
}

func NewTCellInput(screen tcell.Screen) *TCellInput {
	ti := &TCellInput{
		screen:  screen,
		inputCh: make(chan _interface.InputAction, 16),
		stopCh:  make(chan struct{}),
	}
	go ti.readLoop()
	return ti
}

func (i *TCellInput) readLoop() {
	for {
		ev := i.screen.PollEvent()
		select {
		case <-i.stopCh:
			return
		default:
		}
		if ev == nil {
			// prevent busy-loop
			time.Sleep(5 * time.Millisecond)
			continue
		}

		action := i.eventToAction(ev)
		select {
		case i.inputCh <- action:
		default:
			// discard event, buffer is full
		}
	}
}

func (i *TCellInput) eventToAction(ev tcell.Event) _interface.InputAction {
	switch tev := ev.(type) {
	case *tcell.EventKey:
		switch tev.Key() {
		case tcell.KeyLeft:
			return _interface.MoveLeft
		case tcell.KeyRight:
			return _interface.MoveRight
		case tcell.KeyUp:
			return _interface.RotateCW
		case tcell.KeyDown:
			return _interface.SoftDrop
		case tcell.KeyCtrlC, tcell.KeyEscape:
			return _interface.Quit
		default:
			switch tev.Rune() {
			case 'a', 'A':
				return _interface.MoveLeft
			case 'd', 'D':
				return _interface.MoveRight
			case 'w', 'W':
				return _interface.RotateCW
			case 's', 'S':
				return _interface.SoftDrop
			case 'q', 'Q':
				return _interface.Quit
			case 'R', 'r':
				return _interface.Restart
			}
		}
	}
	return _interface.NoAction
}

func (i *TCellInput) Poll() _interface.InputAction {
	select {
	case a := <-i.inputCh:
		return a
	default:
		return _interface.NoAction
	}
}

func (i *TCellInput) Close() {
	close(i.stopCh)
	select {
	case <-i.inputCh:
	default:
	}
}
