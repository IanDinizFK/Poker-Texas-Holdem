package main

import (
	"fmt"
)

// distribuirCartasMesa distribui uma quantidade especificada de cartas para as cartas comunitárias da mesa
// e retorna o baralho restante.
// mesa: A mesa atual de jogo.
// baralho: O conjunto de cartas disponíveis para distribuição.
// quantidade: O número de cartas a serem distribuídas para a mesa.
func distribuirCartasMesa(mesa *Mesa, baralho []Carta, quantidade int) []Carta {
	mesa.CartasComunitarias = append(mesa.CartasComunitarias, baralho[:quantidade]...)
	return baralho[quantidade:]
}

// printarInfoMesa exibe as informações da mesa no console, incluindo cartas comunitárias e o valor do pote.
// mesa: A mesa atual de jogo.
func printarInfoMesa(mesa *Mesa) {
	fmt.Print("CARTAS COMUNITÁRIAS\n")
	for _, carta := range mesa.CartasComunitarias {
		fmt.Printf("[%v%s]", carta.Valor, carta.Naipe)
	}
	fmt.Printf("\nPOTE: %d Fichas", mesa.Pote)
	fmt.Print("\n\n")
}
