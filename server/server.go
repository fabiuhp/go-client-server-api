package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "modernc.org/sqlite"
)

type (
	CotacaoResponse struct {
		USDBRL USDBRL `json:"USDBRL"`
	}
	USDBRL struct {
		Bid string `json:"bid"`
	}
)

var db *sql.DB

func iniciarBanco() error {
	var err error
	db, err = sql.Open("sqlite", "file:cotacoes.db?mode=rwc")
	if err != nil {
		return err
	}

	query := `
	CREATE TABLE IF NOT EXISTS cotacoes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		bid TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err = db.Exec(query)
	return err
}

func main() {
	if err := iniciarBanco(); err != nil {
		log.Fatal("erro ao inicializar banco de dados", err)
	}
	defer db.Close()

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

	if err := salvarCotacao(data); err != nil {
		log.Println("erro ao salvar no banco:", err)
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

func salvarCotacao(bid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	_, err := db.ExecContext(ctx, "INSERT INTO cotacoes (bid) VALUES (?)", bid)
	if err != nil {
		return err
	}

	return nil
}
