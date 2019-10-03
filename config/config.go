package config

import (
	"fmt"
	"log"
	"os"
)

var (
	TG_BOTAPI_KEY string = readEnv("TTOMBOT_TG_BOTAPI_KEY")
)

func readEnv(name string) (out string) {
	if out = os.Getenv(name); out == "" {
		log.Fatalln(fmt.Sprintf("Configuration error: %s env var is missing", name))
	}

	return
}
