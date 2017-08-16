package vidar

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

type Vidar struct {
	Router   *Router
	Server   *http.Server
	Listener net.Listener
	address  string
}

type vidarListener struct {
	*net.TCPListener
}

func New() (v *Vidar) {
	v = &Vidar{
		Router: NewRouter(),
		Server: new(http.Server),
	}

	v.Server.Handler = v.Router

	return
}

//TODO: Implement cgi and fast cgi interface
func (v *Vidar) Run() (err error) {
	v.Server.Addr, err = resolveAddress()
	if err != nil {
		fmt.Printf("Resolve address failed: %v\n", err)
	}

	fmt.Printf("Running on %s", v.Server.Addr)

	if err := v.StartServer(v.Server); err != nil {
		fmt.Printf("Server start failed: %v", err)
	}

	return nil
}

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
		if host := Config.Server.Host; len(host) > 0 {
			fmt.Printf("Read host value from config file: %s\n", host)

			if port := Config.Server.Port; len(port) > 0 {
				fmt.Printf("Read port value from config file: %s\n", port)

				return host + ":" + port, nil
			}
		}
	case 2:
		return strings.Join(addr, ":"), nil
	default:
		fmt.Printf("The number of parameters should be given as 0 or 2, but %s is given\n", len(addr))
	}

	fmt.Println("Use defalt address: localhost:8080")
	return "localhost:8080", nil
}
