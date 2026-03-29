package main

import "fmt"

// ProdutoA é a interface para o produto do tipo A
type ProdutoA interface {
	UsarA()
}

// ProdutoB é a interface para o produto do tipo B
type ProdutoB interface {
	UsarB()
}

// ConcreteProdutoA1 é uma implementação concreta de ProdutoA
type ConcreteProdutoA1 struct{}

func (p *ConcreteProdutoA1) UsarA() {
	fmt.Println("Usando Produto A1")
}

// ConcreteProdutoA2 é outra implementação concreta de ProdutoA
type ConcreteProdutoA2 struct{}

func (p *ConcreteProdutoA2) UsarA() {
	fmt.Println("Usando Produto A2")
}

// ConcreteProdutoB1 é uma implementação concreta de ProdutoB
type ConcreteProdutoB1 struct{}

func (p *ConcreteProdutoB1) UsarB() {
	fmt.Println("Usando Produto B1")
}

// ConcreteProdutoB2 é outra implementação concreta de ProdutoB
type ConcreteProdutoB2 struct{}

func (p *ConcreteProdutoB2) UsarB() {
	fmt.Println("Usando Produto B2")
}

// AbstractFactory define os métodos para criar produtos
type AbstractFactory interface {
	CriarProdutoA() ProdutoA
	CriarProdutoB() ProdutoB
}

// ConcreteFactory1 cria produtos do tipo 1
type ConcreteFactory1 struct{}

func (f *ConcreteFactory1) CriarProdutoA() ProdutoA {
	return &ConcreteProdutoA1{}
}

func (f *ConcreteFactory1) CriarProdutoB() ProdutoB {
	return &ConcreteProdutoB1{}
}

// ConcreteFactory2 cria produtos do tipo 2
type ConcreteFactory2 struct{}

func (f *ConcreteFactory2) CriarProdutoA() ProdutoA {
	return &ConcreteProdutoA2{}
}

func (f *ConcreteFactory2) CriarProdutoB() ProdutoB {
	return &ConcreteProdutoB2{}
}

// Função cliente
func main() {
	var factory AbstractFactory

	// Usando a primeira fábrica concreta
	factory = &ConcreteFactory1{}
	produtoA := factory.CriarProdutoA()
	produtoB := factory.CriarProdutoB()
	produtoA.UsarA()
	produtoB.UsarB()

	// Usando a segunda fábrica concreta
	factory = &ConcreteFactory2{}
	produtoA = factory.CriarProdutoA()
	produtoB = factory.CriarProdutoB()
	produtoA.UsarA()
	produtoB.UsarB()
}
