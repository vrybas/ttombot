package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	// Bot API key, given by BotFather. Required.
	APIKey string = func(name string) string {
		s := os.Getenv(name)
		if s == "" {
			log.Fatalln(errMsg(name))
		}
		return s
	}("TTOMBOT_TG_BOTAPI_KEY")

	// Poll timeout in seconds. Optional. Default: 60 seconds.
	PollTimeoutSec int = func(name string) int {
		defaultSec := 60

		s := os.Getenv(name)
		if s == "" {
			log.Println(warnMsg(name))
			return defaultSec
		}

		nSec, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalln(errMsg(name))
		}

		return nSec
	}("TTOMBOT_POLL_TIMEOUT_SEC")
)

func errMsg(varName string) string {
	return fmt.Sprintf("ERROR: %s env var is missing or incorrect.", varName)
}

func warnMsg(varName string) string {
	return fmt.Sprintf("WARNING: %s env var is missing or incorrect.", varName)
}
