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

const configFile string = "/etc/vidar.conf"

var tc tomlConfig

func init() {
	f, err := os.Open(configFile)
	if err != nil {
		log.Error(err)
	}

	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		log.Error(err)
	}

	if err := toml.Unmarshal(buf, &tc); err != nil {
		log.Error(err)
	}
}
