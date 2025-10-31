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

	gameOver := false

	for !gameOver {
		select {
		case <-deps.TickChan:
			deps.GameService.Tick()
		default:
			action := deps.InputHandler.Poll()
			switch action {
			case _interface.MoveLeft:
				deps.GameService.MoveLeft()
			case _interface.MoveRight:
				deps.GameService.MoveRight()
			case _interface.RotateCW:
				deps.GameService.RotateCW()
			case _interface.SoftDrop:
				deps.GameService.MoveDown()
			case _interface.Quit:
				gameOver = true
			case _interface.NoAction:
				// do nothing
			default:
				log.Printf("unhandled action: %v", action)
			}
		}

		deps.Renderer.Begin(deps.Board, deps.GameService.Score())
		deps.Renderer.DrawPiece(deps.Board, deps.GameService.CurrentPiece())
		deps.Renderer.DrawBoard(deps.Board)
		deps.Renderer.Flush()

		if deps.GameService.IsGameOver() {
			gameOver = true
		}

		time.Sleep(10 * time.Millisecond)
	}

	log.Println("Game Over!")
}
