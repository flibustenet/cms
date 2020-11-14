package hdl

import (
	"net/http"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	conf, err := readConfTest()
	if err != nil {
		t.Fatalf("appel readConfTest : %v", err)
	}

	go func() {
		err := RunServer(conf)
		if err != nil {
			t.Fatalf("lance serveur : %v", err)
		}
	}()
	time.Sleep(500 * time.Millisecond)

	for _, p := range conf.Pages {
		url := "http://" + conf.Listen + "/" + p.Nom
		resp, err := http.Get(url)
		if err != nil {
			t.Errorf("GET %s : %v", url, err)
		}
		if resp.StatusCode != 200 {
			t.Errorf("Mauvais status %s : %d %s", url, resp.StatusCode, resp.Status)
		}
		resp.Body.Close()
	}
}
