package config

var Instance = New()

type Config struct {
	ServicePort         int
}

func New() Config {
	// FIXME get it from env or config file
	c:= Config{ServicePort: 80}
	return c
}