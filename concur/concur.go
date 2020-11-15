package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/flibustenet/cms/app"
)

func main() {
	file, err := os.Open("../config.json")
	if err != nil {
		log.Fatalf("Impossible d'ouvrir la config : %v", err)
	}
	conf, err := app.ReadConf(file)
	if err != nil {
		log.Fatalf("Impossible de lire la config : %v", err)
	}
	avecWaitGroup(conf)
	avecChan(conf)
	avecChanSelect(conf)
}

// Get des pages en goroutine
func avecWaitGroup(conf *app.Conf) {
	log.Println("--- Avec WaitGroup")
	var wg sync.WaitGroup

	for _, p := range conf.Pages {
		url := "http://" + conf.Listen + "/" + p.Nom
		log.Printf("get %s", url)
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			start := time.Now()
			resp, err := http.Get(url)
			if err != nil {
				log.Printf("GET %s : %v", url, err)
			}
			if resp.StatusCode != 200 {
				log.Printf("Mauvais status %s : %d %s", url, resp.StatusCode, resp.Status)
			}
			log.Printf("Reçoit %s en %s", url, time.Since(start))
			resp.Body.Close()
		}(url)
	}
	wg.Wait()
}

// Get des pages en channel

// Msg est passé dans le chan
type Msg struct {
	url   string
	err   error
	start time.Time
}

func (m *Msg) String() string {
	if m.err != nil {
		return fmt.Sprintf("Erreur sur %s : %v", m.url, m.err)
	}
	return m.url
}

func avecChan(conf *app.Conf) error {
	log.Println("--- Avec Chan")
	msgCh := make(chan *Msg)

	for _, p := range conf.Pages {
		url := "http://" + conf.Listen + "/" + p.Nom
		log.Printf("get %s", url)
		go func(url string) {
			now := time.Now()
			resp, err := http.Get(url)
			if err != nil {
				msgCh <- &Msg{url, err, now}
				return
			}
			defer resp.Body.Close()
			if resp.StatusCode != 200 {
				msgCh <- &Msg{url, fmt.Errorf("Status %d", resp.StatusCode), now}
				return
			}
			msgCh <- &Msg{url, nil, now}
		}(url)
	}

	for range conf.Pages {
		msg := <-msgCh
		log.Printf("Reçoit %s en %s", msg.url, time.Since(msg.start))
		if msg.err != nil {
			return msg.err
		}
	}
	return nil
}

func avecChanSelect(conf *app.Conf) error {
	log.Println("--- Avec Chan et select")
	msgCh := make(chan *Msg)

	for _, p := range conf.Pages {
		url := "http://" + conf.Listen + "/" + p.Nom
		log.Printf("get %s", url)
		go func(url string) {
			now := time.Now()
			resp, err := http.Get(url)
			if err != nil {
				msgCh <- &Msg{url, err, now}
				return
			}
			defer resp.Body.Close()
			if resp.StatusCode != 200 {
				msgCh <- &Msg{url, fmt.Errorf("Status %d", resp.StatusCode), now}
				return
			}
			msgCh <- &Msg{url, nil, now}
		}(url)
	}

	tick := time.Tick(1 * time.Second)
	nb := 0
	for nb < len(conf.Pages) {
		select {
		case msg := <-msgCh:
			log.Printf("Reçoit %s en %s", msg.url, time.Since(msg.start))
			if msg.err != nil {
				return msg.err
			}
			nb++
		case <-tick:
			log.Printf("wait...")

		}
	}
	return nil
}
