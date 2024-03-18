package bsp

import (
	"bspoom/app/config"
	"bspoom/app/level"
	"bspoom/app/utils/slice"
	"fmt"
	"math"
)

type Builder struct {
	config.Config
	segments                     []level.Segment //segments created during the bsp tree building
	segmentID                    int
	numFront, numBack, numSplits int
}

func NewBuilder(cfg config.Config) *Builder {
	return &Builder{
		Config:   cfg,
		segments: []level.Segment{},
	}
}

func (b *Builder) Build(levelData level.LevelData) *Node {
	var seed int64
	if levelData.Seed == 0 {
		seed = b.getBestSeed(0, 50000, 3, levelData.Segments)
	} else {
		seed = levelData.Seed
	}

	segments := slice.Shuffle(levelData.Segments, seed)
	n := &Node{}
	b.buildTree(n, segments)
	fmt.Println("numFront", b.numFront)
	fmt.Println("numBack", b.numBack)
	fmt.Println("numSplits", b.numSplits)
	return n
}

func (b *Builder) buildTree(n *Node, segments []level.Segment) {
	if len(segments) == 0 {
		return
	}

	frontSegments, backSegments := b.splitSpace(n, segments)

	if len(backSegments) > 0 {
		b.numBack++
		n.Back = &Node{}
		b.buildTree(n.Back, backSegments)
	}

	if len(frontSegments) > 0 {
		b.numFront++
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
				b.numSplits++
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

func (b *Builder) getBestSeed(startSeed, endSeed, weightFactor int64, levelSegments []level.Segment) int64 {
	bestSeed := int64(-1)
	bestScore := math.MaxFloat32
	var seed int64
	for seed = startSeed; seed < endSeed; seed++ {
		segments := make([]level.Segment, len(levelSegments))
		copy(segments, levelSegments)
		segments = slice.Shuffle(segments, seed)
		fmt.Println(segments)
		candidateBuilder := NewBuilder(b.Config)
		rootNode := &Node{}
		candidateBuilder.buildTree(rootNode, segments)
		score := math.Abs(float64(candidateBuilder.numBack-candidateBuilder.numFront)) + float64(weightFactor*int64(candidateBuilder.numSplits))
		fmt.Println(candidateBuilder.numBack, candidateBuilder.numFront, candidateBuilder.numSplits, score)
		if score < bestScore {
			bestScore = score
			bestSeed = seed
		}
	}

	fmt.Println("bestSeed", bestSeed, "bestScore", bestScore)
	return bestSeed
}
