package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/cosrnic/b173-server"
	"github.com/cosrnic/b173-server/pkg/level"
	"github.com/cosrnic/b173-server/pkg/player"
	"github.com/cosrnic/b173-server/pkg/util"
	"github.com/cosrnic/b173-server/pkg/world"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// chunk1 := level.NewChunk(util.NewPoint(0, 0, 0))
	// chunk1.FillBlock(util.NewPoint(0, 0, 0), util.NewPoint(level.CHUNK_SIZE_X, 1, level.CHUNK_SIZE_Z), level.NewStoneBlock())
	// chunk1.FillLight(util.NewPoint(0, 0, 0), util.NewPoint(level.CHUNK_SIZE_X, level.CHUNK_SIZE_Y, level.CHUNK_SIZE_Z), 0, 15)

	// chunk2 := level.NewChunk(util.NewPoint(1*level.CHUNK_SIZE_X, 0, 0))
	// chunk2.FillBlock(util.NewPoint(0, 1, 0), util.NewPoint(level.CHUNK_SIZE_X, 1, level.CHUNK_SIZE_Z), level.NewGrassBlock())
	// chunk2.FillLight(util.NewPoint(0, 1, 0), util.NewPoint(level.CHUNK_SIZE_X, level.CHUNK_SIZE_Y, level.CHUNK_SIZE_Z), 0, 15)

	var chunks []level.Chunk
	for x := -player.RENDER_DISTANCE / 2; x < player.RENDER_DISTANCE/2; x++ {
		for z := -player.RENDER_DISTANCE / 2; z < player.RENDER_DISTANCE/2; z++ {
			chunk := level.NewChunk(util.NewPoint(int32(x)*level.CHUNK_SIZE_X, 0, int32(z)*level.CHUNK_SIZE_Z))
			chunk.FillBlock(util.NewPoint(0, 0, 0), util.NewPoint(level.CHUNK_SIZE_X, 1, level.CHUNK_SIZE_Z), level.NewGrassBlock())
			// chunk.FillBlock(util.NewPoint(0, 1, 0), util.NewPoint(level.CHUNK_SIZE_X, 1, level.CHUNK_SIZE_X), level.Block{
			// 	TypeId:   0x9,
			// 	Metadata: 0x00,
			// 	Light:    0x00,
			// 	SkyLight: 0x00,
			// })
			chunk.FillLight(util.NewPoint(0, 0, 0), util.NewPoint(level.CHUNK_SIZE_X, level.CHUNK_SIZE_Y, level.CHUNK_SIZE_Z), 0, 15)
			chunks = append(chunks, *chunk)
		}
	}

	s := b173.Server{
		World: &world.World{
			Chunks: &chunks,
		},
	}

	s.Start()

	<-ctx.Done()

	stopCtx, stopCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer stopCancel()

	_ = stopCtx
	// if err := s.Stop(); err != nil {
	// panic(err)
	// }

	println("Goodbye, World!")
}
