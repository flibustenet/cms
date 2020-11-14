package app

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"
)

// ConfJSON contient la configuration du site
// majuscule pour export éventuel
type ConfJSON struct {
	Pages  []*Page // Majuscules pour export json et autre éventuel
	Listen string  // Ecoute du serveur
	Root   string  // chemin où se trouve ./templates
}

// confUtil champs utilitaires
type confUtil struct {
	tmpl     *template.Template
	mapPages map[string]*Page // contient les pages par leur nom
}

// Conf est composé de ConfJson et des champs utilitaires de confUtil
type Conf struct {
	ConfJSON
	confUtil
}

// CONF contient la configuration du site accessible globalement
var CONF *Conf

// ReadConf lit la configuration à partir d'un io.Reader
// L'interface permet d'utiliser n'importe quel objet
// implémentant :
//	   Read(p []byte) (n int, err error)
// chaque erreur est annotée et renvoyée
// compilation des templates
func ReadConf(file io.Reader) (*Conf, error) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("lecture conf : %v", err)
	}
	// parse le json
	conf := &Conf{}
	err = json.Unmarshal(data, conf)
	if err != nil {
		return nil, fmt.Errorf("unmarshal conf : %v", err)
	}

	// indexation des pages
	conf.mapPages = map[string]*Page{}
	for _, page := range conf.Pages {
		conf.mapPages[page.Nom] = page
	}

	// compilation des templates
	if conf.Root == "" {
		// ils se trouvent dans ../templates par rapport à ici
		_, b, _, _ := runtime.Caller(0)
		conf.Root = filepath.Join(path.Dir(b), "..", "templates")
	}
	pathTemplate := filepath.Join(conf.Root, "*.html")

	t, err := template.New("").Funcs(fmap).ParseGlob(pathTemplate)
	if err != nil {
		return nil, fmt.Errorf("parsing templates %s : %v", pathTemplate, err)
	}
	conf.tmpl = t

	CONF = conf

	return conf, nil
}
