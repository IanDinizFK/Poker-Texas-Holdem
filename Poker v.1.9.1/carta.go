// Package main é o pacote principal onde este código está localizado.
package main

import (
	"math/rand" // Utilizado para embaralhar o baralho
	"time"      // Utilizado para inicializar o gerador de números aleatórios
)

// valores é um mapa que relaciona as strings representando valores de cartas a seus respectivos valores inteiros.
var valores = map[string]int{
	"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10,
	"J": 11, "Q": 12, "K": 13, "A": 14,
}

// valorCarta recebe uma Carta e retorna seu valor inteiro correspondente.
func valorCarta(carta Carta) int {
	return valores[carta.Valor] // Retorna o valor inteiro da carta fornecida
}

// criarBaralho cria um novo baralho de cartas.
func criarBaralho() []Carta {
	var baralho []Carta // Inicializa um novo baralho vazio
	for _, valor := range append([]string{"2", "3", "4", "5", "6", "7", "8", "9", "10"}, "A", "J", "Q", "K") {
		for _, naipe := range []Naipe{Paus, Ouros, Copas, Espadas} {
			carta := Carta{Valor: valor, Naipe: naipe}
			baralho = append(baralho, carta)
		}
	}
	return baralho // Retorna o baralho criado
}

// embaralharBaralho recebe um baralho e o embaralha.
func embaralharBaralho(baralho []Carta) {
	// Cria um novo gerador de números aleatórios e o usa para embaralhar o baralho
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range baralho {
		j := r.Intn(len(baralho))
		baralho[i], baralho[j] = baralho[j], baralho[i]
	}
}

// criarEEmbaralharBaralho cria um novo baralho de cartas e o embaralha antes de retorná-lo.
func criarEEmbaralharBaralho() []Carta {
	baralho := criarBaralho()  // Cria um novo baralho
	embaralharBaralho(baralho) // Embaralha o baralho criado
	return baralho             // Retorna o baralho embaralhado
}

// distribuirCartas distribui um número definido de cartas para cada jogador ativo
// e retorna o baralho atualizado.
func distribuirCartas(jogadores []Jogador, baralho []Carta, quantidade int) []Carta {
	for i := range jogadores {
		jogador := &jogadores[i]
		if jogador.Ativo {
			jogador.Mao = baralho[:quantidade]
			baralho = baralho[quantidade:]
		}
	}
	return baralho // Retorna o baralho atualizado após a distribuição de cartas
}
