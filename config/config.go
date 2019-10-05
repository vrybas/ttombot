package config

import (
	"fmt"
	"log"
	"os"
)

var (
	APIKey string = func(name string) (v string) {
		if v = os.Getenv(name); v == "" {
			log.Fatalln(errMsg(name))
		}
		return
	}("TTOMBOT_TG_BOTAPI_KEY")
)

func errMsg(varName string) string {
	return fmt.Sprintf("ERROR: %s env var is missing.", varName)
}

func warnMsg(varName string) string {
	return fmt.Sprintf("WARNING: %s env var is missing.", varName)
}
