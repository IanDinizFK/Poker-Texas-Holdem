// Package main é o pacote principal que contém o código do seu programa.
package main

import (
	"fmt"  // Pacote fmt é usado para formatar e imprimir strings.
	"time" // Pacote time é usado para lidar com tempo, como pausas.
)

// suspense exibe uma sequência de pontos (..........) com pequenas pausas entre cada ponto
// para criar um efeito de suspense ou espera visual.
func suspense() {
	for i := 0; i < 10; i++ { // Um loop for que itera 10 vezes, uma para cada ponto.
		fmt.Print(".")                     // Imprime um ponto na mesma linha.
		time.Sleep(500 * time.Millisecond) // Pausa o programa por 500 milissegundos.
	}
	clearTerminal() // Chama a função clearTerminal para limpar o terminal.
}
