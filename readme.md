# Aplicativo de Cotação de Dólar (GoLang)

Este projeto foi desenvolvido como parte da pós-graduação em GoLang. Ele consiste em uma aplicação cliente-servidor que obtém a cotação do dólar em relação ao real, armazena no banco de dados SQLite e permite que um cliente recupere e salve essa informação em um arquivo.

## Tecnologias Utilizadas
- **Linguagem:** GoLang
- **Banco de Dados:** SQLite
- **API de Cotação:** [AwesomeAPI - USD-BRL](https://economia.awesomeapi.com.br/json/last/USD-BRL)

## Estrutura do Projeto

O projeto é dividido em dois arquivos principais:

1. **server.go**:
   - Inicializa um servidor HTTP na porta `8080`.
   - Obtém a cotação do dólar através da API externa.
   - Armazena os valores no banco SQLite.
   - Responde a requisições no endpoint `/cotacao`.

2. **client.go**:
   - Faz uma requisição ao servidor para obter a cotação.
   - Exibe a cotação no terminal.
   - Salva a cotação em um arquivo `cotacao.txt`.

## Como Executar

### 1. Clonar o Repositório
```sh
 git clone <URL_DO_REPOSITORIO>
 cd <NOME_DO_REPOSITORIO>
```

### 2. Instalar Dependências
```sh
go mod init cotacao_app
go mod tidy
```

### 3. Rodar o Servidor
```sh
go run server.go
```
O servidor estará rodando na porta `8080`.

### 4. Rodar o Cliente
Em outro terminal, execute:
```sh
go run client.go
```
Isso irá exibir a cotação do dólar e salvar o valor no arquivo `cotacao.txt`.

## Notas Técnicas
- O **server.go** utiliza **contextos com timeout** para evitar requisições bloqueadas.
- O **banco SQLite** é criado automaticamente se não existir.
- O **client.go** tem um timeout de 300ms para garantir performance na requisição.

---

Desenvolvido para a pós-graduação em GoLang 🚀

