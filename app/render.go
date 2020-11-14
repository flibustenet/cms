package app

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
)

// ErrNotFound est définie pour identifier un 404
var ErrNotFound = errors.New("Page introuvable")

// RenderPage trouve la page et envoi le rendu sur le writer
// Une méthode peut se trouver dans un fichier séparé
func (c *Conf) RenderPage(w io.Writer, nom string) error {
	page, ok := c.mapPages[nom]
	if !ok {
		return ErrNotFound
	}

	err := page.Render(c, w)
	if err != nil {
		return fmt.Errorf("Rendu de la page %s : %v", nom, err)
	}
	return nil
}

// GetArticle renvoi le contenu d'un article trouvé sur disque
func (c *Conf) GetArticle(name string) string {
	path := filepath.Join(c.Root, name)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err.Error()
	}
	return string(data)
}
