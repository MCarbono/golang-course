package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, Cursos{
		Curso{"GO", 40},
		Curso{"Java", 10},
		Curso{"Python", 65},
	})
	if err != nil {
		panic(err)
	}
}
