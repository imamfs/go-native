package dto

type TopScoreRequest struct {
	Data TopScorerData `json:"data"`
}

type TopScoreResponse[T any] struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type TopScorerData struct {
	Name     string `json:"name"`
	Club     string `json:"club"`
	Position string `json:"position"`
	Goals    int    `json:"goals"`
}
