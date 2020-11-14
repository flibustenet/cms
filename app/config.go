package app

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"
	"text/template"
)

// ConfJSON contient la configuration du site
// majuscule pour export éventuel
type ConfJSON struct {
	Pages []*Page // Majuscules pour export json et autre éventuel
}

// confUtil champs utilitaires
type confUtil struct {
	tmpl *template.Template
}

// Conf est composé de ConfJson et des champs utilitaires de confUtil
type Conf struct {
	ConfJSON
	confUtil
}

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

	// compilation des templates
	// ils se trouvent dans ../templates par rapport à ici
	_, b, _, _ := runtime.Caller(0)
	ici := path.Join(path.Dir(b))
	pathTemplate := filepath.Join(ici, "..", "templates", "*.html")
	t, err := template.ParseGlob(pathTemplate)
	if err != nil {
		return nil, fmt.Errorf("parsing templates %s : %v", pathTemplate, err)
	}
	conf.tmpl = t

	return conf, nil
}
