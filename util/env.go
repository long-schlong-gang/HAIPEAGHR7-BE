/*
 *	Load environment variables
 */

package util

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/kpango/glg"
)

var GlobalEnvs = struct {
	LogFile string `env:"HPR_LOG_FILE" env-default:"haipeaghr7.log"`
	Port    int    `env:"HPR_PORT"`
}{}

func LoadConfig() {
	err := cleanenv.ReadConfig(EnvFile, &GlobalEnvs)
	if err != nil {
		glg.Errorf("Failed to parse '.env' file: %v", err)
		os.Exit(1)
	}
	glg.Info("Loaded ENV variables")
}
