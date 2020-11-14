package app

import (
	"fmt"
	"io"
)

// Page contient les informations d'une page
// exportable avec la majuscule
type Page struct {
	Menu string
	Nom  string
}

// String permet d'afficher des informations de debug sur la page
func (p *Page) String() string {
	return fmt.Sprintf("Page menu=%s nom=%s", p.Menu, p.Nom)
}

// Render envoi le rendu du template du même nom sur le writer
func (p *Page) Render(conf *Conf, w io.Writer) error {
	err := conf.tmpl.ExecuteTemplate(w, p.Nom, map[string]interface{}{
		"Conf": conf,
		"Page": p,
	})
	if err != nil {
		return fmt.Errorf("execute template %s : %v", p.Nom, err)
	}
	return nil
}
