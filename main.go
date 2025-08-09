// main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// Estrutura para armazenar as URLs. Usaremos dois mapas para facilitar as buscas.
// urlStore: guarda a associação de código_curto -> url_longa
// reverseUrlStore: guarda a associação de url_longa -> código_curto (para evitar duplicatas)
var urlStore = make(map[string]string)
var reverseUrlStore = make(map[string]string)

const shortCodeLength = 6
const allowedChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// init() é uma função especial em Go que roda antes da main()
func init() {
	// Inicializa o gerador de números aleatórios com uma semente baseada no tempo atual
	rand.Seed(time.Now().UnixNano())
}

// generateShortCode cria uma string aleatória para ser nosso código curto
func generateShortCode() string {
	b := make([]byte, shortCodeLength)
	for i := range b {
		b[i] = allowedChars[rand.Intn(len(allowedChars))]
	}
	return string(b)
}

// --- NOSSAS FUNÇÕES DE ROTA (HANDLERS) VIRÃO AQUI ---

// Função principal que configura o servidor e as rotas
func main() {
	fmt.Println("Servidor Encurtador de URL rodando na porta 8080")

	// Ativando as rotas para usar nossos handlers
	http.HandleFunc("/encurtar", shortenHandler)
	http.HandleFunc("/", redirectHandler) // O "/" vai pegar todas as rotas que não sejam "/encurtar"

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Estruturas para lidar com o JSON da requisição e da resposta
type URLRequest struct {
	URL string `json:"url"`
}

type URLResponse struct {
	ShortURL string `json:"short_url"`
}

// shortenHandler lida com as requisições para encurtar uma URL
func shortenHandler(w http.ResponseWriter, r *http.Request) {
	// Garante que estamos recebendo um método POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var req URLRequest
	// Decodifica o JSON do corpo da requisição para a nossa struct
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Corpo da requisição inválido", http.StatusBadRequest)
		return
	}

	// Verifica se a URL foi enviada
	if req.URL == "" {
		http.Error(w, "URL não pode ser vazia", http.StatusBadRequest)
		return
	}

	// Verifica se a URL já foi encurtada antes para evitar duplicatas
	if shortCode, ok := reverseUrlStore[req.URL]; ok {
		// Se já existe, retorna a URL curta existente
		response := URLResponse{ShortURL: "http://localhost:8080/" + shortCode}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// Gera um novo código curto e único
	var shortCode string
	for {
		shortCode = generateShortCode()
		// Garante que o código gerado já não está em uso (muito improvável, mas é uma boa prática)
		if _, ok := urlStore[shortCode]; !ok {
			break
		}
	}

	// Armazena a associação nos nossos dois mapas
	urlStore[shortCode] = req.URL
	reverseUrlStore[req.URL] = shortCode

	// Constrói a URL de resposta completa
	shortURL := "http://localhost:8080/" + shortCode
	response := URLResponse{ShortURL: shortURL}

	// Define o cabeçalho como JSON e envia a resposta com o status "Created" (201)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// redirectHandler lida com as requisições para acessar uma URL curta
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	// Pega o caminho da URL, que é o nosso código curto (ex: "/aB1cD2")
	// [1:] remove a primeira barra "/"
	shortCode := r.URL.Path[1:]

	// Procura o código curto no nosso mapa
	longURL, ok := urlStore[shortCode]
	if !ok {
		// Se não encontrar, retorna um erro 404 Not Found
		http.NotFound(w, r)
		return
	}

	// Se encontrar, redireciona o navegador para a URL original
	// http.StatusFound é o código para um redirecionamento temporário (302)
	http.Redirect(w, r, longURL, http.StatusFound)
}
