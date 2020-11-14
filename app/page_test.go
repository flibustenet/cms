package app

import (
	"bytes"
	"strings"
	"testing"
)

func TestPageRender(t *testing.T) {
	conf, err := readConfTest()
	if err != nil {
		t.Errorf("appel readConfTest : %v", err)
	}
	var bf bytes.Buffer
	err = conf.Pages[0].Render(conf, &bf)
	if err != nil {
		t.Errorf("Render pages 0 : %v", err)
	}
	res := bf.String()
	if !strings.Contains(res, "<!doctype html>") {
		t.Errorf("Ne contient pas doctype : %s", res)
	}
	if !strings.Contains(res, "Bienvenue") {
		t.Errorf("Ne contient pas Bienvenue : %s", res)
	}
}
