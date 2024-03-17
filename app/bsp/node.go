package bsp

import "bspoom/app/level"

type Node struct {
	Front          *Node
	Back           *Node
	SplitterPoint1 level.Point
	SplitterPoint2 level.Point
	SplitterVector level.Point
	SegmentID      int
}
