package packets

import (
	"log"

	"github.com/cosrnic/b173-server/pkg/util"
)

// keepalive
type ServerboundKeepAlivePacket struct {
	util.Packet
}

func ReadServerboundKeepAlivePacket(data *[]byte) ServerboundKeepAlivePacket {
	reader := util.PacketReader{
		Data: data,
	}

	packet := ServerboundKeepAlivePacket{}
	packet.PacketId = reader.ReadPacketId()

	return packet
}

// login
type ServerboundLoginPacket struct {
	util.Packet
	ProtocolVersion int
	Username        string
	MapSeed         int64
	Dimension       byte
}

func ReadServerboundLoginPacket(data *[]byte) ServerboundLoginPacket {
	reader := util.PacketReader{
		Data: data,
	}

	packet := ServerboundLoginPacket{}
	packet.PacketId = reader.ReadPacketId()
	packet.ProtocolVersion = reader.ReadInt()
	packet.Username = reader.ReadString16()
	packet.MapSeed = reader.ReadLong()
	packet.Dimension = reader.ReadByte()

	return packet

}

// handshake
type ServerboundHandshakePacket struct {
	util.Packet
	Username string
}

func ReadServerboundHandshakePacket(data *[]byte) ServerboundHandshakePacket {
	reader := util.PacketReader{
		Data: data,
	}

	packet := ServerboundHandshakePacket{}
	packet.PacketId = reader.ReadPacketId()
	packet.Username = reader.ReadString16()

	return packet
}

// chat message
type ServerboundChatMessagePacket struct {
	util.Packet
	Message string
}

func ReadServerboundChatMessagePacket(data *[]byte) ServerboundChatMessagePacket {
	reader := util.PacketReader{
		Data: data,
	}

	packet := ServerboundChatMessagePacket{}
	packet.PacketId = reader.ReadPacketId()
	packet.Message = reader.ReadString16()

	return packet
}

// entity equipment

// interact entity

// respawn

// player

// player position

// player look

// player position and look
type ServerboundPlayerPositionLook struct {
	util.Packet
	X        float64
	Y        float64
	Z        float64
	Yaw      float32
	Pitch    float32
	Stance   float64
	OnGround bool
}

func ReadServerboundPlayerPositionLookPacket(data *[]byte) ServerboundPlayerPositionLook {
	reader := util.PacketReader{
		Data: data,
	}

	packet := ServerboundPlayerPositionLook{}
	packet.PacketId = reader.ReadPacketId()
	packet.X = reader.ReadFloat64()
	packet.Y = reader.ReadFloat64()
	packet.Stance = reader.ReadFloat64()
	packet.Z = reader.ReadFloat64()
	packet.Yaw = reader.ReadFloat32()
	packet.Pitch = reader.ReadFloat32()
	packet.OnGround = reader.ReadBool()

	return packet
}

// mine
type ServerboundPlayerMine struct {
	util.Packet
	Status byte
	X      int
	Y      byte
	Z      int
	Face   byte
}

func ReadServerboundPlayerMine(data *[]byte) ServerboundPlayerMine {
	reader := util.PacketReader{
		Data: data,
	}

	packet := ServerboundPlayerMine{}
	packet.PacketId = reader.ReadPacketId()
	packet.Status = reader.ReadByte()
	packet.X = reader.ReadInt()
	log.Print(packet.X)
	packet.Y = reader.ReadByte()
	packet.Z = reader.ReadInt()
	packet.Face = reader.ReadByte()

	return packet
}

// place

// held item

// animation

// spawn item

// spawn painting

// stance [unused]

// entity velocity

// mount entity

// multi block change

// block change

// explosion

// sound effect

// game state

// open window

// close window

// window click

// transaction

// update sign

// statistic

// disconnect
