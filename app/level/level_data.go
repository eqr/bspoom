package level

type LevelData struct {
	Segments []Segment
}

func NewLevelData(segments []Segment) LevelData {
	return LevelData{
		Segments: segments,
	}
}
