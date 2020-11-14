package app

import (
	"fmt"
	"strings"
	"testing"
)

var confTest = `
{
    "pages":[
        {"Menu":"Accueil", "Nom":"index.html"},
        {"Menu":"Contact", "Nom":"contact.html"},
        {"Menu":"Plan", "Nom":"plan.html"}
    ]
}
`

func readConfTest() (*Conf, error) {
	conf, err := ReadConf(strings.NewReader(confTest))
	if err != nil {
		return nil, fmt.Errorf("lecture conf test : %v", err)
	}
	return conf, nil
}

func TestConfig(t *testing.T) {
	conf, err := readConfTest()
	if err != nil {
		t.Fatalf("lecture conf test : %v", err)
	}

	if len(conf.Pages) != 3 {
		t.Errorf("Pages != 3 %+v", conf)
	}
	if conf.Pages[1].Menu != "Contact" {
		t.Errorf("Page 2 menu : %v", conf.Pages[1])
	}
}
