package main

import (
	"html/template"
	"os"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}
func main() {
	templates := []string{
		"header.html",
		"content.html",
	}
	t := template.New("content.html")
	t.Funcs(template.FuncMap{
		"toUpper": ToUpper,
	})
	t = template.Must(t.ParseFiles(templates...))
	err := t.Execute(os.Stdout, Cursos{
		Curso{"GO", 40},
		Curso{"Java", 10},
		Curso{"Python", 65},
	})
	if err != nil {
		panic(err)
	}
}
