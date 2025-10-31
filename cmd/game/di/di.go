package di

import (
	"log"
	"time"

	"github.com/fsjorgeluis/tetrix/infrastructure"
	_interface "github.com/fsjorgeluis/tetrix/interface"
	"github.com/fsjorgeluis/tetrix/internal/domain"
	"github.com/fsjorgeluis/tetrix/internal/usecase"
)

type GameDependencies struct {
	Board        *domain.Board
	GameService  *usecase.GameService
	Renderer     _interface.Renderer
	InputHandler *infrastructure.TCellInput
	TickChan     <-chan struct{}
	Timer        *infrastructure.Timer
	Sound        *infrastructure.SoundPlayer
}

func SetupGame() *GameDependencies {
	renderer, err := infrastructure.NewTCellRenderer()
	if err != nil {
		log.Fatalf("failed to init renderer: %v", err)
	}

	sound := infrastructure.NewSoundPlayer()
	preloadFiles := []string{
		"assets/sounds/shot.mp3",
		"assets/sounds/line_clear.mp3",
	}
	for _, file := range preloadFiles {
		if err := sound.Preload(file); err != nil {
			log.Printf("failed to preload sound %s: %v", file, err)
		}
	}

	// background music
	sound.PlayMusic("assets/sounds/t_sound.mp3")

	inputHandler := infrastructure.NewTCellInput(renderer.Screen())

	tickInterval := 500 * time.Millisecond
	timer := infrastructure.NewTimer(tickInterval)
	tickChan := make(chan struct{})
	timer.Start(tickChan)

	board, err := domain.NewBoard(10, 20)
	if err != nil {
		log.Fatalf("failed to init board: %v", err)
	}
	spawner := usecase.NewDefaultSpawner()

	gameService, err := usecase.NewGameService(board, spawner, sound)
	if err != nil {
		log.Fatalf("failed to init game service: %v", err)
	}

	return &GameDependencies{
		Board:        board,
		GameService:  gameService,
		Renderer:     renderer,
		InputHandler: inputHandler,
		TickChan:     tickChan,
		Timer:        timer,
		Sound:        sound,
	}
}

func (d *GameDependencies) Close() {
	if d.InputHandler != nil {
		d.InputHandler.Close()
	}
	if d.Renderer != nil {
		d.Renderer.Close()
	}
	if d.Timer != nil {
		d.Timer.Stop()
	}
	if d.Sound != nil {
		d.Sound.Close()
	}
}
