package main

import (
	"log"
	"math"
	"sort"
)

type RankingMao int

const (
	Invalido RankingMao = iota
	HighCard
	OnePair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	RoyalFlush
)

// Representação em string dos rankings.
var rankingString = map[RankingMao]string{
	Invalido:      "Invalido",
	HighCard:      "HighCard",
	OnePair:       "OnePair",
	TwoPair:       "TwoPair",
	ThreeOfAKind:  "ThreeOfAKind",
	Straight:      "Straight",
	Flush:         "Flush",
	FullHouse:     "FullHouse",
	FourOfAKind:   "FourOfAKind",
	StraightFlush: "StraightFlush",
	RoyalFlush:    "RoyalFlush",
}

// valorCartas mapeia a representação em string do valor de uma carta para seu valor numérico.
var valorCartas = map[string]int{
	"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10,
	"J": 11, "Q": 12, "K": 13, "A": 14,
}

// quantificarMao calcula um valor único para uma mão baseado em seu ranking e seus valores.
func quantificarMao(ranking RankingMao, valores []int) float64 {
	quantificacao := float64(ranking)
	for i, valor := range valores {
		quantificacao += float64(valor) / math.Pow(10, float64((i+1)*2))
	}
	return quantificacao
}

// avaliarMao avalia a melhor mão que um jogador pode formar com suas cartas e cartas comunitárias.
func avaliarMao(jogadores []Jogador, comunitarias []Carta) {
	for i := range jogadores {
		if jogadores[i].Ativo {
			var melhorQuantificacao float64
			combinacoes := combinarCartas(jogadores[i].Mao, comunitarias)
			for _, combinacao := range combinacoes {
				quantificacao := rankingMao(combinacao)
				if quantificacao > melhorQuantificacao {
					melhorQuantificacao = quantificacao
				}
			}
			jogadores[i].levelMao = melhorQuantificacao
			jogadores[i].tipoMao = rankingString[RankingMao(int(melhorQuantificacao))]
		}
	}
}

// combinarCartas gera todas as combinações possíveis de 5 cartas a partir de um conjunto de cartas.
func combinarCartas(mao, comunitarias []Carta) [][]Carta {
	var combinacoes [][]Carta
	todasCartas := append(mao, comunitarias...)
	for i := 0; i < (1 << len(todasCartas)); i++ {
		var combinacao []Carta
		for j := 0; j < len(todasCartas); j++ {
			if (i & (1 << j)) != 0 {
				combinacao = append(combinacao, todasCartas[j])
			}
		}
		if len(combinacao) == 5 {
			combinacoes = append(combinacoes, combinacao)
		}
	}
	return combinacoes
}

// rankingMao calcula o ranking de uma mão de poker.
func rankingMao(mao []Carta) float64 {
	ranking := avaliarForcaDaMao(mao)
	valores := make([]int, len(mao))
	for i, carta := range mao {
		valores[i] = valorCartas[carta.Valor]
	}
	return quantificarMao(ranking, valores)
}

// avaliarForcaDaMao determina o ranking de uma mão de poker.
func avaliarForcaDaMao(mao []Carta) RankingMao {
	contaNaipes := make(map[Naipe]int)
	contaValores := make(map[int]int)
	var valores []int
	for _, carta := range mao {
		contaNaipes[carta.Naipe]++
		valor := valorCartas[carta.Valor]
		if valor == 0 {
			log.Fatalf("Valor de carta inválido: %v", carta.Valor)
		}
		contaValores[valor]++
		if contaValores[valor] == 1 {
			valores = append(valores, valor)
		}
	}

	sort.Ints(valores)
	if isFlush(contaNaipes) {
		if isStraight(valores) {
			if valores[len(valores)-1] == 14 {
				return RoyalFlush
			}
			return StraightFlush
		}
		return Flush
	}
	if isFourOfAKind(contaValores) {
		return FourOfAKind
	}
	if isFullHouse(contaValores) {
		return FullHouse
	}
	if isThreeOfAKind(contaValores) {
		return ThreeOfAKind
	}
	if isTwoPair(contaValores) {
		return TwoPair
	}
	if isOnePair(contaValores) {
		return OnePair
	}
	if isStraight(valores) {
		return Straight
	}
	if len(valores) > 0 {
		return HighCard
	}
	return Invalido
}

// Funções auxiliares para verificar condições específicas de classificações de mãos de poker.
func isFlush(contaNaipes map[Naipe]int) bool {
	for _, conta := range contaNaipes {
		if conta >= 5 {
			return true
		}
	}
	return false
}

func isStraight(valores []int) bool {
	for i := 0; i <= len(valores)-5; i++ {
		if valores[i] == 2 && valores[i+3] == 5 && contains(valores, 14) {
			return true
		}
		if valores[i]+4 == valores[i+4] {
			return true
		}
	}
	return false
}

func isFourOfAKind(contaValores map[int]int) bool {
	for _, conta := range contaValores {
		if conta == 4 {
			return true
		}
	}
	return false
}

func isFullHouse(contaValores map[int]int) bool {
	var hasThree, hasTwo bool
	for _, conta := range contaValores {
		if conta == 3 {
			hasThree = true
		} else if conta == 2 {
			hasTwo = true
		}
	}
	return hasThree && hasTwo
}

func isThreeOfAKind(contaValores map[int]int) bool {
	for _, conta := range contaValores {
		if conta == 3 {
			return true
		}
	}
	return false
}

func isTwoPair(contaValores map[int]int) bool {
	count := 0
	for _, conta := range contaValores {
		if conta == 2 {
			count++
		}
	}
	return count >= 2
}

func isOnePair(contaValores map[int]int) bool {
	for _, conta := range contaValores {
		if conta == 2 {
			return true
		}
	}
	return false
}

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
