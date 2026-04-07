package main

import (
	"fmt"
	"sync"
)

// Database define a estrutura da nossa conexão.
type Database struct {
	connectionString string
}

// instance armazena a instância única (equivalente ao campo estático privado).
var instance *Database

// once é usado para garantir que a inicialização ocorra apenas uma vez.
var once sync.Once

// GetInstance provê o ponto de acesso global à instância do Singleton.
func GetInstance() *Database {
	if instance != nil {
		fmt.Println("Instância já criada, retornando a instância existente...")
	} else {
		once.Do(func() {
			// Este bloco é executado apenas uma única vez,
			// mesmo se chamado por múltiplas goroutines simultaneamente.
			fmt.Println("Criando instância única do Banco de Dados...")
			instance = &Database{
				connectionString: "postgres://user:password@localhost:5432/mydb",
			}
		})
	}
	return instance
}

// Query representa a lógica de negócio (método da instância).
func (d *Database) Query(sql string) {
	fmt.Printf("Executando query: %s\n", sql)
}

// Application (main)
func main() {
	// Primeira chamada: cria a instância.
	foo := GetInstance()
	foo.Query("SELECT * FROM users")

	// Segunda chamada: retorna a instância já existente.
	bar := GetInstance()
	bar.Query("SELECT * FROM products")

	// Verificação de que são a mesma instância (mesmo endereço de memória).
	if foo == bar {
		fmt.Println("Sucesso: Ambas as variáveis contêm a mesma instância.")
	}
}
