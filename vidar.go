package vidar

import (
	"net/http"
	"strings"

	"github.com/johanliu/Vidar/config"
	"github.com/johanliu/Vidar/logger"
	"github.com/johanliu/Vidar/router"
)

type Vidar struct {
	Route *router.Router
}

func New() *Vidar {
	return &Vidar{router.New()}
}

//TODO: Implement cgi and fast cgi interface
func (v *Vidar) Run(addr ...string) error {

	address, err := resolveAddress(addr...)
	if err != nil {
		logger.Error.Printf("Resolve address failed: %v\n", err)
	}

	logger.Info.Printf("Running on %s", address)

	if err := http.ListenAndServe(address, v.Route); err != nil {
		logger.Error.Printf("Server start failed: %v", err)
	}

	return nil
}

func resolveAddress(addr ...string) (string, error) {

	switch len(addr) {
	case 0:
		if host := config.Config.Server.Host; len(host) > 0 {
			logger.Debug.Printf("Read host value from config file: %s\n", host)

			if port := config.Config.Server.Port; len(port) > 0 {
				logger.Debug.Printf("Read port value from config file: %s\n", port)

				return host + ":" + port, nil
			}
		}
	case 2:
		return strings.Join(addr, ":"), nil
	default:
		logger.Error.Printf("The number of parameters should be given as 0 or 2, but %s is given\n", len(addr))
	}

	logger.Info.Println("Use defalt address: localhost:8080")
	return "localhost:8080", nil
}
