package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("template/*index.html"))
}
func Homehandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}
func Contacthandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome to contact page")
}

func Abouthandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome to about page")
}
func Signuphandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome to signup page")
}
func Errorhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "error")
}
func Productshandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "products")
}
func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	http.HandleFunc("/", Homehandler)
	http.HandleFunc("/Contact", Contacthandler)
	http.HandleFunc("/products", Productshandler)
	http.HandleFunc("/about", Abouthandler)
	http.HandleFunc("/signup", Signuphandler)
	http.HandleFunc("/error", Errorhandler)
	fmt.Println("starting server at port :5555")
	http.ListenAndServe(":5555", nil)
}
