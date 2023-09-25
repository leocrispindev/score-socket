package service

import (
	"math/rand"
	"score-socket/internal/model"
	"time"
)

var mensagensDeFutebol = []string{
	"Gol! O time da casa abre o placar com um chute incrível no ângulo direito!",
	"Defesa espetacular do goleiro visitante para evitar o gol!",
	"Escanteio para o time visitante. Eles têm uma ótima oportunidade de marcar aqui.",
	"Falta perigosa na entrada da área! O time da casa tem uma chance de ouro agora.",
	"Cartão amarelo para o jogador número 7 do time visitante por uma falta dura.",
	"Substituição no time da casa. O jogador número 10 está saindo e o número 23 está entrando em campo.",
	"Impedimento! O jogador número 9 do time da casa estava adiantado quando a bola foi lançada.",
	"Pênalti! O árbitro marca um pênalti a favor do time visitante após uma falta dentro da área.",
	"Gol anulado! O jogador estava em posição de impedimento quando marcou o gol.",
	"Final do primeiro tempo. O placar está empatado em 1 a 1. O jogo está emocionante!",
}

func Start(channel chan model.Match) {
	sendMessage(channel)
}

func sendMessage(channel chan model.Match) {

	for {

		currentTime := time.Now()

		// Extract the minute and second components
		minute := currentTime.Minute()
		second := currentTime.Second()

		for _, m := range GetMatches() {
			m.Event = getRandomMessage()
			m.Minute = minute
			m.Second = second
			updateScore(&m)

			Dispatch(&m)
		}

		time.Sleep(time.Duration(1000))
	}
}

func getRandomMessage() string {

	return mensagensDeFutebol[rand.Intn(9)]

}

func updateScore(match *model.Match) {
	match.Score.TeamA += 1
	match.Score.TeamB += 1
}
