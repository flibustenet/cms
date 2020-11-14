package hdl

import (
	"fmt"
	"net/http"

	"github.com/flibustenet/cms/app"
)

// Handler envoi la page demandée par http.Request vers http.ResponseWriter
func Handler(w http.ResponseWriter, r *http.Request) {
	nom := r.URL.Path[1:len(r.URL.Path)]
	if nom == "" {
		nom = "index.html"
	}
	err := app.CONF.RenderPage(w, nom)
	if err != nil {
		http.Error(w, fmt.Sprintf("handler pour page %s : %v", nom, err), 500)
	}
}
