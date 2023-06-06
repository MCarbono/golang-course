package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	return a == b
}

func main() {
	m := map[string]int{"Wesley": 1000, "João": 2000, "Marcelo": 3000}
	println(Soma(m))
	m2 := map[string]float64{"Wesley": 1000.5, "João": 2000.5, "Marcelo": 3000.5}
	println(Soma(m2))
	m3 := map[string]MyNumber{"Wesley": 1000, "João": 2000, "Marcelo": 3000}
	println(Soma(m3))

	println(Compara(10, 10))
}
