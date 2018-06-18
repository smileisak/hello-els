package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("example")

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func initLogger() {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	backendLeveled := logging.AddModuleLevel(backend)
	backendLeveled.SetLevel(logging.INFO, "")
	logging.SetBackend(backendLeveled, backendFormatter)
}

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	initLogger()
	log.Notice("Server Listeing on 0.0.0.0:8080")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Notice("Request Handled: URI %s", r.URL.Path)
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
		fmt.Fprintf(w, "Pod handling this request: %s \n", hostname)
		log.Notice("Request Handled: URI %s", r.URL.Path)
	})
	http.ListenAndServe(":8080", nil)
}
