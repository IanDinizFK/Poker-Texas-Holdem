// Package main é o pacote principal onde este código está localizado.
package main

import (
	// Importações para operações de entrada/saída, formatação, aleatoriedade, execução de comandos e sistema operacional.
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

// criarJogador cria uma nova instância de um jogador.
func criarJogador(nome string, dinheiro int, tipo bool) Jogador {
	return Jogador{Nome: nome, Mao: []Carta{}, Dinheiro: dinheiro, Aposta: 0, Ativo: true, Tipo: tipo}
}

// inicializarJogadores inicializa uma lista de jogadores.
func inicializarJogadores() []Jogador {
	return []Jogador{
		criarJogador("Ian", 1000, false),
		criarJogador("Adeilson", 1000, false),
		criarJogador("Kaio", 1000, false),
		criarJogador("CPU 1", 1000, false),
	} // Cria jogadores específicos.

}

// printarInfos exibe as informações dos jogadores ativos.
func printarInfos(jogadores []Jogador) {
	for _, jogador := range jogadores {
		if jogador.Ativo {
			fmt.Printf("Nome: %s | Fichas: %d\n", jogador.Nome, jogador.Dinheiro)
			fmt.Print("\n")
			fmt.Println()
		} // Exibe informações de jogadores ativos.

	}
}

// printarCarta exibe as cartas de cada jogador humano ativo.
func printarCarta(jogadores []Jogador) {
	time.Sleep(2 * time.Second) // Aguarda 2 segundos antes de exibir as cartas.
	fmt.Printf("\n\n\n")
	clearTerminal() // Limpa o terminal para uma exibição mais limpa.

	// Itera sobre os jogadores ativos.
	for _, jogador := range jogadores {
		if jogador.Ativo && jogador.Tipo == true {
			fmt.Printf("Olá Jogador: %s, digite qualquer tecla para visualizar suas cartas:", jogador.Nome)
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()

			// Imprime as cartas do jogador.
			fmt.Print("Cartas: ")
			for _, carta := range jogador.Mao {
				fmt.Printf("[%v%s]", carta.Valor, carta.Naipe)
				time.Sleep(2 * time.Second) // Aguarda 2 segundos entre cada carta.
			}
			println()                                                             // Imprime uma nova linha.
			fmt.Println("ESSAS SÃO SUAS CARTAS, ANOTE POIS ELAS VÃO SUMIR EM 5s") // Mensagem informativa.
			fmt.Print("\n")                                                       // Imprime uma nova linha em branco.
			fmt.Println()                                                         // Imprime uma nova linha em branco.
			time.Sleep(2 * time.Second)                                           // Aguarda 2 segundos antes de limpar o terminal.
			clearTerminal()                                                       // Limpa o terminal.
		}
	}
}

// printarCartaFinal exibe as cartas de todos os jogadores ativos no final do jogo ou rodada.
func printarCartaFinal(jogadores []Jogador) {
	for _, jogador := range jogadores {
		if jogador.Ativo {
			fmt.Printf("Jogador: %s\n", jogador.Nome)
			fmt.Print("Cartas: ")
			for _, carta := range jogador.Mao {
				fmt.Printf("[%v%s]", carta.Valor, carta.Naipe)
			}
			println()
			fmt.Print("\n")
			fmt.Println()
			time.Sleep(2 * time.Second)
		}
	} // Exibe as cartas de jogadores ativos.
}

// clearTerminal limpa o terminal.
func clearTerminal() {
	// Verifique o sistema operacional para determinar o comando de limpeza apropriado.
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // Para Windows
	} else {
		cmd = exec.Command("clear") // Para sistemas Unix-like (Linux, macOS, etc.)
	}

	// Defina a saída do comando para a saída padrão
	cmd.Stdout = os.Stdout

	// Execute o comando para limpar o terminal
	if err := cmd.Run(); err != nil {
		fmt.Println("Erro ao limpar o terminal:", err)
	}
} // Limpa o terminal com base no sistema operacional.

// RearranjarJogadores reorganiza a ordem dos jogadores com base no dealer.
func RearranjarJogadores(jogadores []Jogador, dealer Jogador) {
	// Encontre a posição do dealer na lista de jogadores.
	dealerIndex := -1
	for i, jogador := range jogadores {
		if jogador.Nome == dealer.Nome {
			dealerIndex = i
			break
		}
	}

	// Se o dealer não for encontrado na lista de jogadores, não há nada a fazer.
	if dealerIndex == -1 {
		return
	}

	// Determina o número de jogadores na lista.
	numJogadores := len(jogadores)

	// Cria uma nova slice de jogadores com a mesma capacidade.
	novaOrdem := make([]Jogador, numJogadores)

	// Copia os jogadores à direita do dealer para a nova ordem.
	copy(novaOrdem[0:], jogadores[dealerIndex+1:])

	// Copia os jogadores à esquerda do dealer para o final da nova ordem.
	copy(novaOrdem[numJogadores-dealerIndex-1:], jogadores[:dealerIndex+1])

	// Atualiza a lista de jogadores com a nova ordem.
	copy(jogadores, novaOrdem)
} // Reordena os jogadores de forma que o dealer seja o último.

// DeterminarDealer determina o dealer com base em uma sorte aleatória.
func DeterminarDealer(jogadores []Jogador) Jogador {
	// Inicialize a semente do gerador de números aleatórios
	rand.Seed(time.Now().UnixNano())

	// Atribua números aleatórios aos jogadores, representando as cartas de 2 a Ás.
	for i := range jogadores {
		jogadores[i].NumSorte = rand.Intn(13) + 2 // Números de 2 a 14
	}

	// Encontre o jogador com o número mais alto como dealer
	var dealer Jogador
	for _, jogador := range jogadores {
		if jogador.NumSorte > dealer.NumSorte {
			dealer = jogador
		}
	}
	return dealer
} // Atribui um número aleatório para cada jogador e retorna o jogador com o número mais alto como dealer.
