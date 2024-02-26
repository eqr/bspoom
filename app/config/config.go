package config

const (
	mapOffset = 50
)

type Config struct {
	WinWidth  int
	WinHeight int
	MapWidth  int
	MapHeight int
	MapOffset int
}

func NewConfig() Config {
	c := Config{
		WinWidth:  1600,
		WinHeight: 900,
		MapOffset: mapOffset,
	}

	c.MapWidth = c.WinWidth - mapOffset
	c.MapHeight = c.WinHeight - mapOffset
	return c
}
