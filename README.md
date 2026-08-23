# Padrões de Projeto em Go

Coleção de exemplos práticos dos padrões de projeto GoF implementados em Go, organizada nas categorias criacionais e estruturais.

[![Go](https://img.shields.io/badge/Go-1.22.2-00ADD8?logo=go)](https://go.dev/)
[![Licença](https://img.shields.io/badge/licença-MIT-green.svg)](#licença)
[![Build](https://img.shields.io/badge/build-exemplos%20compiláveis-brightgreen.svg)](#testes)

## Recursos

- Exemplos pequenos e independentes, fáceis de ler e executar.
- Padrões criacionais: Abstract Factory, Builder, Factory Method, Prototype e Singleton.
- Padrões estruturais: Adapter e Bridge.
- Código focado na aplicação dos conceitos, sem dependências externas.
- Organização por categoria para facilitar o estudo e a comparação entre padrões.

## Estrutura do projeto

```text
Criacionais/
├── AbstractFactory/
│   └── animation/
├── Builder/
├── FactoryMethod/
├── Prototype/
└── Singleton/

Estruturais/
├── Adapter/
└── Bridge/
```

Os arquivos dentro de `AbstractFactory/animation` formam um único exemplo. Nos demais diretórios, os arquivos com sufixo `2` são variações independentes do padrão correspondente.

## Pré-requisitos

- [Go 1.22.2 ou superior](https://go.dev/dl/).
- Git, caso queira clonar o repositório.

Confira a instalação com:

```bash
go version
```

## Como instalar e rodar

Clone o repositório e entre no diretório do projeto:

```bash
git clone <URL_DO_REPOSITORIO>
cd PadroesProjetoGO
go mod download
```

Como este repositório é uma coleção de exemplos, não existe um único comando de inicialização na raiz. Execute o padrão desejado informando o caminho do pacote:

```bash
go run ./Criacionais/AbstractFactory/animation
go run ./Criacionais/Builder/Builder.go
go run ./Criacionais/FactoryMethod/FactoryMethod.go
go run ./Criacionais/Prototype/Prototype.go
go run ./Criacionais/Singleton/Singleton.go
go run ./Estruturais/Adapter/Adapter.go
go run ./Estruturais/Bridge/Bridge.go
```

Os arquivos com sufixo `2` são exemplos alternativos e devem ser executados individualmente:

```bash
go run ./Criacionais/AbstractFactory/AbstractFactory2.go
go run ./Criacionais/Prototype/Prototype2.go
go run ./Criacionais/Singleton/Singleton2.go
go run ./Estruturais/Adapter/Adapter2.go
go run ./Estruturais/Bridge/Bridge2.go
```

## Como usar

### Builder

O exemplo de Builder permite construir uma `Person` em etapas, usando uma API encadeada:

```go
person := NewPersonBuilder().
	FirstName("Ana").
	LastName("Silva").
	Age(30).
	Email("ana@example.com").
	Build()
```

Para observar esse exemplo em funcionamento:

```bash
go run ./Criacionais/Builder/Builder.go
```

### Abstract Factory

O exemplo de Abstract Factory seleciona uma fábrica de personagem para `manga`, `anime` ou `cartoon`. O valor padrão está definido no código como `manga`:

```bash
go run ./Criacionais/AbstractFactory/animation
```

Para estudar outro tipo, altere `animationType` em `AnimationApp.go` para `anime` ou `cartoon` e execute o comando novamente.

## Testes

Execute a verificação padrão de todos os pacotes:

```bash
go test ./Criacionais/AbstractFactory/animation/...
```

No momento, o repositório contém exemplos executáveis, mas ainda não possui arquivos de teste automatizados. O comando acima verifica o pacote de Abstract Factory que reúne múltiplos arquivos. O comando abrangente `go test ./...` não pode ser usado enquanto os arquivos base e as variações `2` permanecerem no mesmo pacote, pois eles possuem declarações duplicadas.

## Contribuição

1. Faça um fork do projeto.
2. Crie uma branch para sua alteração: `git checkout -b feat/novo-padrao`.
3. Implemente ou melhore um exemplo mantendo o foco didático.
4. Formate o código com `gofmt` e execute `go test ./...`.
5. Abra um pull request descrevendo a mudança e o padrão abordado.

Sugestões, correções e novos padrões são bem-vindos.

## Licença

Este projeto é destinado a ser distribuído sob a licença MIT. O arquivo `LICENSE` ainda deve ser adicionado ao repositório com os termos completos.
