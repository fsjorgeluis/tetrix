package infrastructure

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type SoundPlayer struct {
	initialized bool
	cache       map[string]*CachedSound
	mu          sync.RWMutex
}
type CachedSound struct {
	buffer *beep.Buffer
	format beep.Format
}

func NewSoundPlayer() *SoundPlayer {
	return &SoundPlayer{
		cache: make(map[string]*CachedSound),
	}
}

func (s *SoundPlayer) Init(sampleRate beep.SampleRate) {
	if s.initialized {
		return
	}
	if err := speaker.Init(sampleRate, sampleRate.N(time.Second/10)); err != nil {
		log.Printf("failed to init speaker: %v", err)
		return
	}
	s.initialized = true
}

// Preload loads and decodes a sound into memory
func (s *SoundPlayer) Preload(path string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.cache[path]; ok {
		return nil
	}

	fullPath := assetPath(path)
	f, err := os.Open(fullPath)
	if err != nil {
		return err
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		return fmt.Errorf("decode sound: %w", err)
	}
	defer streamer.Close()
	s.Init(format.SampleRate)
	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)

	s.cache[path] = &CachedSound{
		buffer: buffer,
		format: format,
	}

	return nil
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
		defer streamer.Close()

		s.Init(format.SampleRate)
		speaker.Play(beep.Loop(-1, streamer))

		select {} // keep alive
	}()
}

func (s *SoundPlayer) PlayEffect(path string) {
	s.mu.RLock()
	cached := s.cache[path]
	s.mu.RUnlock()

	if cached == nil {
		// fallback: load on demand
		if err := s.Preload(path); err != nil {
			log.Printf("failed to preload effect: %v", err)
			return
		}
		s.mu.RLock()
		cached = s.cache[path]
		s.mu.RUnlock()
	}

	if cached == nil {
		return
	}

	streamer := cached.buffer.Streamer(0, cached.buffer.Len())

	done := make(chan struct{})
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		close(done)
	})))
	go func() {
		<-done
	}()
}

func (s *SoundPlayer) Close() {
	s.mu.Lock()
	defer s.mu.Unlock()
	for k := range s.cache {
		delete(s.cache, k)
	}
}
