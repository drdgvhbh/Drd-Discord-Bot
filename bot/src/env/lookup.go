package env

import (
	"fmt"
	"os"
)

func Lookup(key string) string {
	value, isSuccessful := os.LookupEnv(key)
	if !isSuccessful {
		panic(fmt.Sprintf("Environment variable \"%s\" not set", key))
	}

	return value
}
