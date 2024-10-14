package main

import (
	"flag"
	"keycloak-example/confs"
	"keycloak-example/web"
)

var f = flag.String("f", "confs.yaml", "load configuration from yaml file")

func main() {
	flag.Parse()
	
	// Load config
	if err := confs.Load(*f); err != nil {
		panic(err)
	}

	// Init webserver
	if err := web.Setup(); err != nil {
		panic(err)
	}
}
