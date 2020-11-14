package app

import (
	"html/template"

	"github.com/russross/blackfriday/v2"
)

// Version v2 à indiquer à la fin

// markdown renvoi le texte sous forme HTML
func markdown(s string) template.HTML {
	output := blackfriday.Run([]byte(s))
	return template.HTML(output)
}

// fonctions pour servir dans les templates
var fmap = template.FuncMap{
	"MD": markdown,
}
