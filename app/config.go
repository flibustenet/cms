package app

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

// Conf contient la configuration du site
// majuscule pour export éventuel
type Conf struct {
	Pages []*Page // Majuscules pour export json et autre éventuel
}

// ReadConf lit la configuration à partir d'un io.Reader
// L'interface permet d'utiliser n'importe quel objet
// implémentant :
//	   Read(p []byte) (n int, err error)
// chaque erreur est annotée et renvoyée
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
	return conf, nil
}
