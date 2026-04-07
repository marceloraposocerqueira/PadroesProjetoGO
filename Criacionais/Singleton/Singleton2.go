package main

import (
	"fmt"
	"sync"
)

// 1. Estrutura privada
type single struct {
}

// 2. Variável da instância única e objeto sync.Once
var singleInstance *single
var once sync.Once

// 3. Função pública para obter a instância
func getInstance() *single {
	if singleInstance != nil {
		fmt.Println("Instância única já criada.")
	} else {
		// once.Do garante que o código dentro seja executado apenas uma vez
		once.Do(func() {
			fmt.Println("Criando instância única agora.")
			singleInstance = &single{}
		})
	}
	return singleInstance
}

func main() {
	// Simulando múltiplas goroutines tentando obter a instância
	for i := 1; i <= 5; i++ {
		getInstance()
	}
}
