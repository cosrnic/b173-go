package world

import (
	"net"

	"github.com/cosrnic/b173-server/pkg/level"
	"github.com/cosrnic/b173-server/pkg/proto/packets"
	"github.com/cosrnic/b173-server/pkg/util"
)

type Player interface {
	GetConnection() net.Conn
	GetUsername() string
	SetUsername(string)
	GetWorld() *World
	SendPacket(packets.ClientboundPacket)
	GetPosition() util.Pos
	SetPosition(util.Pos)
}

type World struct {
	Players []Player
	Chunks  *[]level.Chunk
}

func (w *World) Tick() {
	for _, p := range w.Players {
		timePacket := &packets.ClientboundTimePacket{
			Time: int64(6000),
		}
		p.SendPacket(timePacket)
	}
}
