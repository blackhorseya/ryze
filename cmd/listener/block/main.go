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
	service, err := CreateService(*path)
	if err != nil {
		log.Fatal(err)
	}
	_ = service
}
