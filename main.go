package main

import (
	"fmt"
	"log"
	"os"
	"github.com/calvincabral/calculet/fetcher"
	"github.com/calvincabral/calculet/instruction"
	"github.com/calvincabral/calculet/filter"
)

func main() {
	// Verifica se a URL do endpoint foi fornecida
	if len(os.Args) < 2 {
		log.Fatal("Por favor, forneça o URL do endpoint como argumento.")
	}

	endpoint := os.Args[1] // A URL do endpoint é fornecida como argumento

	// Autenticação opcional (se fornecida)
	var authType, authValue string
	if len(os.Args) > 2 {
		authType = os.Args[2] // Tipo de autenticação (pode ser 'Bearer', 'Basic', etc.)
		if len(os.Args) > 3 {
			authValue = os.Args[3] // O valor de autenticação (token ou credenciais)
		}
	}

	// Definindo os headers dinamicamente (só inclui autenticação se fornecido)
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	// Se autenticação for fornecida, adiciona ao header
	if authType != "" && authValue != "" {
		headers["Authorization"] = fmt.Sprintf("%s %s", authType, authValue)
	}

	// Criando a instrução com os parâmetros fornecidos
	inst := instruction.NewInstruction(
		endpoint,                // URL fornecida pelo usuário
		"GET",                   // Método HTTP
		headers,                 // Headers configurados dinamicamente
		[]string{"id", "title"}, // Campos a serem extraídos
		nil,                     // Corpo da requisição (se necessário)
	)

	// Fazendo a requisição ao endpoint
	result, err := fetcher.DoRequest(inst.Endpoint, inst.Method, inst.Headers, inst.RequestBody)
	if err != nil {
		log.Fatalf("Erro ao fazer requisição: %s", err)
	}

	// Exibindo os dados recebidos
	fmt.Println("Dados recebidos:", result)

	// Se houver um filtro a ser aplicado, ele será passado como argumento adicional
	if len(os.Args) > 4 {
		filter := os.Args[4] // Filtro passado como argumento
		filteredResult := filter.ApplyFilter(result, filter)
		fmt.Println("Dados filtrados:", filteredResult)
	}
}
