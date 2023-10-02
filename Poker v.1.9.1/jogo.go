package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Variáveis globais
var rodada int      // Controla a rodada atual
var vencedorON bool // Indica se um vencedor foi determinado
var passarVez bool  // Indica se um jogador passou a vez ou desistiu

// inicioRodadaApostas inicia as apostas dos dois primeiros jogadores.
func inicioRodadaApostas(jogadores []Jogador, mesa *Mesa) {
	apostaMinima := 20
	for i := 0; i <= 1; i++ {
		jogador := &jogadores[i]
		if i == 0 {
			fmt.Printf("SMALLBLIND: %s, FICHAS: %d\n", jogador.Nome, apostaMinima/2)
			jogador.Dinheiro -= apostaMinima / 2
			mesa.Pote += apostaMinima / 2
		} else {
			fmt.Printf("BIGBLIND: %s, FICHAS: %d\n", jogador.Nome, apostaMinima)
			jogador.Dinheiro -= apostaMinima
			mesa.Pote += apostaMinima
		}
		timer := time.NewTimer(2 * time.Second)
		<-timer.C
	}
}

// primeiraRodadaApostas gerencia apostas dos jogadores após a distribuição das cartas.
func primeiraRodadaApostas(jogadores []Jogador, mesa *Mesa) {
	apostaMinima := 20
	for i := 2; i < len(jogadores); i++ {
		jogador := &jogadores[i]
		if jogador.Ativo && jogador.Tipo {
			if !passarVez {
				fmt.Printf("%s, Deseja apostar (Sim/Não)? !NÃO PARA DESISTIR DA MÃO!: ", jogador.Nome)
				if solicitarRespostaSimOuNao() {
					aposta := fazerAposta(jogador, apostaMinima)
					mesa.Pote += aposta
					if aposta > apostaMinima {
						apostaMinima = aposta
					}
				} else {
					jogador.Ativo = false
					passarVez = true
				}
			} else {
				fmt.Printf("%s, O jogador anterior desistiu/Passou a vez. Deseja apostar? (Sim/Não): ", jogador.Nome)
				if solicitarRespostaSimOuNao() {
					aposta := fazerAposta(jogador, apostaMinima)
					mesa.Pote += aposta
					if aposta > apostaMinima {
						apostaMinima = aposta
					}
				}
			}
		} else if jogador.Ativo && !jogador.Tipo {
			aposta := decidirApostaBot(jogador, mesa.CartasComunitarias, apostaMinima, rodada, mesa.Pote)
			if aposta < 0 {
				jogador.Ativo = false
				passarVez = true
				fmt.Printf("-->!AÇÃO BOT! O BOT %s Desistiu da sua mão <--\n", jogador.Nome)
			} else if aposta > 0 {
				apostaMinima = aposta
				mesa.Pote += aposta
				jogador.Dinheiro -= aposta
				fmt.Printf("-->!AÇÃO BOT! O BOT %s Apostou: %d <--\n", jogador.Nome, aposta)
			} else {
				fmt.Printf("-->!AÇÃO BOT! O BOT %s PASSOU SUA VEZ <--\n", jogador.Nome)
			}
		}
	}
}

// rodadaApostas gerencia apostas em rodadas subsequentes.
func verificaVencedorDesistencia(jogadores []Jogador) {
	cont := 0
	for i := range jogadores {
		jogador := &jogadores[i]
		if jogador.Ativo {
			cont++
		}
	}
	if cont == 1 {
		vencedorON = true
	}
}

// fazerAposta permite que um jogador faça uma aposta.
func solicitarRespostaSimOuNao() bool {
	for {
		var resposta string
		_, err := fmt.Scan(&resposta)
		if err != nil {
			fmt.Println(err)
			return false
		}
		resposta = strings.ToLower(strings.TrimSpace(resposta))
		if resposta == "sim" || resposta == "s" {
			return true
		} else if resposta == "não" || resposta == "nao" || resposta == "n" {
			return false
		} else {
			fmt.Print("Por favor, responda sim ou não: ")
		}
	}
}

// solicitarRespostaSimOuNao aguarda uma resposta "sim" ou "não" do usuário.
func infoAllIn(jogador *Jogador) {
	fmt.Printf("Ola Player %s\n", jogador.Nome)
	fmt.Printf("Foi notado que você possui quantidade de fichas menor que a aposta Minima\n")
	fmt.Printf("Caso você concorde em apostar, terá que fazer All-In\n")
}

