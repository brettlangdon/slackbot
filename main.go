package main

import (
	"flag"
	"log"

	"github.com/brettlangdon/slackbot/plugins/sqwiggle"
	"github.com/brettlangdon/slackbot/server"
)

var bind string

func main() {
	flag.StringVar(&bind, "bind", "127.0.0.1:8000", "which address to bind to [defualt: 127.0.0.1:8000]")
	flag.Parse()

	s := server.NewServer()

	s.AddPlugin(sqwiggle.NewSqwiggle())

	s.SetListenAddress(bind)
	log.Fatal(s.Start())
}
