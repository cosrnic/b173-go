package b173

import (
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/cosrnic/b173-server/pkg/player"
	"github.com/cosrnic/b173-server/pkg/world"
)

type Server struct {
	World *world.World

	listener net.Listener
}

func (s *Server) Start() (err error) {
	s.listener, err = net.Listen("tcp", "localhost:25565")
	if err != nil {
		return fmt.Errorf("failed to start listener: %w", err)
	}
	go s.listenLoop()
	go s.Ticker()

	log.Printf("started listening!")

	return nil
}

func (s *Server) listenLoop() {
	for {
		cc, err := s.listener.Accept()
		if errors.Is(err, net.ErrClosed) {
			break
		} else if err != nil {
			log.Printf("Failed to accept connection: %s\n", err)
			continue
		}

		player := player.NewPlayer(s.World, cc)
		go player.ReadLoop()

	}
}

func (s *Server) Ticker() {
	ticker := time.NewTicker(time.Second / 20)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				s.World.Tick()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