// infoAllIn informa ao jogador sobre a situação do All-In.
func rodadaApostas(jogadores []Jogador, mesa *Mesa) {
	apostaMinima := 20
	passarVez = false
	for i := range jogadores {
		jogador := &jogadores[i]
		if jogador.Ativo && jogador.Tipo {
			if !passarVez {
				if jogador.Dinheiro < apostaMinima {
					infoAllIn(jogador)
				}
				fmt.Printf("%s, Deseja apostar (Sim/Não)? !NÃO PARA DESISTIR DA MÃO!: ", jogador.Nome)
				if solicitarRespostaSimOuNao() {
					aposta := fazerAposta(jogador, apostaMinima)
					mesa.Pote += aposta
					if aposta > apostaMinima {
						apostaMinima = aposta
					}
				} else {
					jogador.Ativo = false
					passarVez = true
				}
			} else {
				if jogador.Dinheiro < apostaMinima {
					infoAllIn(jogador)
				}
				fmt.Printf("%s, O jogador anterior desistiu/Passou a vez. Deseja apostar? (Sim/Não): ", jogador.Nome)
				if solicitarRespostaSimOuNao() {
					aposta := fazerAposta(jogador, apostaMinima)
					mesa.Pote += aposta
					if aposta > apostaMinima {
						apostaMinima = aposta
					}
				}
			}
		} else if jogador.Ativo && !jogador.Tipo {
			aposta := decidirApostaBot(jogador, mesa.CartasComunitarias, apostaMinima, rodada, mesa.Pote)
			if aposta < 0 {
				jogador.Ativo = false
				apostaMinima = 0
				fmt.Printf("-->!AÇÃO BOT! %s Desistiu da sua mão <--\n", jogador.Nome)
			} else if aposta > 0 {
				apostaMinima = aposta
				mesa.Pote += aposta
				jogador.Aposta += aposta
				fmt.Printf("-->!AÇÃO BOT! %s Apostou: %d <-- \n", jogador.Nome, aposta)
			}
		}
	}
}

// verificaVencedorDesistencia verifica se há um vencedor devido à desistência dos jogadores.
func fazerAposta(jogador *Jogador, apostaMinima int) int {
	fmt.Printf("JOGADOR: %s, você possui %d Fichas\n", jogador.Nome, jogador.Dinheiro)
	for {
		var apostaInput string
		fmt.Printf("%s, quanto deseja apostar? (mínimo R$%d ou '0' para all in): ", jogador.Nome, apostaMinima)
		_, err := fmt.Scan(&apostaInput)
		if err != nil {
			fmt.Printf("\nAposta inválida! Deve ser um número ou '0'. Tente novamente: ")
			continue
		}

		if apostaInput == "0" {
			aposta := jogador.Dinheiro
			jogador.Aposta = aposta
			jogador.Dinheiro -= aposta
			return aposta
		}

		aposta, err := strconv.Atoi(apostaInput)
		if err != nil {
			fmt.Printf("\nAposta inválida! Deve ser um número ou '0'. Tente novamente: ")
			continue
		}

		if aposta < apostaMinima {
			fmt.Printf("\nAposta inválida! Deve ser maior ou igual a %d ou '0' para all in. Tente novamente: ", apostaMinima)
			continue
		}

		if aposta > jogador.Dinheiro {
			fmt.Printf("\nVocê não possui fichas suficientes para essa aposta. Tente novamente ou faça um all in: ")
			continue
		}

		jogador.Aposta = aposta
		jogador.Dinheiro -= aposta
		return aposta
	}
}

