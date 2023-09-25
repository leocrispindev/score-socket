package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"score-socket/internal/model"
	"score-socket/internal/service"
	"time"
)

type Matchs struct {
	Datetime string        `json:"datetime"`
	MatchDay []model.Match `json:"matchs"`
}

func HandlerGet(resp http.ResponseWriter, req *http.Request) {
	mat := Matchs{
		Datetime: time.Now().Format(time.RFC822),
		MatchDay: service.GetMatches(),
	}

	body, err := json.Marshal(mat)
	if err != nil {
		log.Println(err)
		return
	}

	resp.Write(body)
}
