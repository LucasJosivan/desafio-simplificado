package domains

import (
	"fmt"

	"github.com/LucasJosivan/desafio-simplificado/repository"
)

var resp string

func StartGame() {
	fmt.Println("Bem vindo ao jogo Cara ou Coroa!")
	fmt.Print("Você gostaria de jogar uma partida? [s/n]: ")
	fmt.Scan(&resp)
	switch ToLower(resp) {
	case "s":
		play()
	case "n":
		fmt.Println("Tudo bem, caso queira jogar futuramente, basta executar o programa novamente.")
	default:
		fmt.Println("Opção inválida. Execute o programa novamente digitando 's' para sim ou 'n' para não.")
	}
}

func play() {
	player := Game{}
	fmt.Print("Muito bem, vou explicar como o jogo funciona.\nSempre que quiser jogar a moeda, você será 50% de chance de ganhar e 50% de chance de perder de acordo com a sua escolha de cara ou coroa\nQuando quiser parar de jogar, será mostrado seu score total, com a quantidade de partidas ganhas e perdidas.\nBoa sorte, você vai precisar!\n\n")
	play := true
	for play {
		fmt.Print("Cara ou coroa? [cara/coroa]: ")
		fmt.Scan(&resp)
		if err := Verify(ToLower(resp)); err != nil {
			fmt.Println("Você cometeu um erro: ", err)
			continue
		}
		flip := repository.Flip()
		if flip == ToLower(resp) {
			player.Wins++
			fmt.Println("Você acertou! A moeda caiu em", flip)
		} else {
			player.Loses++
			fmt.Println("Você errou! A moeda caiu em", flip)
		}
		fmt.Println("Você quer jogar novamente? [s/n]")
		fmt.Scan(&resp)
		if err := Check(ToLower(resp)); err != nil {
			fmt.Println("Você cometeu um erro: ", err)
			continue
		}
		if ToLower(resp) == "n" {
			play = false
		}
	}
	fmt.Printf("Você optou por parar de jogar, aqui está o seu placar:\nVitórias: %v\nDerrotas: %v\n", player.Wins, player.Loses)
	fmt.Println("Obrigado por jogar Cara ou Coroa. Se quiser jogar novamente, execute o programa.")
}
