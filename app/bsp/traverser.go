package bsp

import (
	"bspoom/app/level"
)

type Traverser struct {
	node             *Node
	segments         []level.Segment
	cameraPosition   level.Point
	segmentIDsToDraw []int
}

func NewTraverser(tree *Node, segments []level.Segment) *Traverser {
	return &Traverser{
		node:             tree,
		segments:         segments,
		cameraPosition:   level.Point{6, 7},
		segmentIDsToDraw: make([]int, 0),
	}
}

func (t *Traverser) Update() {
	t.segmentIDsToDraw = make([]int, 0)
	t.traverse(t.node)
}

func (t *Traverser) traverse(node *Node) {
	if node == nil {
		return
	}

	v1 := t.cameraPosition.Minus(node.SplitterPoint1)
	onFront := v1.IsOnFront(node.SplitterVector)
	// traversing from front to back
	if onFront {
		t.traverse(node.Front)
		t.segmentIDsToDraw = append(t.segmentIDsToDraw, node.SegmentID)
		t.traverse(node.Back)
	} else {
		t.traverse(node.Back)
		t.segmentIDsToDraw = append(t.segmentIDsToDraw, node.SegmentID)
		t.traverse(node.Front)
	}
}

func (t *Traverser) GetCameraPosition() level.Point {
	return t.cameraPosition
}

func (t *Traverser) GetSegmentIDsToDraw() []int {
	return t.segmentIDsToDraw
}
