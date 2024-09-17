package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Hit struct {
	URL   string // remember to use upper case for tmpl str
	Count int
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		//searchTerm := r.URL.Query().Get("term")
		w.Header().Set("Content-Type", "text/html")

		hits := []Hit{
			{
				"/foo.html", 3,
			},
			{
				"/bar.html", 4,
			},
		}
		tmplBody := "<ol> {{range .}} <li>{{.URL}} {{.Count}}</li> {{end}} </ol>"
		tmpl, err := template.New("demo").Parse(tmplBody)
		if err != nil {
			fmt.Printf("template.Parse returned %v\n", err)
		}
		tmpl.Execute(w, hits)
		//w.Write([]byte("search term was: " + searchTerm))
	})
	http.ListenAndServe(":8080", nil)
}