// jogarPoker controla o fluxo de uma partida de poker.
// Recebe uma lista de jogadores, uma mesa e um baralho como entrada.
func jogarPoker(jogadores []Jogador, mesa *Mesa, baralho []Carta) {
	// Inicializa variáveis de controle da partida
	vencedorON = false
	rodada = 1

	// Ativa os jogadores
	ativaJogadores(jogadores)

	// Exibe mensagem de "Decidindo o dealer" com uma pausa
	fmt.Print("Decidindo o dealer...\n")
	timer := time.NewTimer(2 * time.Second)
	<-timer.C

	// Determina o dealer e rearranja a ordem dos jogadores
	dealer := DeterminarDealer(jogadores)
	fmt.Printf("O dealer sorteado foi: %s\n", dealer.Nome)
	RearranjarJogadores(jogadores, dealer)

	// Loop principal da partida
	for !vencedorON {
		switch rodada {
		case 1:
			fmt.Println("Início da Rodada 1: Pré-flop")
			inicioRodadaApostas(jogadores, mesa)
			baralho = distribuirCartas(jogadores, baralho, 2)
			printarCarta(jogadores)
			primeiraRodadaApostas(jogadores, mesa)
			baralho = distribuirCartasMesa(mesa, baralho, 3) // flop

		case 2:
			avaliarMao(jogadores, mesa.CartasComunitarias) // CHAMA A FUNÇÃO DE AVALIAR MAO
			time.Sleep(4 * time.Second)
			clearTerminal()
			fmt.Println("Início da Rodada 2: Flop")
			printarInfoMesa(mesa)
			rodadaApostas(jogadores, mesa)
			verificaVencedorDesistencia(jogadores)
			if vencedorON {
				break
			}
			baralho = distribuirCartasMesa(mesa, baralho, 1) // turn

		case 3:
			avaliarMao(jogadores, mesa.CartasComunitarias) // CHAMA A FUNÇÃO DE AVALIAR MAO
			time.Sleep(4 * time.Second)
			clearTerminal()
			fmt.Println("Início da Rodada 3: Turn")
			printarInfoMesa(mesa)
			rodadaApostas(jogadores, mesa)
			verificaVencedorDesistencia(jogadores)
			if vencedorON {
				break
			}
			baralho = distribuirCartasMesa(mesa, baralho, 1) // river

		case 4:
			avaliarMao(jogadores, mesa.CartasComunitarias) // CHAMA A FUNÇÃO DE AVALIAR MAO
			time.Sleep(4 * time.Second)
			clearTerminal()
			fmt.Println("Início da Rodada 4: River")
			printarInfos(jogadores)
			printarInfoMesa(mesa)
			rodadaApostas(jogadores, mesa)
			time.Sleep(3 * time.Second)
			clearTerminal()
			fmt.Println("-->REVELANDO AS CARTAS DOS JOGADORES<--")
			suspense()
			printarCartaFinal(jogadores)
			fmt.Printf("\n")
			printarInfoMesa(mesa)
			fmt.Printf("\n\n")
			time.Sleep(5 * time.Second)
			decidirVencedor(jogadores, mesa)
			bonificaVencedorDesistencia(jogadores, mesa)
		}
		rodada++
	}
}

// ativaJogadores ativa todos os jogadores no início de uma rodada.
func ativaJogadores(jogadores []Jogador) {
	for i := range jogadores {
		jogadores[i].Ativo = true
	}
}

// bonificaVencedorDesistencia concede o prêmio ao vencedor e reinicia o pote.
func bonificaVencedorDesistencia(jogadores []Jogador, mesa *Mesa) {
	fmt.Printf("PRÊMIO DO VENCEDOR: %d\n", mesa.Pote)

	// Loop pelos jogadores para determinar o vencedor e conceder o prêmio
	for i := range jogadores {
		jogador := &jogadores[i]
		if jogador.Ativo {
			jogador.Dinheiro += mesa.Pote
			mesa.Pote = 0
		}
	}
}

// decidirVencedor determina o vencedor com base na força da mão de cada jogador.
func decidirVencedor(jogadores []Jogador, mesa *Mesa) {
	var vencedor *Jogador

	// Loop pelos jogadores para determinar o vencedor
	for i := range jogadores {
		jogador := &jogadores[i]
		if jogador.Ativo {
			nivelMao := jogador.levelMao

			// Verifica se o jogador atual tem uma mão mais forte do que o vencedor atual
			if vencedor == nil || nivelMao > vencedor.levelMao {
				if vencedor != nil {
					// Define o jogador anterior como inativo se houver um novo vencedor
					vencedor.Ativo = false
				}
				vencedor = jogador
				jogador.levelMao = nivelMao
			} else {
				// Se o jogador atual não tem uma mão mais forte, torna-o inativo
				jogador.Ativo = false
			}
		}
	}

	// Exibe o vencedor, a mão vencedora e a premiação
	if vencedor != nil {
		fmt.Print("E o vencedor é")
		suspense()
		fmt.Printf("------>%s<------\n", vencedor.Nome)
		time.Sleep(500 * time.Microsecond)
		fmt.Printf("LEVOU A BOLADA DE %d FICHAS\n", mesa.Pote)
		time.Sleep(500 * time.Microsecond)
		fmt.Println("COM A MÃO")
		printarCartaFinal(jogadores)
		fmt.Printf("E VOCÊ FEZ UM: %s\n", vencedor.tipoMao)
	} else {
		fmt.Println("Não há vencedor.")
	}

	// Ativa a flag de vencedor para finalizar a partida
	vencedorON = true
}

func confirmarContinuacao() bool {
	fmt.Print("Deseja continuar jogando? (Sim/Não): ")
	var resposta string
	fmt.Scanln(&resposta)
	resposta = strings.ToLower(resposta)
	return resposta == "sim" || resposta == "s"
}
