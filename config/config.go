package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	// Bot API key, given by BotFather. Required.
	APIKey string = func() string {
		envVar := "TTOMBOT_TG_BOTAPI_KEY"

		s := os.Getenv(envVar)
		if s == "" {
			log.Fatalln(errMsg(envVar))
		}
		return s
	}()

	// Poll timeout in seconds. Optional. Default: 60 seconds.
	PollTimeoutSec int = func() int {
		envVar := "TTOMBOT_POLL_TIMEOUT_SEC"
		defaultSec := 60

		s := os.Getenv(envVar)
		if s == "" {
			log.Println(warnMsg(envVar))
			return defaultSec
		}

		nSec, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalln(errMsg(envVar))
		}

		return nSec
	}()
)

func errMsg(varName string) string {
	return fmt.Sprintf("ERROR: %s env var is missing or incorrect.", varName)
}

func warnMsg(varName string) string {
	return fmt.Sprintf("WARNING: %s env var is missing or incorrect.", varName)
}
