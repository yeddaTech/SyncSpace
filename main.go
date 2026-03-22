package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	// 1. Diciamo a Go di servire i file statici (il nostro CSS) dalla cartella /static/
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// 2. Creiamo la rotta principale ("/") che carica il template HTML
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
	})

	// 3. Avviamo il server sulla porta 8080
	log.Println("Server avviato su http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
