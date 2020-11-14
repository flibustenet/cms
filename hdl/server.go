package hdl

import (
	"fmt"
	"net/http"

	"github.com/flibustenet/cms/app"
)

// RunServer lance le serveur
func RunServer(conf *app.Conf) error {
	http.HandleFunc("/", Middle(Handler))
	fmt.Println("Listen " + conf.Listen)
	err := http.ListenAndServe(conf.Listen, nil)
	if err != nil {
		return fmt.Errorf("Run serveur %v", err)
	}
	return nil
}
