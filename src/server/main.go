package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error {
	return ioutil.WriteFile(p.Title + ".txt", p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	// WARNING: This is relative to where this script is called
	// E.g. if you call `go run src/server/main.go`, it will search
	// at the calling directory
	// So if your txt file is at src/server/<text_file>.txt, it will
	// fail.
	body, err := ioutil.ReadFile(title + ".txt")
	if (err != nil) {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	if p, err := loadPage(title); err == nil {
		fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	} else {
		fmt.Fprintf(w, "%s", err)
	}
}

func test_WriteAndReadFile() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}

func test_RunServer() {
	// Order don't matter
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func test_RunWikiPage() {
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	test_RunWikiPage()
}
