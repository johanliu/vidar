package vidar

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
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
	// Initialize config with defaults
	tc = tomlConfig{
		Version: "1.0",
	}
	tc.Log.Level = "INFO"
	tc.Server.Host = DefaultHost
	tc.Server.Port = DefaultPort

	// Try to load config file if it exists
	f, err := os.Open(configFile)
	if err != nil {
		// Config file doesn't exist, use defaults
		return
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		// Can't read config file, use defaults
		return
	}

	// Try to parse TOML config
	if err := toml.Unmarshal(buf, &tc); err != nil {
		// Invalid TOML format, use defaults
		return
	}
}
