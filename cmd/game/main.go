package main

import (
	"log"
	"time"

	"github.com/fsjorgeluis/tetrix/cmd/game/di"
	_interface "github.com/fsjorgeluis/tetrix/interface"
)

func main() {
	deps := di.SetupGame()
	defer deps.Close()

	//gameOver := false

	for /*!gameOver*/ {
		select {
		case <-deps.TickChan:
			if !deps.GameService.IsGameOver() && !deps.GameService.Paused() {
				deps.GameService.Tick()
			}
			//deps.GameService.Tick()
		default:
			action := deps.InputHandler.Poll()
			switch action {
			case _interface.MoveLeft:
				if !deps.GameService.Paused() {
					deps.GameService.MoveLeft()
				}
			case _interface.MoveRight:
				if !deps.GameService.Paused() {
					deps.GameService.MoveRight()
				}
			case _interface.RotateCW:
				if !deps.GameService.Paused() {
					deps.GameService.RotateCW()
				}
			case _interface.SoftDrop:
				if !deps.GameService.Paused() {
					deps.GameService.MoveDown()
				}
			case _interface.Restart:
				if deps.GameService.IsGameOver() {
					deps.GameService.Reset()
				}
			case _interface.Pause:
				deps.GameService.TogglePause()
			case _interface.Quit:
				return
				//gameOver = true
			case _interface.NoAction: // do nothing
			default:
				log.Printf("unhandled action: %v", action)
			}
		}

		deps.Renderer.Begin(deps.Board, deps.GameService.Score())
		deps.Renderer.DrawBoard(deps.Board)
		deps.Renderer.DrawNextPiece(deps.GameService.NextPiece(), deps.Board.Width+2, 2)

		if deps.GameService.Paused() {
			deps.Renderer.DrawLabel(&_interface.GameLabel{
				X:     (deps.Board.Width / 2) + 5,
				Y:     deps.Board.Height + 2,
				Text:  "PAUSED (ESC to resume)",
				Color: "yellow"})
		}

		if deps.GameService.IsGameOver() {
			deps.Renderer.DrawLabel(&_interface.GameLabel{
				X:     (deps.Board.Width / 2) + 1,
				Y:     deps.Board.Height + 2,
				Text:  "GAME OVER! Press 'R' to retry",
				Color: "red"})
		} else {
			deps.Renderer.DrawPiece(deps.Board, deps.GameService.CurrentPiece())
		}

		deps.Renderer.Flush()
		time.Sleep(10 * time.Millisecond)
	}
}
