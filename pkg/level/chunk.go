package level

import (
	"bytes"
	"compress/zlib"
	"log"

	"github.com/cosrnic/b173-server/pkg/util"
)

const (
	CHUNK_SIZE_X = 16
	CHUNK_SIZE_Y = 128
	CHUNK_SIZE_Z = 16
)

type Chunk struct {
	Point util.Point
	SizeX byte
	SizeY byte
	SizeZ byte
	Data  []byte

	BlockTypes    []byte
	BlockMetadata []byte
	BlockLight    []byte
	SkyLight      []byte
	HeightMap     []byte
}

func (c *Chunk) Compress() []byte {
	var buf bytes.Buffer

	buf.Write(c.BlockTypes)
	buf.Write(c.BlockMetadata)
	buf.Write(c.BlockLight)
	buf.Write(c.SkyLight)

	var compressed bytes.Buffer
	w := zlib.NewWriter(&compressed)
	w.Write(buf.Bytes())
	w.Close()

	return compressed.Bytes()
}

func (c *Chunk) FillBlock(from util.Point, to util.Point, block Block) {
	for x := from.X; x < from.X+to.X; x++ {
		for z := from.Z; z < from.Z+to.Z; z++ {
			for y := from.Y; y < from.Y+to.Y; y++ {
				index := calc3DIndex(int(x), int(y), int(z))
				c.BlockTypes[index] = block.TypeId
				c.setMetadata(index, block.Metadata)
			}
		}
	}
}

func (c *Chunk) FillLight(from util.Point, to util.Point, blockLight, skyLight byte) {
	for x := from.X; x < from.X+to.X; x++ {
		for z := from.Z; z < from.Z+to.Z; z++ {
			for y := from.Y; y < from.Y+to.Y; y++ {
				index := calc3DIndex(int(x), int(y), int(z))
				c.setBlockLight(index, blockLight)
				c.setSkyLight(index, skyLight)
			}
		}
	}
}

func (c *Chunk) SetBlock(pos util.Point, block Block) {
	index := calc3DIndex(int(pos.X), int(pos.Y), int(pos.Z))
	c.BlockTypes[index] = block.TypeId
	c.setMetadata(index, block.Metadata)
}

func (c *Chunk) GetBlock(pos util.Point) (byte, byte) {
	index := calc3DIndex(int(pos.X), int(pos.Y), int(pos.Z))
	return c.BlockTypes[index], c.getMetadata(index)
}

func calc3DIndex(x, y, z int) int {
	return (x&0xF)<<11 | (z&0xF)<<7 | (y & 0x7F)
}

func (c *Chunk) getMetadata(index int) byte {
	meta := c.BlockMetadata[index/2]
	if index%2 == 0 {
		return meta & 0x0F
	}
	return (meta >> 4) & 0x0F
}

func (c *Chunk) setMetadata(index int, value byte) {
	if value > 0x0F {
		log.Panicf("Metadata value %d out of range", value)
	}
	if index%2 == 0 {
		c.BlockMetadata[index/2] = (c.BlockMetadata[index/2] & 0xF0) | value
	} else {
		c.BlockMetadata[index/2] = (c.BlockMetadata[index/2] & 0x0F) | (value << 4)
	}
}

func (c *Chunk) getBlockLight(index int) byte {
	light := c.BlockLight[index/2]
	if index%2 == 0 {
		return light & 0x0F
	}
	return (light >> 4) & 0x0F
}

func (c *Chunk) setBlockLight(index int, value byte) {
	if value > 0x0F {
		log.Panicf("Light value %d out of range", value)
	}
	if index%2 == 0 {
		c.BlockLight[index/2] = (c.BlockLight[index/2] & 0xF0) | value
	} else {
		c.BlockLight[index/2] = (c.BlockLight[index/2] & 0x0F) | (value << 4)
	}
}

func (c *Chunk) getSkyLight(index int) byte {
	light := c.SkyLight[index/2]
	if index%2 == 0 {
		return light & 0x0F
	}
	return (light >> 4) & 0x0F
}

func (c *Chunk) setSkyLight(index int, value byte) {
	if value > 0x0F {
		log.Panicf("Sky light value %d out of range", value)
	}
	if index%2 == 0 {
		c.SkyLight[index/2] = (c.SkyLight[index/2] & 0xF0) | value
	} else {
		c.SkyLight[index/2] = (c.SkyLight[index/2] & 0x0F) | (value << 4)
	}
}

func NewChunk(pos util.Point) *Chunk {
	blocksAmount := CHUNK_SIZE_X * CHUNK_SIZE_Y * CHUNK_SIZE_Z
	return &Chunk{
		Point:         pos,
		SizeX:         byte(CHUNK_SIZE_X - 1),
		SizeY:         byte(CHUNK_SIZE_Y - 1),
		SizeZ:         byte(CHUNK_SIZE_Z - 1),
		BlockTypes:    make([]byte, blocksAmount),
		BlockMetadata: make([]byte, divEvenOrRoundUp(blocksAmount, 2)),
		BlockLight:    make([]byte, divEvenOrRoundUp(blocksAmount, 2)),
		SkyLight:      make([]byte, divEvenOrRoundUp(blocksAmount, 2)),
		HeightMap:     make([]byte, CHUNK_SIZE_X*CHUNK_SIZE_Z),
	}
}
