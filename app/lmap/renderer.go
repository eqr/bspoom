package lmap

import (
	"bspoom/app/config"
	"bspoom/app/level"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

type MapRenderer interface {
	Draw()
}

type mapRenderer struct {
	cfg                    config.Config
	levelData              level.LevelData
	rawSegments            []level.Segment
	xMin, yMin, xMax, yMax float32
}

func NewRenderer(levelData level.LevelData, cfg config.Config) mapRenderer {
	mr := mapRenderer{
		levelData: levelData,
		cfg:       cfg,
	}
	mr.xMin, mr.yMin, mr.xMax, mr.yMax = getBounds(levelData.Segments)
	fmt.Println(mr.xMin, mr.yMin, mr.xMax, mr.yMax)

	mr.rawSegments = mr.RemapSegments(levelData.Segments)

	return mr
}

func (mr mapRenderer) Draw() {
	mr.DrawRawSegments()
}

func (mr mapRenderer) DrawRawSegments() {
	for _, segment := range mr.rawSegments {
		rl.DrawLineV(segment.P1.ToVector(), segment.P2.ToVector(), rl.Orange)
		xCenter := int32(segment.P1.X)
		yCenter := int32(segment.P1.Y)
		rl.DrawCircle(xCenter, yCenter, 3.0, rl.White)
	}
}

// TODO isolate it to remaper
func (mr mapRenderer) RemapSegments(segments []level.Segment) []level.Segment {
	res := make([]level.Segment, len(segments))
	for ix, segment := range segments {
		res[ix] = mr.RemapSegment(segment)
	}
	return res
}

func (mr mapRenderer) RemapSegment(s level.Segment) level.Segment {
	p1 := mr.RemapPoint(s.P1)
	p2 := mr.RemapPoint(s.P2)
	return level.Segment{P1: p1, P2: p2}
}

func (mr mapRenderer) RemapPoint(p level.Point) level.Point {
	x := mr.RemapX(p.X)
	y := mr.RemapY(p.Y)
	return level.Point{X: x, Y: y}
}

func (mr mapRenderer) RemapX(x float32) float32 {
	return (x-mr.xMin)*float32(mr.cfg.MapWidth-mr.cfg.MapOffset)/(mr.xMax-mr.xMin) + float32(mr.cfg.MapOffset)
}

func (mr mapRenderer) RemapY(y float32) float32 {
	return (y-mr.yMin)*float32(mr.cfg.MapHeight-mr.cfg.MapOffset)/(mr.yMax-mr.yMin) + float32(mr.cfg.MapOffset)
}

func getBounds(segments []level.Segment) (float32, float32, float32, float32) {
	xMin, yMin := float32(math.MaxFloat32), float32(math.MaxFloat32)
	xMax, yMax := float32(-math.MaxFloat32), float32(-math.MaxFloat32)
	for _, segment := range segments {
		if segment.P1.X < xMin {
			xMin = segment.P1.X
		}

		if segment.P2.X < xMin {
			xMin = segment.P2.X
		}

		if segment.P1.X > xMax {
			xMax = segment.P1.X
		}

		if segment.P2.X > xMax {
			xMax = segment.P2.X
		}

		if segment.P1.Y < yMin {
			yMin = segment.P1.Y
		}

		if segment.P1.Y > yMax {
			yMax = segment.P1.Y
		}

		if segment.P2.Y < yMin {
			yMin = segment.P2.Y
		}

		if segment.P2.Y > yMax {
			yMax = segment.P2.Y
		}
	}

	return xMin, yMin, xMax, yMax
}
