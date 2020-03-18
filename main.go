package main

import (
	"fmt"
	"log"
	"novelweb/app"
	"novelweb/config"
)

func main() {
	generate()
}

func generate() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	app := app.Init()
	listenAddress := fmt.Sprintf("%s:%d", config.GetConfig().Host, config.GetConfig().Port)
	fmt.Println(" main address ", listenAddress)
	if err := app.Run(listenAddress); err != nil {
		log.Fatal("run app, but meet err: ", err)
	}
}
