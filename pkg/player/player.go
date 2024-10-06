package player

import (
	"errors"
	"io"
	"log"
	"net"

	"github.com/cosrnic/b173-server/pkg/inventory"
	"github.com/cosrnic/b173-server/pkg/level"
	"github.com/cosrnic/b173-server/pkg/proto/handler"
	"github.com/cosrnic/b173-server/pkg/proto/packets"
	"github.com/cosrnic/b173-server/pkg/util"
	"github.com/cosrnic/b173-server/pkg/world"
)

const (
	RENDER_DISTANCE = 32
)

type Player struct {
	Username   string
	Inventory  inventory.Inventory
	World      *world.World
	Connection net.Conn
	Position   util.Pos
}

func (p *Player) GetUsername() string {
	return p.Username
}

func (p *Player) SetUsername(username string) {
	p.Username = username
}

func (p *Player) GetConnection() net.Conn {
	return p.Connection
}

func (p *Player) GetWorld() *world.World {
	return p.World
}

func (p *Player) GetPosition() util.Pos {
	return p.Position
}

func (p *Player) SetPosition(pos util.Pos) {
	p.Position = pos
}

func (p *Player) SendPacket(packet packets.ClientboundPacket) {
	p.Connection.Write(packet.Serialise())
}

func NewPlayer(world *world.World, conn net.Conn) Player {
	return Player{
		Username:   "",
		World:      world,
		Connection: conn,
		Inventory:  inventory.NewInventory(inventory.PLAYER_INVENTORY_SIZE),
	}
}

func (p *Player) ReadLoop() {
	var buf []byte
	for {
		buf = make([]byte, 1024)

		_, err := p.Connection.Read(buf)
		if errors.Is(err, io.EOF) || errors.Is(err, net.ErrClosed) {
			return
		} else if err != nil {
			panic(err)
		}

		p.processPackets(&buf)
	}
}

func (p *Player) processPackets(data *[]byte) {
	pp := util.ReadPacket(data)

	switch pp.PacketId {
	case packets.Handshake:
		packet := packets.ReadServerboundHandshakePacket(data)
		handler.HandleServerboundHandshakePacket(p, packet)
	case packets.LoginRequest:
		packet := packets.ReadServerboundLoginPacket(data)
		handler.HandleServerboundLoginPacket(p, packet)
	case packets.ChatMessage:
		serverboundPacket := packets.ReadServerboundChatMessagePacket(data)

		newMessage := "<" + p.GetUsername() + "> " + serverboundPacket.Message

		clientboundPacket := &packets.ClientboundChatMessagePacket{
			Message: newMessage,
		}
		p.SendPacket(clientboundPacket)
	case packets.PlayerPositionLook:
		packet := packets.ReadServerboundPlayerPositionLookPacket(data)
		p.SetPosition(util.Pos{
			X:     packet.X,
			Y:     packet.Y,
			Z:     packet.Z,
			Yaw:   packet.Yaw,
			Pitch: packet.Pitch,
		})
	case packets.Mine:
		packet := packets.ReadServerboundPlayerMine(data)
		for _, chunk := range *p.World.Chunks {
			// log.Printf("chunk x %d, packet x %d", chunk.Point.ChunkX, packet.X/16)
			// log.Printf("chunk z %d, packet z %d", chunk.Point.ChunkZ, packet.Z/16)
			if chunk.Point.ChunkX == int32(packet.X/16) {

				if chunk.Point.ChunkZ == int32(packet.Z/16) {
					chunk.SetBlock(util.NewPoint(int32(packet.X), int16(packet.Y), int32(packet.Z)), level.NewAirBlock())
					log.Printf("breaking block at %d, %d, %d", packet.X, packet.Y, packet.Z)
					chunkpacket := &packets.ClientboundChunkPacket{}
					chunkpacket.Apply(chunk)

					p.SendPacket(chunkpacket)
				}
			}
		}
	default:
		// log.Printf("Unprocessed packet: %+v", pp)
	}
}
