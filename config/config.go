package config

import (
	"fmt"
	"log"
	"os"
)

var (
	TG_BOTAPI_KEY string = readEnv("TTOMBOT_TG_BOTAPI_KEY")
)

// readEnv tries to read ENV variable, or fails if variable is not set or has
// zero length.
func readEnv(name string) (out string) {
	if out = os.Getenv(name); out == "" {
		log.Fatalln(fmt.Sprintf("Configuration error: %s env var is missing", name))
	}

	return
}
