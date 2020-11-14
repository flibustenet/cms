package hdl

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/flibustenet/cms/app"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Middle : middleware de log et rattrapge erreurs
// voir par la suite https://www.gorillatoolkit.org/
func Middle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wait := 0
		if app.CONF.Slow > 0 {
			wait = rand.Intn(app.CONF.Slow)
		}

		log.Printf("%s %s %ds", r.Method, r.URL.Path, wait)
		time.Sleep(time.Duration(wait) * time.Second)
		defer func() {
			// nécessiterait un buffer sur http.ResponseWriter...
			if rec := recover(); rec != nil {
				log.Printf("Erreur : %s", rec)
				http.Error(w, "Erreur grave", 500)
			}
		}()
		next(w, r) // original function call
	}
}
