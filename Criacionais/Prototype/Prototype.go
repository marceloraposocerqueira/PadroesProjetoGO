package main

import "fmt"

// 1. Interface Prototype
type Documento interface {
	Clone() Documento
	Exibir()
}

// 2. Concreto Prototype
type Relatorio struct {
	Titulo   string
	Conteudo string
}

// Implementação do Clone (Cópia Superficial)
func (r *Relatorio) Clone() Documento {
	// Cria uma cópia do valor, não da referência
	return &Relatorio{
		Titulo:   r.Titulo + " (Cópia)",
		Conteudo: r.Conteudo,
	}
}

func (r *Relatorio) Exibir() {
	fmt.Printf("Relatório: %s | Conteúdo: %s\n", r.Titulo, r.Conteudo)
}

func main() {
	// Criar protótipo original
	relatorioOriginal := &Relatorio{
		Titulo:   "Relatório Anual",
		Conteudo: "Dados financeiros...",
	}

	// Clonar
	relatorioClone := relatorioOriginal.Clone()

	// Alterar o original não afeta o clone (se for cópia profunda)
	relatorioOriginal.Titulo = "Relatório Anual 2026"

	relatorioOriginal.Exibir()
	relatorioClone.Exibir()
}
