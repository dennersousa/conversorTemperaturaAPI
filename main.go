package main

import (
	"conversorTemperaturaAPI/conversor"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/converter", ConvertHandler)
	fmt.Println("Servidor iniciado na porta 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}

func ConvertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	if r.Method == http.MethodGet {
		// Aqui você pode adicionar um tratamento para solicitações GET, se necessário.
		// Por exemplo, pode renderizar um formulário HTML para inserir os parâmetros.
		// No entanto, por razões de segurança, considera-se melhor realizar operações sensíveis usando POST.

		// Se desejar fornecer um formulário HTML, você pode fazer algo como:
		// w.Write([]byte(`<html><body><form method="post" action="/converter">...</form></body></html>`))
		return
	}

	unidadeOriginal := r.FormValue("unidade_original")
	unidadeConvertida := r.FormValue("unidade_convertida")
	temperaturaInseridaStr := r.FormValue("temperatura_inserida")

	temperaturaInserida, err := strconv.ParseFloat(temperaturaInseridaStr, 64)
	if err != nil {
		http.Error(w, "Erro ao converter a temperatura", http.StatusBadRequest)
		return
	}

	output := w

	conversor.ConverterTemperatura(unidadeOriginal, unidadeConvertida, temperaturaInserida, output)
}
