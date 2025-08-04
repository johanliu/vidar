package vidar

import (
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/johanliu/mlog"
)

var log *mlog.Logger = nil

// Vidar represents the main web framework instance.
// It contains the router, HTTP server, listener, plugin system, and configuration.
type Vidar struct {
	Router   *Router      // HTTP request router
	Server   *http.Server // Underlying HTTP server
	Listener net.Listener // Network listener
	Plugin   *Plugin      // Plugin/middleware system
	address  string       // Server address
	log      *mlog.Logger // Logger instance
}

// vidarListener wraps a TCP listener with keep-alive functionality
type vidarListener struct {
	*net.TCPListener
}

// New creates a new Vidar instance with default configuration
func New() *Vidar {
	log = mlog.NewLogger()
	log.SetLevelByName("INFO")

	v := &Vidar{
		Router: NewRouter(),
		Server: new(http.Server),
		Plugin: NewPlugin(),
		log:    log,
	}

	v.Server.Handler = v.Router

	return v
}

// TODO: Implement cgi and fast cgi interface
// Run starts the HTTP server and listens for incoming requests.
// It resolves the server address and starts the HTTP server.
func (v *Vidar) Run() (err error) {
	v.Server.Addr, err = v.resolveAddress()
	if err != nil {
		v.log.Error(err)
	}

	v.log.Info("Running on %s", v.Server.Addr)

	if err := v.StartServer(v.Server); err != nil {
		v.log.Error(err)
	}

	return nil
}

// TODO: should be used by unix domain socket as well
// StartServer starts the HTTP server with the given configuration.
// It creates a new listener and serves HTTP requests.
func (v *Vidar) StartServer(s *http.Server) (err error) {
	v.Listener, err = newListener("tcp", v.Server.Addr)
	if err != nil {
		return err
	}

	return s.Serve(v.Listener)
}

// Accept accepts a TCP connection and sets keep-alive parameters.
func (vl *vidarListener) Accept() (c net.Conn, err error) {
	tc, err := vl.AcceptTCP()
	if err != nil {
		return
	}

	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(5 * time.Minute)
	return tc, nil
}

// newListener creates a new vidarListener with the specified protocol and address.
func newListener(proto string, address string) (*vidarListener, error) {
	l, err := net.Listen(proto, address)
	if err != nil {
		return nil, err
	}
	return &vidarListener{l.(*net.TCPListener)}, nil
}

// resolveAddress determines the server address from configuration or defaults.
// It accepts 0 or 2 parameters for host and port.
func (v *Vidar) resolveAddress(addr ...string) (string, error) {
	switch len(addr) {
	case 0:
		// Try to get config from global config if available
		if tc.Server.Host != "" && tc.Server.Port != "" {
			return tc.Server.Host + ":" + tc.Server.Port, nil
		}
		// Fall back to default
		v.log.Info("Use default address: 0.0.0.0:8080")
		return "0.0.0.0:8080", nil
	case 2:
		return strings.Join(addr, ":"), nil
	default:
		v.log.Info("The number of parameters should be given as 0 or 2, but %d is given", len(addr))
		return "0.0.0.0:8080", nil
	}
}
