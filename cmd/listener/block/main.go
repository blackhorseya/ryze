package main

import (
	"flag"
	"log"
)

var path = flag.String("c", "./deployments/configs/listener/block/local.yaml", "set config file path")

func init() {
	flag.Parse()
}

func main() {
	app, err := CreateApplication(*path)
	if err != nil {
		log.Fatal(err)
	}
	_ = app
}
