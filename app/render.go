package app

import (
	"errors"
	"fmt"
	"io"
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
