package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type (
	CotacaoResponse struct {
		USDBRL USDBRL `json:"USDBRL"`
	}
	USDBRL struct {
		Bid string `json:"bid"`
	}
)

func main() {
	http.HandleFunc("/cotacao", cotacaoHandler)

	fmt.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func cotacaoHandler(w http.ResponseWriter, r *http.Request) {
	data, err := fetchData()
	if err != nil {
		http.Error(w, "erro ao buscar cotacao", http.StatusInternalServerError)
		log.Println("erro ao buscar cotacao:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"cotacao": data})
}

func fetchData() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var cotacao CotacaoResponse
	err = json.NewDecoder(resp.Body).Decode(&cotacao)
	if err != nil {
		return "", err
	}

	return cotacao.USDBRL.Bid, nil
}
