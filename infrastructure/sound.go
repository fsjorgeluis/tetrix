package infrastructure

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type SoundPlayer struct {
	initialized bool
}

func NewSoundPlayer() *SoundPlayer {
	return &SoundPlayer{}
}

func (s *SoundPlayer) Init(sampleRate beep.SampleRate) {
	if s.initialized {
		return
	}
	err := speaker.Init(sampleRate, sampleRate.N(time.Second/10))
	if err != nil {
		log.Printf("failed to init speaker: %v", err)
		return
	}
	s.initialized = true
}

func (s *SoundPlayer) PlayMusic(path string) {
	go func() {
		fullPath := assetPath(path)
		f, err := os.Open(fullPath)
		if err != nil {
			log.Printf("error opening music: %v", err)
			return
		}
		defer f.Close()

		streamer, format, err := mp3.Decode(f)
		if err != nil {
			log.Printf("error decoding music: %v", err)
			return
		}
		s.Init(format.SampleRate)

		// loop infinito
		speaker.Play(beep.Loop(-1, streamer))

		// keeps the goroutine alive
		select {}
	}()
}

func (s *SoundPlayer) PlayEffect(path string) {
	go func() {
		fullPath := assetPath(path)
		f, err := os.Open(fullPath)
		if err != nil {
			log.Printf("error opening effect: %v", err)
			return
		}
		defer f.Close()

		streamer, format, err := mp3.Decode(f)
		if err != nil {
			log.Printf("error decoding effect: %v", err)
			return
		}
		defer streamer.Close()

		s.Init(format.SampleRate)

		done := make(chan struct{})
		speaker.Play(beep.Seq(streamer, beep.Callback(func() {
			close(done)
		})))
		<-done
	}()
}
