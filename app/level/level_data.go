package level

type LevelData struct {
	Segments []Segment
	Seed     int64
}

func NewLevelData(segments []Segment, seed int64) LevelData {
	return LevelData{
		Segments: segments,
		Seed:     seed,
	}
}
