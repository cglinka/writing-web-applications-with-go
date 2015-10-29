package main

import (
	// "fmt"
	"io/ioutil"
)

// Define the primary data structure for a wiki, a page.
// A page consists of a title and a body.
type Page struct {
	Title string
	Body  []byte
}

// Create a 'save' method on the Page struct to save the info
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func main() {

}
