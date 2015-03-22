package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"bitbucket.org/tebeka/selenium"
)

var (
	eventURL = envVar("EVENT_URL")
	// FireFox driver without specific version
	caps = selenium.Capabilities{"browserName": "firefox"}
)

type page struct {
	driver selenium.WebDriver
}

func NewPage() page {
	wd, err := selenium.NewRemote(caps, "")
	if err != nil {
		log.Fatalf("failed to connect to selenium => {%s}", err)
	}

	return page{
		driver: wd,
	}
}

func (p page) Refresh() error {
	err := p.driver.Refresh()
	if err != nil {
		return fmt.Errorf("failed to connect to selenium => {%s}", err)
	}

	err = p.driver.Get(eventURL)
	if err != nil {
		return fmt.Errorf("failed to connect to selenium => {%s}", err)
	}

	return nil
}

func (p page) Available() bool {
	elem, err := p.driver.FindElement(selenium.ByCSSSelector, "#primary_cta")
	if err != nil {
		return false
	}

	displayed, err := elem.IsDisplayed()
	if !displayed {
		return false
	}

	return true
}

func (p page) dumpPage() {
	src, _ := p.driver.PageSource()
	filename := fmt.Sprintf("logs/source-%s.html", time.Now().Format(time.RFC3339))
	ioutil.WriteFile(filename, []byte(src), os.ModePerm)
}

func (p page) Start() (string, error) {
	elem, err := p.driver.FindElement(selenium.ByCSSSelector, "#primary_cta")
	if err != nil {
		return "", fmt.Errorf("Should have register button but failed => {%s}", err)
	}

	err = elem.Click()
	if err != nil {
		return "", fmt.Errorf("Should have register button but failed => {%s}", err)
	}

	_, err = p.driver.ExecuteScript("freeCheckout();", []interface{}{})
	if err != nil {
		log.Printf("Failed to call JS => {%s}", err)
	}

	url, err := p.driver.CurrentURL()
	if err != nil {
		p.dumpPage()
		return "", fmt.Errorf("failed to get current URL => {%s}", err)
	}
	log.Printf("new page: %s", url)

	return url, nil
}
