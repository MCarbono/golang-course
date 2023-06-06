package main

// type ViaCEP struct {
// 	Cep         string `json:"cep"`
// 	Logradouro  string `json:"logradouro"`
// 	Complemento string `json:"complemento"`
// 	Bairro      string `json:"bairro"`
// 	Localidade  string `json:"localidade"`
// 	Uf          string `json:"uf"`
// 	Ibge        string `json:"ibge"`
// 	Gia         string `json:"gia"`
// 	Ddd         string `json:"ddd"`
// 	Siafi       string `json:"siafi"`
// }

// func (v ViaCEP) String() string {
// 	var s strings.Builder
// 	s.WriteString(fmt.Sprintf("CEP: %v\n", v.Cep))
// 	s.WriteString(fmt.Sprintf("Localidade: %v\n", v.Localidade))
// 	s.WriteString(fmt.Sprintf("UF: %v\n", v.Uf))
// 	return s.String()
// }

// func cep() {
// 	for _, cep := range os.Args[1:] {
// 		req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "erro ao fazer req %v\n", err)
// 		}
// 		defer req.Body.Close()
// 		res, err := io.ReadAll(req.Body)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "erro ao ler a req %v\n", err)
// 		}
// 		cep := ViaCEP{}
// 		err = json.Unmarshal(res, &cep)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "erro ao fazer o parse de json %v\n", err)
// 		}
// 		f, err := os.Create("cidade.txt")
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "erro ao criar arquivo %v\n", err)
// 		}
// 		defer f.Close()
// 		_, err = f.Write([]byte(cep.String()))
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "erro ao escrever no  arquivo %v\n", err)
// 		}
// 	}
// }
