package main

import (
	"fmt"
	"log"
)

// --- Interfaces (Produtos Abstratos) ---

type Button interface {
	Paint()
}

type Checkbox interface {
	Paint()
}

// --- Interfaces (Fábrica Abstrata) ---

type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

// --- Implementações Windows (Variante Concreta 1) ---

type WinButton struct{}

func (w *WinButton) Paint() {
	fmt.Println("Renderizando um botão no estilo Windows.")
}

type WinCheckbox struct{}

func (w *WinCheckbox) Paint() {
	fmt.Println("Renderizando uma caixa de seleção no estilo Windows.")
}

type WinFactory struct{}

func (w *WinFactory) CreateButton() Button {
	return &WinButton{}
}

func (w *WinFactory) CreateCheckbox() Checkbox {
	return &WinCheckbox{}
}

// --- Implementações Mac (Variante Concreta 2) ---

type MacButton struct{}

func (m *MacButton) Paint() {
	fmt.Println("Renderizando um botão no estilo macOS.")
}

type MacCheckbox struct{}

func (m *MacCheckbox) Paint() {
	fmt.Println("Renderizando uma caixa de seleção no estilo macOS.")
}

type MacFactory struct{}

func (m *MacFactory) CreateButton() Button {
	return &MacButton{}
}

func (m *MacFactory) CreateCheckbox() Checkbox {
	return &MacCheckbox{}
}

// --- Implementações Linux (Variante Concreta 3) ---

type LinuxButton struct{}

func (m *LinuxButton) Paint() {
	fmt.Println("Renderizando um botão no estilo linux.")
}

type LinuxCheckbox struct{}

func (m *LinuxCheckbox) Paint() {
	fmt.Println("Renderizando uma caixa de seleção no estilo linux.")
}

type LinuxFactory struct{}

func (m *LinuxFactory) CreateButton() Button {
	return &LinuxButton{}
}

func (m *LinuxFactory) CreateCheckbox() Checkbox {
	return &LinuxCheckbox{}
}

// --- Código Cliente ---

type Application struct {
	factory  GUIFactory
	button   Button
	checkbox Checkbox
}

func NewApplication(f GUIFactory) *Application {
	return &Application{
		factory: f,
	}
}

func (a *Application) CreateUI() {
	a.button = a.factory.CreateButton()
	a.checkbox = a.factory.CreateCheckbox()
}

func (a *Application) Paint() {
	a.button.Paint()
	a.checkbox.Paint()
}

// --- Configuração e Main ---

func main() {
	// Simulando a leitura de uma configuração
	osType := "Linux" // Poderia vir de um arquivo de config

	var factory GUIFactory

	switch osType {
	case "Windows":
		factory = &WinFactory{}
	case "Mac":
		factory = &MacFactory{}
	case "Linux":
		factory = &LinuxFactory{}
	default:
		log.Fatal("Erro! Sistema operacional desconhecido.")
	}

	app := NewApplication(factory)
	app.CreateUI()
	app.Paint()
}
