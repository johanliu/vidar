package vidar

import (
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/johanliu/mlog"
)

type Vidar struct {
	Router   *Router
	Server   *http.Server
	Listener net.Listener
	Chain    []Chain
	address  string
}

type vidarListener struct {
	*net.TCPListener
}

var log = mlog.NewLogger()

func New() (v *Vidar) {
	v = &Vidar{
		Router: NewRouter(),
		Server: new(http.Server),
	}

	v.Server.Handler = v.Router
	log.SetLevelByName("INFO")

	return
}

//TODO: Implement cgi and fast cgi interface
func (v *Vidar) Run() (err error) {
	v.Server.Addr, err = resolveAddress()
	if err != nil {
		log.Error(err)
	}

	log.Info("Running on %s", v.Server.Addr)

	if err := v.StartServer(v.Server); err != nil {
		log.Error(err)
	}

	return nil
}

// TODO: should be used by unix domain socket as well
func (v *Vidar) StartServer(s *http.Server) (err error) {
	v.Listener, err = newListener("tcp", v.Server.Addr)
	if err != nil {
		return err
	}

	return s.Serve(v.Listener)
}

func (vl *vidarListener) Accept() (c net.Conn, err error) {
	tc, err := vl.AcceptTCP()
	if err != nil {
		return
	}

	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(5 * time.Minute)
	return tc, nil
}

func newListener(proto string, address string) (*vidarListener, error) {
	l, err := net.Listen(proto, address)
	if err != nil {
		return nil, err
	}
	return &vidarListener{l.(*net.TCPListener)}, nil
}

func resolveAddress(addr ...string) (string, error) {
	switch len(addr) {
	case 0:
		if host := tc.Server.Host; len(host) > 0 {
			if port := tc.Server.Port; len(port) > 0 {
				return host + ":" + port, nil
			}
		}
	case 2:
		return strings.Join(addr, ":"), nil
	default:
		log.Info("The number of parameters should be given as 0 or 2, but %s is given\n", len(addr))
	}

	log.Info("Use defalt address: localhost:8080")
	return "localhost:8080", nil
}
