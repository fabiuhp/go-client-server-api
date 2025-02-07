package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Cotacao string `json:"cotacao"`
}

func main() {
	data, err := fetchCotacao()
	if err != nil {
		fmt.Println("Erro ao obter cotação:", err)
		return
	}

	fmt.Println("Cotação do dólar:", data)

	if err := salvarCotacaoNoArquivo(data); err != nil {
		fmt.Println("Erro ao salvar cotação no arquivo:", err)
	}
}

func fetchCotacao() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}

	return response.Cotacao, nil
}

func salvarCotacaoNoArquivo(data string) error {
	file, err := os.Create("cotacao.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("Dólar: " + data)
	return err
}
