package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var (
	TG_BOTAPI_KEY string
)

func init() {
	viper.SetEnvPrefix("TTOMBOT")

	viper.BindEnv("TG_BOTAPI_KEY")
	if TG_BOTAPI_KEY = viper.GetString("TG_BOTAPI_KEY"); TG_BOTAPI_KEY == "" {
		log.Fatalln(errorMsg("TG_BOTAPI_KEY is missing"))
	}
}

func errorMsg(in string) string {
	return fmt.Sprintf("Configuration error: %s\n", in)
}
