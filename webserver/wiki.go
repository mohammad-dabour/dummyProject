package main

import (
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {

	filename := title + ".txt"
	if body, err := ioutil.ReadFile(filename); err != nil {
		return nil, err
	} else {
		return &Page{Title: title, Body: body}, nil
	}

}

func main() {

	runServer()

}
