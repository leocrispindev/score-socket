package service

import (
	"score-socket/internal/model"
)

func GetMatches() []model.Match {
	return []model.Match{
		{
			TeamA: "Equipe A1",
			TeamB: "Equipe B1",
			ID:    1,
			Score: model.Score{
				TeamA: 0,
				TeamB: 0,
			},
		},
		{
			TeamA: "Equipe A2",
			TeamB: "Equipe B2",
			ID:    2,
			Score: model.Score{
				TeamA: 0,
				TeamB: 0,
			},
		},
		{
			TeamA: "Equipe A3",
			TeamB: "Equipe B3",
			ID:    3,
			Score: model.Score{
				TeamA: 0,
				TeamB: 0,
			},
		},
	}

}
