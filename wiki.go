package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/ds4tech/pipeline-calculator-ws/calculator_ws"
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
	body, err := ioutil.ReadFile(filename)
	fmt.Println("Loading file " + filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFiles(tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	fmt.Println("Starting server...")
 
	//"strings"
	fmt.Printf("Sum of %d and %d is %d\n", 3,4, calc.Sum(3,4))
	fmt.Printf("Factorial of 4 is %v\n", calc.Factorial(4))
  fmt.Printf("4 is prime: %v\n", calc.IsPrime(4))
  fmt.Printf("3 is prime: %v\n", calc.IsPrime(3))
  fmt.Printf("Log(9) = %v\n", calc.Log(9))
	fmt.Println("End of calculations")

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)
	//	p1 := &Page{Title: "TestPage", Body: []byte("Simple Page.")}
}
