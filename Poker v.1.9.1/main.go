package main

import (
	"fmt"
)

// Naipe representa os possíveis naipes de uma carta no baralho.
type Naipe string

// Constantes dos naipes, representados por seus respectivos símbolos Unicode.
const (
	Paus    Naipe = "\u2663" // ♣
	Ouros   Naipe = "\u2666" // ♦
	Copas   Naipe = "\u2665" // ♥
	Espadas Naipe = "\u2660" // ♠
)

// Carta define a estrutura de uma carta no baralho, composta por um valor e um naipe.
type Carta struct {
	Valor string
	Naipe Naipe
}

// Jogador representa uma entidade jogador, que pode ser um humano ou CPU.
type Jogador struct {
	Nome     string  // Nome do jogador
	Tipo     bool    // Se true, o jogador é humano, caso contrário é CPU
	Mao      []Carta // Cartas atualmente na mão do jogador
	Cartas   []Carta // Cartas totais do jogador (mão + comunitárias)
	Dinheiro int     // Dinheiro atual do jogador
	Aposta   int     // Valor da aposta atual feita pelo jogador
	Ativo    bool    // Status atual do jogador na rodada (ativo ou não)
	NumSorte int     // Usado para determinar decisões para a CPU
	levelMao float64 // Nível da mão, usado para avaliar força da mão
	tipoMao  string  // Descrição textual do tipo da mão (e.g., "Royal Flush")
}

// Mesa representa o estado atual da mesa de poker.
type Mesa struct {
	CartasComunitarias []Carta // Cartas que são reveladas comum a todos os jogadores
	Pote               int     // Quantidade total de dinheiro no pote
}

// Mao é um tipo enum para representar a força da mão do jogador.
type Mao int

// Constantes que representam diferentes forças de mão.
const (
	ForcaDaMaoMuitoFraca Mao = 1
	ForcaDaMaoFraca      Mao = 2
	ForcaDaMaoRazoavel   Mao = 3
	ForcaDaMaoBoa        Mao = 4
	ForcaDaMaoMuitoBoa   Mao = 5
)

// AvaliacaoMao armazena informações sobre a avaliação da mão de um jogador.
type AvaliacaoMao struct {
	Tipo     Mao     // Força da mão
	Cartas   []Carta // Cartas que contribuem para a força da mão
	ValorPar string  // Usado quando a mão é um par
}

// jogadores é a lista de jogadores para a partida.
var jogadores = inicializarJogadores()

// main é a função de entrada do programa.
func main() {
	for {
		jogarPartida()         // Inicia uma partida
		confirmarContinuacao() // Pergunta ao jogador se ele quer jogar novamente
	}
}

// jogarPartida controla o fluxo de uma única partida de poker.
func jogarPartida() {
	clearTerminal()                      // Limpa a tela do terminal para uma nova partida
	baralho := criarEEmbaralharBaralho() // Prepara um novo baralho de cartas
	mesa := &Mesa{}                      // Reseta o estado da mesa
	jogarPoker(jogadores, mesa, baralho) // Começa a partida
	fmt.Printf("\nFim da partida!\n\n")  // Informa o fim da partida ao jogador
}
