package main

import (
	"log"
	"time"
)

func main() {
	// setup page /w selenium
	p := NewPage()
	defer p.driver.Close()
	p.Refresh()

	for {
		// write page to file in case we need to recalibrate later
		p.dumpPage()

		// if the "register" button is present
		if p.Available() {
			// try to complete registration
			url, err := p.Start()
			if err != nil {
				log.Printf("Failed to get registration page => %s", err)
				continue
			}

			// send alerts to user
			alert(url)
			return
		}

		// refresh the page
		err := p.Refresh()
		if err != nil {
			log.Fatalf("Failed to refresh connection => {%s}", err)
		}

		time.Sleep(time.Second)
	}
}
