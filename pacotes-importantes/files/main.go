package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	f, err := os.Create("files/arquivo.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	size, err := f.Write([]byte("Hello, World! dsadsda adasdeqeqwewq fdfdfdfdfdreerewrerw"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso. Tamanho: %d\n", size)
	file, err := os.ReadFile("files/arquivo.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(file))

	//leitura de pouco em pouco abrindo o arquivo
	fileLines, err := os.Open("files/arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(fileLines)
	buffer := make([]byte, 3)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}
	time.Sleep(time.Second)
	err = os.Remove("files/arquivo.txt")
	if err != nil {
		panic(err)
	}
}
