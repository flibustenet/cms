package hdl

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/flibustenet/cms/app"
)

// Handler envoi la page demandée par http.Request vers http.ResponseWriter
func Handler(w http.ResponseWriter, r *http.Request) {
	nom := r.URL.Path[1:len(r.URL.Path)]
	if nom == "" {
		nom = "index.html"
	}

	// routage des fichiers statics
	// pourrait être mis sur le mux en amont
	if strings.HasPrefix(nom, "static/") {
		http.ServeFile(w, r, filepath.Join(app.CONF.Root, nom))
		return
	}

	// Utilise un buffer tampon pour envoyer le code 500 en cas d'erreur
	// avant les données
	// pourrait être effectué dans un middleware
	var buf bytes.Buffer
	err := app.CONF.RenderPage(&buf, nom)
	if errors.Is(err, app.ErrNotFound) {
		http.Error(w, err.Error(), 404)
		return
	}

	if err != nil {
		http.Error(w, fmt.Sprintf("handler pour page %s : %v", nom, err), 500)
	}

	io.Copy(w, &buf)
}
