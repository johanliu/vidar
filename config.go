package vidar

import (
	"io/ioutil"
	"os"

	"github.com/naoina/toml"
)

type tomlConfig struct {
	Version string
	Log     struct {
		Level string
		Path  string
	}
	Server struct {
		Host string
		Port string
	}
}

const configFile string = "default.toml"

var Config tomlConfig

func init() {
	f, err := os.Open(configFile)
	if err != nil {
		log.Error("Failed to open config file: %s", err)
	}

	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		log.Error("Failed to read config file: %s", err)
	}

	if err := toml.Unmarshal(buf, &Config); err != nil {
		log.Error("Failed to unmarshal config file: %s", err)
	}
}
