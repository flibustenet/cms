package app

import "fmt"

// Page contient les informations d'une page
// exportable avec la majuscule
type Page struct {
	Menu string
	Nom  string
}

// String permet d'affiche des informations de debug sur la page
func (p *Page) String() string {
	return fmt.Sprintf("Page menu=%s nom=%s", p.Menu, p.Nom)
}
