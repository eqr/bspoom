package config

type Config struct {
	WinWidth  int
	WinHeight int
}

func NewConfig() Config {
	return Config{
		WinWidth:  1600,
		WinHeight: 900,
	}
}
