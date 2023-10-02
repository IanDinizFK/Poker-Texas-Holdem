// Package main é o pacote principal que contém o código do seu programa.
package main

import (
	"log"       // Usado para logar mensagens de erro fatal
	"math/rand" // Para gerar números aleatórios
	"sort"      // Para ordenar slices
	"time"      // Para inicializar o gerador de números aleatórios com o tempo atual
)

// nivelMesa é uma variável global para armazenar o nível da mesa.
var nivelMesa int

// Valor é um tipo para representar o valor da carta.
type Valor string

// Inicialize o gerador de números aleatórios.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// decidirApostaBot decide a aposta do bot com base em várias condições.
// Leva em consideração o jogador, cartas da mesa, aposta mínima, rodada e tamanho do pote.
func decidirApostaBot(jogador *Jogador, cartasMesa []Carta, apostaMinima int, rodada int, tamanhoPote int) int {
	if rodada == 1 {
		return apostaMinima
	}

	// Obter o nível da mão do bot a partir das informações do jogador
	levelMao := jogador.levelMao

	// Ajustar a aposta com base no nível da mão (escala de 1 a 10)
	var aposta int
	switch {
	case levelMao >= 10:
		// Mão extremamente forte
		aposta = apostaMinima * 4
		if rand.Float64() < 0.2 {
			aposta = apostaMinima * 8 // 20% de chance de aumentar a aposta ainda mais
		}
	case levelMao >= 7:
		// Mão muito forte
		aposta = apostaMinima * 2
		if rand.Float64() < 0.3 {
			aposta = apostaMinima * 4 // 30% de chance de aumentar a aposta
		}
	case levelMao >= 5:
		// Mão forte
		aposta = apostaMinima
		if rand.Float64() < 0.4 {
			aposta = apostaMinima * 2 // 40% de chance de aumentar a aposta
		}
	case levelMao >= 2:
		// Mão média
		aposta = apostaMinima
		if rand.Float64() < 0.2 {
			aposta = -1 // 20% de chance de desistir
		}
	default:
		// Mão fraca a muito fraca, desistir
		aposta = -1
	}

	// Levar em consideração o tamanho do pote
	if tamanhoPote > 2*apostaMinima {
		if aposta > 0 {
			// Aumentar a aposta com uma probabilidade de 30%
			if rand.Float64() < 0.3 {
				aposta = int(float64(aposta) * 1.5)
			}
		} else {
			// Igualar a aposta com uma probabilidade de 40%
			if rand.Float64() < 0.4 {
				aposta = apostaMinima
			}
		}
	}

	// Fazer "All In" se o dinheiro for menor do que a aposta mínima
	if jogador.Dinheiro < apostaMinima {
		aposta = jogador.Dinheiro
	}

	// Se for a última rodada e o bot estiver confiante, seja mais agressivo
	if rodada == 4 && aposta > 0 && aposta < jogador.Dinheiro {
		if rand.Float64() < 0.6 {
			aposta = int(float64(aposta) * 1.5)
		}
	}

	// Garantir que o bot não aposte mais do que tem
	if aposta > jogador.Dinheiro {
		aposta = jogador.Dinheiro
	}

	return aposta
}

// Função possibilidadeSequencia verifica a possibilidade de uma sequência nas cartas.
// Retorna o número de cartas faltando para uma sequência e se é possível.
func possibilidadeSequencia(cartasMao, cartasMesa []Carta) (int, bool) {
	todasCartas := append(cartasMao, cartasMesa...)
	var valores []int
	for _, carta := range todasCartas {
		valor, existe := valorCartas[carta.Valor]
		if !existe {
			log.Fatalf("Valor de carta inválido: %v", carta.Valor)
		}
		valores = append(valores, valor)
	}
	sort.Ints(valores)

	for i := 0; i < len(valores); i++ {
		sequencia := 1
		for j := i + 1; j < len(valores) && valores[j] <= valores[i]+4; j++ {
			if valores[j] > valores[j-1] {
				sequencia++
			}
		}
		if sequencia >= 4 {
			// Retornando o número de cartas faltando para uma sequência
			return 5 - sequencia, true
		}
	}
	// Retornando -1 para indicar que não há possibilidade de sequência.
	return -1, false
}
