package util

import (
	"encoding/binary"
	"math"
)

type PacketReader struct {
	index uint
	Data  *[]byte
}

func (r *PacketReader) ReadByte() byte {
	data := (*r.Data)[r.index]
	r.index++
	return data
}

func (r *PacketReader) ReadPacketId() byte {
	return r.ReadByte()
}

func (r *PacketReader) ReadBool() bool {
	// 0x00 = False; 0x01 = True
	return r.ReadByte() == 0x01
}

func (r *PacketReader) ReadShort() uint16 {
	data := make([]byte, SHORT_SIZE)

	shortIndex := 0
	endIndex := r.index + SHORT_SIZE
	for i := r.index; i < endIndex; i++ {
		data[shortIndex] = (*r.Data)[i]
		r.index++
		shortIndex++
	}

	return binary.BigEndian.Uint16(data)
}

func (r *PacketReader) ReadInt() int {
	data := make([]byte, INT_SIZE)

	intIndex := 0
	endIndex := r.index + INT_SIZE
	for i := r.index; i < endIndex; i++ {
		data[intIndex] = (*r.Data)[i]
		r.index++
		intIndex++
	}

	return int(int32(binary.BigEndian.Uint32(data)))
}

func (r *PacketReader) ReadLong() int64 {
	data := make([]byte, LONG_SIZE)

	intIndex := 0
	endIndex := r.index + LONG_SIZE
	for i := r.index; i < endIndex; i++ {
		data[intIndex] = (*r.Data)[i]
		r.index++
		intIndex++
	}

	return int64(binary.BigEndian.Uint64(data))
}

func (r *PacketReader) ReadFloat32() float32 {
	data := make([]byte, FLOAT_SIZE)

	intIndex := 0
	endIndex := r.index + FLOAT_SIZE
	for i := r.index; i < endIndex; i++ {
		data[intIndex] = (*r.Data)[i]
		r.index++
		intIndex++
	}

	bits := binary.BigEndian.Uint32(data)
	float := math.Float32frombits(bits)

	return float
}

func (r *PacketReader) ReadFloat64() float64 {
	data := make([]byte, DOUBLE_SIZE)

	intIndex := 0
	endIndex := r.index + DOUBLE_SIZE
	for i := r.index; i < endIndex; i++ {
		data[intIndex] = (*r.Data)[i]
		r.index++
		intIndex++
	}

	bits := binary.BigEndian.Uint64(data)
	float := math.Float64frombits(bits)

	return float
}

func (r *PacketReader) ReadString16() string {
	// read the length of the string in characters
	strLength := r.ReadShort()

	// convert string length to bytes
	strLengthBytes := uint(strLength * STRING_CHARACTER_SIZE)

	// get string bytes from data
	strData := (*r.Data)[r.index : r.index+strLengthBytes]

	// increment index by string length
	r.index += strLengthBytes

	// convert strData to string
	return string(strData)
}
