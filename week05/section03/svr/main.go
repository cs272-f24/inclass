package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type SearchHit struct {
	URL   string
	Count int
}

func main() {
	//http.HandleFunc("/top10", )
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		//term := r.URL.Query().Get("term")
		w.Header().Set("Content-Type", "text/html")

		hits := []SearchHit{
			{"/foo.html", 5},
			{"/bar.html", 6},
		}

		tmpl, err := template.New("demo.tmpl").ParseFiles("demo.tmpl")
		if err != nil {
			fmt.Printf("Parse returned %v\n", err)
		}
		err = tmpl.Execute(w, hits)
		if err != nil {
			fmt.Println("Execute returned ", err)
		}
	})

	http.Handle("/", http.FileServer(http.Dir("static")))
	http.ListenAndServe(":8080", nil)
	fmt.Println("here I am")
}
