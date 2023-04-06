package config

var config map[string]string = map[string]string{
	"PORT":    ":1234",
	"USEMOCK": "true",
}

func Get(key string) string {
	return config[key]
}
