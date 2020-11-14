package hdl

import (
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/flibustenet/cms/app"
)

var confTest = `
{
	"listen":":8888",
    "pages":[
        {"Menu":"Accueil", "Nom":"index.html"},
        {"Menu":"Contact", "Nom":"contact.html"},
        {"Menu":"Plan", "Nom":"plan.html"}
    ]
}
`

// copy ici car il n'est pas possible d'utiliser une
// méthode de test d'un autre package
func readConfTest() (*app.Conf, error) {
	conf, err := app.ReadConf(strings.NewReader(confTest))
	if err != nil {
		return nil, fmt.Errorf("lecture conf test : %v", err)
	}
	return conf, nil
}

func TestHandler(t *testing.T) {
	conf, err := readConfTest()
	if err != nil {
		t.Errorf("lecture conf test sur test config : %v", err)
	}
	for _, page := range conf.Pages {
		r := httptest.NewRequest("GET", "/"+page.Nom, nil)
		w := httptest.NewRecorder()
		Handler(w, r)
		if w.Code != 200 {
			t.Errorf("GET / status=%d pour page %s body : %s", w.Code, page.Nom, w.Body)
		}
	}
	r := httptest.NewRequest("GET", "/xyz", nil)
	w := httptest.NewRecorder()
	Handler(w, r)
	if w.Code != 404 {
		t.Errorf("GET / status=%d pour page inexistante", w.Code)
	}

}

func TestHandlerStatic(t *testing.T) {
	_, err := readConfTest()
	if err != nil {
		t.Errorf("lecture conf test sur test config : %v", err)
	}
	r := httptest.NewRequest("GET", "/static/style.css", nil)
	w := httptest.NewRecorder()
	Handler(w, r)
	if w.Code != 200 {
		t.Errorf("GET / status=%d pour /static/style.css : %s", w.Code, w.Body)
	}

	r = httptest.NewRequest("GET", "/static/xyz", nil)
	w = httptest.NewRecorder()
	Handler(w, r)
	if w.Code != 404 {
		t.Errorf("GET / status=%d pour page static inexistante", w.Code)
	}
}
