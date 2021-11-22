package app

import (
	"flag"
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
)

// AppEngine struct
type AppEngine struct {
	mainSrv *MainServer
	promSrv *PromServer
	servers errgroup.Group
}

var addr = flag.String("addr", ":8080", "Address:port for HTTP requests.")
var promAddr = flag.String("prom-addr", ":8088", "Address:port for Prometheus requests.")

func init() {
	flag.Parse()
}

// Create function
func Create() (a *AppEngine, err error) {
	a = new(AppEngine)

	a.mainSrv = CreateMainServer()
	a.promSrv, err = CreatePromServer()

	return
}

// Run function
func (a *AppEngine) Run() {
	a.runServer(*addr, a.promSrv.Metrics.Collect(a.mainSrv.Router), "HTTP server")
	a.runServer(*promAddr, a.promSrv.Router, "Metrics server")

	if err := a.servers.Wait(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

// runServer function
func (a *AppEngine) runServer(addr string, srv http.Handler, srvTitle string) {
	s := &http.Server{
		Addr:    addr,
		Handler: srv,
	}

	a.servers.Go(s.ListenAndServe)

	log.Printf("Start %s listening on %s\n", srvTitle, addr)
}
