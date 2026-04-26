package main

import "fmt"

// 1. Interface da Implementação
type Device interface {
	IsEnabled() bool
	Enable()
	Disable()
	SetVolume(percent int)
}

// 2. Implementações Concretas
type Tv struct{}

func (t *Tv) IsEnabled() bool       { return true }
func (t *Tv) Enable()               { fmt.Println("TV ligada") }
func (t *Tv) Disable()              { fmt.Println("TV desligada") }
func (t *Tv) SetVolume(percent int) { fmt.Println("Volume TV:", percent) }

type Radio struct{}

func (r *Radio) IsEnabled() bool       { return false }
func (r *Radio) Enable()               { fmt.Println("Rádio ligado") }
func (r *Radio) Disable()              { fmt.Println("Rádio desligado") }
func (r *Radio) SetVolume(percent int) { fmt.Println("Volume Rádio:", percent) }

// 3. Abstração (O "O que fazer")
type Remote struct {
	device Device // A PONTE
}

func (r *Remote) TogglePower() {
	if r.device.IsEnabled() {
		r.device.Disable()
	} else {
		r.device.Enable()
	}
}

// 4. Abstração Refinada
// Em Go, usamos "Embedding" para simular a extensão de funcionalidades
type AdvancedRemote struct {
	Remote
}

func (ar *AdvancedRemote) Mute() {
	ar.device.SetVolume(0)
}

// Uso
func main() {
	// Exemplo com TV e Controle Básico
	tv := &Tv{}
	basicRemote := &Remote{device: tv}
	basicRemote.TogglePower() // TV ligada

	// Exemplo com Rádio e Controle Avançado
	radio := &Radio{}
	advRemote := &AdvancedRemote{
		Remote: Remote{device: radio},
	}
	advRemote.Mute() // Volume Rádio: 0
}
