package util

const (
	BOOL_SIZE   = 1
	BYTE_SIZE   = 1
	SHORT_SIZE  = 2
	INT_SIZE    = 4
	LONG_SIZE   = 8
	FLOAT_SIZE  = 4
	DOUBLE_SIZE = 8

	STRING_CHARACTER_SIZE = 2
)

type Packet struct {
	PacketId byte
}

func ReadPacket(data *[]byte) Packet {
	packet := Packet{}

	packet.PacketId, _ = readPacketId(data)

	return packet
}

func readPacketId(data *[]byte) (byte, uint) {
	return (*data)[0], 1
}
