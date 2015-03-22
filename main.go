package main

import (
	"log"
	"time"
)

func main() {
	p := NewPage()
	defer p.driver.Close()
	p.Refresh()

	for {
		p.dumpPage()

		b := p.Available()
		if b {
			url, err := p.Start()
			if err != nil {
				log.Printf("Failed to get registration page => %s", err)
				continue
			}
			alert(url)
			return
		}

		err := p.Refresh()
		if err != nil {
			log.Fatalf("Failed to refresh connection => {%s}", err)
		}

		time.Sleep(5 * time.Second)
	}
}
