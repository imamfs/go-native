package middleware

import (
	"go-native/dto"
	"go-native/helper"
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientId := r.Header.Get("Client-ID")

		if clientId == "" {
			helper.WriteJson(w, http.StatusUnauthorized, dto.TopScoreResponse[[]dto.TopScorerData]{
				Status:  false,
				Message: "Client-ID required",
				Data:    []dto.TopScorerData{},
			})
			return
		}

		if ok := validateClientId(clientId); !ok {
			helper.WriteJson(w, http.StatusUnauthorized, dto.TopScoreResponse[[]dto.TopScorerData]{
				Status:  false,
				Message: "invalid Client-ID",
				Data:    []dto.TopScorerData{},
			})
			return
		}

		next.ServeHTTP(w, r)

	})
}

func validateClientId(clientId string) bool {
	if clientId == "box2box-id" {
		return true
	}
	return false
}
