package helper

import (
	"encoding/json"
	"fmt"
	"go-native/dto"
	"net/http"
)

func WriteJson[T any](w http.ResponseWriter, statusCode int, body dto.TopScoreResponse[T]) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(parseToJson(body))
}

func parseToJson[T any](body dto.TopScoreResponse[T]) []byte {
	data, err := json.Marshal(body)
	if err != nil {
		fmt.Printf("failed to marshal response error: %v\n", err)
	}
	return data
}

func TopScorerValidator(request dto.TopScoreRequest) error {
	if request.Data.Name == "" {
		return fmt.Errorf("player name is required")
	}

	if request.Data.Club == "" {
		return fmt.Errorf("club name is required")
	}

	if request.Data.Position == "" {
		return fmt.Errorf("position is required")
	}

	if request.Data.Goals == 0 {
		return fmt.Errorf("goals must filled")
	}
	return nil
}
