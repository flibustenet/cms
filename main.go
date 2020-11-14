package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/flibustenet/cms/app"
	"github.com/flibustenet/cms/hdl"
)

func main() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalf("Impossible d'ouvrir la config : %v", err)
	}
	conf, err := app.ReadConf(file)
	if err != nil {
		log.Fatalf("Impossible de lire la config : %v", err)
	}
	fmt.Printf("Conf = %+v\n", conf)

	http.HandleFunc("/", hdl.Handler)
	fmt.Printf("Listen %s\n", conf.Listen)
	panic(http.ListenAndServe(conf.Listen, nil))
}
