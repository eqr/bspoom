package bsp

import (
	"bspoom/app/config"
	"bspoom/app/level"
	"fmt"
	"math"
)

type Builder struct {
	config.Config
	segments  []level.Segment //segments created during the bsp tree building
	segmentID int
}

func NewBuilder(cfg config.Config) *Builder {
	return &Builder{
		cfg,
		[]level.Segment{},
		0,
	}
}

func (b *Builder) Build(levelData level.LevelData) *Node {
	n := &Node{}
	b.buildTree(n, levelData.Segments)
	return n
}

func (b *Builder) buildTree(n *Node, segments []level.Segment) {
	if len(segments) == 0 {
		return
	}

	frontSegments, backSegments := b.splitSpace(n, segments)
	fmt.Println("frontSegments", frontSegments)
	fmt.Println("backSegments", backSegments)

	if len(backSegments) > 0 {
		n.Back = &Node{}
		b.buildTree(n.Back, backSegments)
	}

	if len(frontSegments) > 0 {
		n.Front = &Node{}
		b.buildTree(n.Front, frontSegments)
	}
}

func (b *Builder) splitSpace(node *Node, segments []level.Segment) ([]level.Segment, []level.Segment) {
	splitterSegment := segments[0]
	node.SplitterPoint1 = splitterSegment.P1
	node.SplitterPoint2 = splitterSegment.P2
	node.SplitterVector = splitterSegment.Vector()

	var frontSegments, backSegments []level.Segment
	inputSegments := segments[1:]
	for _, segment := range inputSegments {
		segmentVec := segment.Vector()
		numerator := segment.P1.Minus(node.SplitterPoint1).Cross2D(node.SplitterVector)
		denominator := node.SplitterVector.Cross2D(segmentVec)

		// if denominator is zero, the lines are parallel
		denominatorIsZero := math.Abs(float64(denominator)) < b.Config.FloatDelta

		// if numerator is zero, if they are parallel and nominator is zero
		numeratorIsZero := math.Abs(float64(numerator)) < b.Config.FloatDelta

		if denominatorIsZero && numeratorIsZero {
			frontSegments = append(frontSegments, segment)
			continue
		}

		if !denominatorIsZero {
			// intersection is the point where splitter divides the segment
			intersection := numerator / denominator

			// segments that are not parallel and t is in (0,1) should be split
			if intersection > 0.0 && intersection < 1.0 {
				intersectionPoint := segment.P1.Plus(segmentVec.Multiply(intersection))
				rightSegment := segment.Copy()
				rightSegment.P1 = segment.P1
				rightSegment.P2 = intersectionPoint

				leftSegment := segment.Copy()
				leftSegment.P1 = intersectionPoint
				leftSegment.P2 = segment.P2

				// swap sides if the beginning of the segment is on the back side of the splitter
				if numerator > 0 {
					rightSegment, leftSegment = leftSegment, rightSegment
				}

				frontSegments = append(frontSegments, rightSegment)
				backSegments = append(backSegments, leftSegment)
				continue
			}
		}

		// what if segment lies completely in front of behind the splitter
		if numerator < 0 || (numerator == 0 && denominator > 0) {
			frontSegments = append(frontSegments, segment)
		} else if numerator > 0 || (numerator == 0 && denominator < 0) {
			backSegments = append(backSegments, segment)
		}
	}

	b.AddSegment(splitterSegment, node)
	return frontSegments, backSegments
}

func (b *Builder) AddSegment(segment level.Segment, node *Node) {
	b.segments = append(b.segments, segment)
	node.SegmentID = b.segmentID
	b.segmentID++
}

func (b *Builder) GetSegments() []level.Segment {
	return b.segments
}
