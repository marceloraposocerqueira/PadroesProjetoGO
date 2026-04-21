package main

import (
	"fmt"
	"math"
)

// --- Interfaces e Estruturas Originais ---

// RoundHole representa o Buraco Redondo
type RoundHole struct {
	radius float64
}

func (h *RoundHole) GetRadius() float64 {
	return h.radius
}

// Fits verifica se um "Pino Redondo" (via interface) cabe no buraco
func (h *RoundHole) Fits(p RoundPegInterface) bool {
	return h.GetRadius() >= p.GetRadius()
}

// RoundPegInterface define o comportamento esperado para algo ser um "Pino Redondo"
type RoundPegInterface interface {
	GetRadius() float64
}

// RoundPeg é a implementação padrão de um pino redondo
type RoundPeg struct {
	radius float64
}

func (p *RoundPeg) GetRadius() float64 {
	return p.radius
}

// --- Classe Incompatível ---

// SquarePeg é o Pino Quadrado (incompatível com RoundHole)
type SquarePeg struct {
	width float64
}

func (s *SquarePeg) GetWidth() float64 {
	return s.width
}

// --- O Adaptador ---

// SquarePegAdapter adapta SquarePeg para satisfazer RoundPegInterface
type SquarePegAdapter struct {
	peg *SquarePeg
}

// GetRadius realiza o cálculo matemático para converter largura em raio circular equivalente
func (a *SquarePegAdapter) GetRadius() float64 {
	// Raio de um círculo que circunscreve um quadrado: (largura * √2) / 2
	return a.peg.GetWidth() * math.Sqrt(2) / 2
}

// --- Código Cliente ---

func main() {
	hole := &RoundHole{radius: 5}
	rpeg := &RoundPeg{radius: 5}

	fmt.Printf("Pino redondo de raio 5 cabe no buraco de raio 5? %v\n", hole.Fits(rpeg))

	smallSqPeg := &SquarePeg{width: 5}
	largeSqPeg := &SquarePeg{width: 10}

	// hole.Fits(smallSqPeg) // Isso não compilaria em Go, tipos diferentes!

	// Usando o Adaptador
	smallSqPegAdapter := &SquarePegAdapter{peg: smallSqPeg}
	largeSqPegAdapter := &SquarePegAdapter{peg: largeSqPeg}

	fmt.Printf("Pino quadrado de largura 5 cabe? %v\n", hole.Fits(smallSqPegAdapter))  // True
	fmt.Printf("Pino quadrado de largura 10 cabe? %v\n", hole.Fits(largeSqPegAdapter)) // False
}
