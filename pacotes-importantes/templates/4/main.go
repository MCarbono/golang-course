package main

import (
	"net/http"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
		err := t.Execute(w, Cursos{
			Curso{"GO", 40},
			Curso{"Java", 10},
			Curso{"Python", 65},
		})
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":3000", nil)
}
