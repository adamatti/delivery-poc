package config

var Instance = New()

type Config struct {
	ServicePort int
	MongoUrl    string
}

func New() Config {
	// FIXME get it from env or config file
	c:= Config{
		ServicePort: 80,
		MongoUrl: "mongodb://mongo:mongo@localhost:27017/products?authSource=admin",
	}
	return c
}