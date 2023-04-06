package env

import (
	"os"

	"github.com/chulipinho/person-api/config"
)

func Parse(key string) string {
	v := os.Getenv(key)
	if v == "" {
		return config.Get(key)
	}

	return v
}
