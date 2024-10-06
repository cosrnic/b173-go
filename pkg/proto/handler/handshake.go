package handler

import (
	"log"

	"github.com/cosrnic/b173-server/pkg/proto/packets"
	"github.com/cosrnic/b173-server/pkg/world"
)

func HandleServerboundHandshakePacket(player world.Player, packet packets.ServerboundHandshakePacket) {
	log.Printf("Handshake: %+v", packet)

	clientboundPacket := packets.ClientboundHandshakePacket{
		ConnectionHash: "-",
	}
	outData := clientboundPacket.Serialise()

	player.GetConnection().Write(outData)
}
