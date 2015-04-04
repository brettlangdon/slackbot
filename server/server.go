package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/brettlangdon/slackbot/plugins"
	"github.com/brettlangdon/slackbot/slack"
)

type Server struct {
	plugs   map[string]plugins.Plugin
	address string
	server  *http.Server
}

func NewServer() *Server {
	server := &Server{
		address: "0.0.0.0:80",
	}

	server.plugs = make(map[string]plugins.Plugin)

	return server
}

func (this *Server) SetListenAddress(address string) {
	this.address = address
}

func (this *Server) AddPlugin(plug plugins.Plugin) {
	this.plugs[strings.ToLower(plug.GetName())] = plug
}

func (this *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// we always want to write back something
	defer fmt.Fprintf(w, "")
	path := strings.Trim(r.URL.Path, "/")
	if path == "" {
		log.Println("Empty path given")
		return
	}
	args := strings.Split(path, "/")

	name := strings.ToLower(args[0])
	var plug plugins.Plugin
	var ok bool
	if plug, ok = this.plugs[name]; !ok {
		log.Println("No plugin gin found with name: ", name)
		return
	}

	apiRequest := slack.ParseAPIRequest(r)
	apiResponse, err := plug.HandleRequest(apiRequest, args)

	if err != nil {
		log.Println(err)
		return
	}

	content, err := json.Marshal(apiResponse)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Fprintf(w, "%s", content)
}

func (this *Server) Start() error {
	this.server = &http.Server{
		Addr:    this.address,
		Handler: this,
	}

	log.Println("Listening to: ", this.address)
	return this.server.ListenAndServe()
}
