package handler

import (
	"log"

	"github.com/cosrnic/b173-server/pkg/inventory"
	"github.com/cosrnic/b173-server/pkg/proto/packets"
	"github.com/cosrnic/b173-server/pkg/world"
)

func HandleServerboundLoginPacket(player world.Player, packet packets.ServerboundLoginPacket) {
	log.Printf("Login: %+v", packet)

	player.SetUsername(packet.Username)

	sendLoginResponse(player)
	// sendInventory(c)
	sendSpawnPosition(player)
	sendPlayerPositionLook(player)

	sendPreChunk(player)
	sendChunk(player)

	player.GetWorld().Players = append(player.GetWorld().Players, player)

}

func sendLoginResponse(player world.Player) {
	clientboundPacket := &packets.ClientboundLoginPacket{
		EntityID:  0,
		MapSeed:   0,
		Dimension: 0,
	}

	player.SendPacket(clientboundPacket)
}

func sendPreChunk(player world.Player) {

	for _, chunk := range *player.GetWorld().Chunks {
		log.Printf("x: %+v, chunkX: %+v", chunk.Point.X, chunk.Point.ChunkX)

		packet := &packets.ClientboundPreChunkPacket{
			X:    chunk.Point.ChunkX,
			Z:    chunk.Point.ChunkZ,
			Load: true,
		}

		player.SendPacket(packet)
	}

}

func sendChunk(player world.Player) {

	for _, chunk := range *player.GetWorld().Chunks {

		packet := &packets.ClientboundChunkPacket{}
		packet.Apply(chunk)

		player.SendPacket(packet)
	}

}

func sendSpawnPosition(player world.Player) {

	clientboundSpawnPosition := &packets.ClientboundSpawnPositionPacket{
		X: 0,
		Y: 130,
		Z: 0,
	}

	player.SendPacket(clientboundSpawnPosition)
}

func sendInventory(player world.Player) {

	inv := inventory.NewInventory(inventory.PLAYER_INVENTORY_SIZE)

	clientboundWindowItems := &packets.ClientboundWindowItemsPacket{
		WindowId: 0,
		Count:    int16(inv.Size),
		Payload:  inv,
	}

	player.SendPacket(clientboundWindowItems)

	clientboundSetSlot := &packets.ClientboundSetSlotPacket{
		WindowId: 0x81,
		Slot:     -1,
		Item:     inventory.NewItem(-1, 1),
	}

	player.SendPacket(clientboundSetSlot)

}

func sendPlayerPositionLook(player world.Player) {

	packet := &packets.ClientboundPlayerPositionLook{
		X:        0,
		Y:        130,
		Z:        0,
		Yaw:      0,
		Pitch:    0,
		OnGround: false,
	}
	packet.Stance = packet.Y + 1.62

	player.SendPacket(packet)
}
