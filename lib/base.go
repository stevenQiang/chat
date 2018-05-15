package lib

import (
	"os"
)

func Getenv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0{
		return defaultValue
	}
	return value
}