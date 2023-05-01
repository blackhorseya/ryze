package main

import (
	"flag"
	"log"
)

var path = flag.String("c", "./deployments/configs/restful/local.yaml", "set config file path")

func init() {
	flag.Parse()
}

// @title ryze
// @version 0.0.1
// @description ryze is a blockchain explorer for the ryze blockchain
//
// @contact.name sean.zheng
// @contact.email blackhorseya@gmail.com
// @contact.url https://blog.seancheng.space
//
// @license.name GPL-3.0
// @license.url https://spdx.org/licenses/GPL-3.0-only.html
//
// @BasePath /api
func main() {
	app, err := CreateApplication(*path)
	if err != nil {
		log.Fatal(err)
	}

	err = app.Start()
	if err != nil {
		log.Fatal(err)
	}

	err = app.AwaitSignal()
	if err != nil {
		log.Fatal(err)
	}
}
