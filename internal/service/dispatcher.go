package service

import (
	"encoding/json"
	"fmt"
	"score-socket/internal/model"
)

func Dispatch(match *model.Match) {
	for _, conn := range GetSubscribersByMatchID(match.ID) {
		body, _ := json.Marshal(match)

		err := conn.WriteMessage(1, body)
		if err != nil {
			fmt.Println("Erro ao escrever para a conexão:", err)
		}
	}
}
