package packets

import (
	"log"

	"github.com/cosrnic/b173-server/pkg/inventory"
	"github.com/cosrnic/b173-server/pkg/level"
	"github.com/cosrnic/b173-server/pkg/util"
)

type ClientboundPacket interface {
	Serialise() []byte
}

// keepalive
type ClientboundKeepAlivePacket struct {
}

func (p *ClientboundKeepAlivePacket) Serialise() []byte {
	w := util.NewPacketWriter()
	w.WriteByte(KeepAlive)
	return w.Bytes()
}

// login
type ClientboundLoginPacket struct {
	EntityID  int
	MapSeed   int64
	Dimension byte
}

func (p *ClientboundLoginPacket) Serialise() []byte {
	w := util.NewPacketWriter()
	w.WriteByte(LoginRequest)
	w.WriteInt32(int32(p.EntityID))
	w.WriteString16("") // unused
	w.WriteInt64(p.MapSeed)
	w.WriteByte(p.Dimension)

	return w.Bytes()
}

// handshake
type ClientboundHandshakePacket struct {
	ConnectionHash string
}

func (p *ClientboundHandshakePacket) Serialise() []byte {
	w := util.NewPacketWriter()
	w.WriteByte(Handshake)
	w.WriteString16(p.ConnectionHash)

	return w.Bytes()
}

// chat message
type ClientboundChatMessagePacket struct {
	Message string
}

func (p *ClientboundChatMessagePacket) Serialise() []byte {
	w := util.NewPacketWriter()
	w.WriteByte(ChatMessage)
	w.WriteString16(p.Message)

	return w.Bytes()
}

// time
type ClientboundTimePacket struct {
	Time int64
}

func (p *ClientboundTimePacket) Serialise() []byte {
	w := util.NewPacketWriter()
	w.WriteByte(Time)
	w.WriteInt64(p.Time)

	return w.Bytes()
}

// entity equipment

// spawn position
type ClientboundSpawnPositionPacket struct {
	X int32
	Y int32
	Z int32
}

func (p *ClientboundSpawnPositionPacket) Serialise() []byte {
	w := util.NewPacketWriter()
	w.WriteByte(SpawnPosition)
	w.WriteInt32(p.X)
	w.WriteInt32(p.Y)
	w.WriteInt32(p.Z)

	return w.Bytes()
}

// health

// respawn

// player position and look
type ClientboundPlayerPositionLook struct {
	X        float64
	Y        float64
	Z        float64
	Yaw      float32
	Pitch    float32
	OnGround bool
	Stance   float64
}

func (p *ClientboundPlayerPositionLook) Serialise() []byte {
	w := util.NewPacketWriter()

	w.WriteByte(PlayerPositionLook)
	w.WriteFloat64(p.X)
	w.WriteFloat64(p.Y)
	w.WriteFloat64(p.Stance)
	w.WriteFloat64(p.Z)
	w.WriteFloat32(p.Yaw)
	w.WriteFloat32(p.Pitch)
	w.WriteBool(p.OnGround)

	return w.Bytes()
}

// mine

// place

// held item

// click bed

// animation

// spawn player

// spawn item

// collect item

// spawn object

// spawn mob

// spawn painting

// stance [unused]

// entity velocity

// despawn entity

// entity

// entity relative position

// entity look

// entity relative position and look

// entity teleport

// entity status

// mount entity

// prechunk
type ClientboundPreChunkPacket struct {
	X    int32
	Z    int32
	Load bool
}

func (p *ClientboundPreChunkPacket) Serialise() []byte {
	w := util.NewPacketWriter()
	w.WriteByte(PreChunk)
	w.WriteInt32(p.X)
	w.WriteInt32(p.Z)
	w.WriteBool(p.Load)

	return w.Bytes()
}

// chunk
type ClientboundChunkPacket struct {
	Point          util.Point
	SizeX          byte
	SizeY          byte
	SizeZ          byte
	CompressedSize int32
	CompresedData  []byte
}

func (p *ClientboundChunkPacket) Apply(chunk level.Chunk) {
	p.Point = chunk.Point
	p.SizeX = chunk.SizeX
	p.SizeY = chunk.SizeY
	p.SizeZ = chunk.SizeZ
	p.CompresedData = chunk.Compress()
	p.CompressedSize = int32(len(p.CompresedData))
}

func (p *ClientboundChunkPacket) Serialise() []byte {
	w := util.NewPacketWriter()

	log.Println(p.Point.X)

	w.WriteByte(Chunk)
	w.WriteInt32(p.Point.X)
	w.WriteInt16(p.Point.Y)
	w.WriteInt32(p.Point.Z)
	w.WriteByte(p.SizeX)
	w.WriteByte(p.SizeY)
	w.WriteByte(p.SizeZ)
	w.WriteInt32(p.CompressedSize)
	w.Write(p.CompresedData)

	return w.Bytes()
}

// multi block change

// block change
type ClientboundBlockChangePacket struct {
	X     int32
	Y     byte
	Z     int32
	Block level.Block
}

func (p *ClientboundBlockChangePacket) Serialise() []byte {
	w := util.NewPacketWriter()

	w.WriteByte(BlockChange)
	w.WriteInt32(p.X)
	w.WriteByte(p.Y)
	w.WriteInt32(p.Z)
	w.WriteByte(p.Block.TypeId)
	w.WriteByte(p.Block.Metadata)

	return w.Bytes()
}

// block action

// explosion

// sound effect

// game state

// lightning bolt

// open window

// close window

// set slot
type ClientboundSetSlotPacket struct {
	WindowId byte
	Slot     int16
	Item     inventory.Item
}

func (p *ClientboundSetSlotPacket) Serialise() []byte {
	w := util.NewPacketWriter()

	w.WriteByte(SetSlot)
	w.WriteByte(p.WindowId)
	w.WriteShort(uint16(p.Slot))
	w.Write(p.Item.Serialize())

	return w.Bytes()
}

// window items
type ClientboundWindowItemsPacket struct {
	WindowId byte
	Count    int16
	Payload  inventory.Inventory
}

func (p *ClientboundWindowItemsPacket) Serialise() []byte {
	w := util.NewPacketWriter()

	w.WriteByte(WindowItems)
	w.WriteByte(p.WindowId)
	w.WriteInt16(p.Count)
	w.Write(p.Payload.Serialise())

	return w.Bytes()
}

// update progress bar

// transaction

// update sign

// item data

// statistic

// disconnect
