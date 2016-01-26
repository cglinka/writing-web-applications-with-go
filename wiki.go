package main

import (
	"fmt"
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

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))

}
