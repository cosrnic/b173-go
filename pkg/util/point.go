package util

type Pos struct {
	X     float64
	Y     float64
	Z     float64
	Yaw   float32
	Pitch float32
}

type Point struct {
	X      int32
	Y      int16
	Z      int32
	ChunkX int32
	ChunkZ int32
}

func NewPoint(x int32, y int16, z int32) Point {
	return Point{
		X:      x,
		Y:      y,
		Z:      z,
		ChunkX: x / 16,
		ChunkZ: z / 16,
	}
}
