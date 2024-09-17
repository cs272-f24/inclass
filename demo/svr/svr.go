package main

import (
	"fmt"
	//"io"
	"html/template"
	"net/http"
)

type FooHandler http.HandlerFunc

func (fh FooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

type TemplateData struct {
	Elements []string
}

func main() {
	/* First try
	var fh FooHandler
	http.Handle("/foo", fh)
	err := http.ListenAndServe(":8080", nil)
	fmt.Println("ListenAndServe returned %v\n", err)
	*/

	/* second try
	fooHandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "foo!")
	}
	*/

	markup := "{{ range .Elements }} <h2> {{.}} </h2>  {{end}}"

	tmpl, err := template.New("some name").Parse(markup)
	if err != nil {
		fmt.Printf("template.Parse returned %v\n", err)
	}

	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		data := TemplateData{
			Elements: []string{"foo", "bar", "baz"},
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			fmt.Printf("tmpl.Execute returned %v\n", err)
		}
		//io.WriteString(w, "more foo!")
	})

	http.ListenAndServe(":8080", nil)

}
