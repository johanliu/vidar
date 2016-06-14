package config

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/naoina/toml"
)

type tomlConfig struct {
	Title string
	Log   struct {
		Level string
		Path  string
	}
}

const configFile string = "../default.toml"

var config tomlConfig

func init() {
	f, err := os.Open(configFile)
	if err != nil {
		log.Panicf("Failed to open config file: %s", err)
	}

	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		log.Panicln("Failed to read config file: %s", err)
	}

	if err := toml.Unmarshal(buf, &config); err != nil {
		log.Panicf("Failed to unmarshal config file: %s", err)
	}
}
