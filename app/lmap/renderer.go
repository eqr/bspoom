package lmap

import (
	"bspoom/app/bsp"
	"bspoom/app/config"
	"bspoom/app/level"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
	"math"
)

const scale = 12.0

type MapRenderer interface {
	Draw()
}

type mapRenderer struct {
	cfg                    config.Config
	levelData              level.LevelData
	rawSegments            []level.Segment
	xMin, yMin, xMax, yMax float32
	segments               []level.Segment
	counter                float32
	traverser              *bsp.Traverser
}

func NewRenderer(levelData level.LevelData, builder *bsp.Builder, traverser *bsp.Traverser, cfg config.Config) *mapRenderer {
	mr := mapRenderer{
		levelData: levelData,
		counter:   0,
		traverser: traverser,
		cfg:       cfg,
	}
	mr.xMin, mr.yMin, mr.xMax, mr.yMax = getBounds(levelData.Segments)
	fmt.Println(mr.xMin, mr.yMin, mr.xMax, mr.yMax)

	mr.rawSegments = mr.RemapSegments(levelData.Segments)
	mr.segments = mr.RemapSegments(builder.GetSegments())
	return &mr
}

func (mr *mapRenderer) Draw() {
	mr.DrawRawSegments()
	mr.DrawSegments(rl.Orange)
	mr.DrawPlayer()
	mr.counter += 0.0005
}

func (mr *mapRenderer) DrawPlayer() {
	p := mr.RemapPoint(mr.traverser.GetCameraPosition())
	rl.DrawCircleV(p.ToVector(), 10, rl.Green)
}

func (mr *mapRenderer) DrawSegments(color color.RGBA) {
	upperBound := int(mr.counter) % (len(mr.traverser.GetSegmentIDsToDraw()) + 1)
	fmt.Println(mr.counter, len(mr.traverser.GetSegmentIDsToDraw())+1, upperBound)

	for _, segmentID := range mr.traverser.GetSegmentIDsToDraw()[:upperBound] {
		seg := mr.segments[segmentID]
		p1 := seg.P1
		p2 := seg.P2

		rl.DrawLineV(p1.ToVector(), p2.ToVector(), color)
		mr.DrawNormal(p1, p2, color, scale)
		rl.DrawCircleV(p1.ToVector(), 3, rl.White)
	}
}

// DrawRawSegments draws the raw segments with barely visible lines.
func (mr *mapRenderer) DrawRawSegments() {
	for _, segment := range mr.rawSegments {
		rl.DrawLineV(segment.P1.ToVector(), segment.P2.ToVector(), rl.DarkGray)
	}
}

func (mr *mapRenderer) DrawNormal(p0, p1 level.Point, color color.RGBA, scale float32) {
	// middle of the vector
	p10 := level.Point{X: p1.X - p0.X, Y: p1.Y - p0.Y}
	// rotate 90 degrees
	p10rotated := level.Point{X: -10 * p10.Y, Y: 10 * p10.X}

	normal := p10rotated.Normalize()

	// beginning of normal vector
	n0 := level.Point{X: (p1.X + p0.X) * 0.5, Y: (p1.Y + p0.Y) * 0.5}
	// end of normal vector
	n1 := level.Point{X: n0.X + normal.X*scale, Y: n0.Y + normal.Y*scale}
	rl.DrawLineV(n0.ToVector(), n1.ToVector(), color)
}

// TODO isolate it to remaper
func (mr *mapRenderer) RemapSegments(segments []level.Segment) []level.Segment {
	res := make([]level.Segment, len(segments))
	for ix, segment := range segments {
		res[ix] = mr.RemapSegment(segment)
	}
	return res
}

func (mr *mapRenderer) RemapSegment(s level.Segment) level.Segment {
	p1 := mr.RemapPoint(s.P1)
	p2 := mr.RemapPoint(s.P2)
	return level.Segment{P1: p1, P2: p2}
}

func (mr *mapRenderer) RemapPoint(p level.Point) level.Point {
	x := mr.RemapX(p.X)
	y := mr.RemapY(p.Y)
	return level.Point{X: x, Y: y}
}

func (mr *mapRenderer) RemapX(x float32) float32 {
	return (x-mr.xMin)*float32(mr.cfg.MapWidth-mr.cfg.MapOffset)/(mr.xMax-mr.xMin) + float32(mr.cfg.MapOffset)
}

func (mr *mapRenderer) RemapY(y float32) float32 {
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
