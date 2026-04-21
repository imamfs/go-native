package handler

import (
	"encoding/json"
	"go-native/dto"
	"go-native/helper"
	"go-native/service"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type handler struct {
	topScorerService service.TopScoreService
}

func NewTopScorerHandler(topScorerService service.TopScoreService) handler {
	return handler{topScorerService: topScorerService}
}

func (h handler) AddScorer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.WriteJson(w, http.StatusMethodNotAllowed, dto.TopScoreResponse[[]dto.TopScorerData]{
			Status:  false,
			Message: "invalid http method",
			Data:    []dto.TopScorerData{},
		})
		return
	}

	var request dto.TopScoreRequest

	decodeBody := json.NewDecoder(r.Body)
	if err := decodeBody.Decode(&request); err != nil {
		helper.WriteJson(w, http.StatusBadRequest, dto.TopScoreResponse[[]dto.TopScorerData]{
			Status:  false,
			Message: "invalid body request",
			Data:    []dto.TopScorerData{},
		})
		return
	}

	if err := helper.TopScorerValidator(request); err != nil {
		helper.WriteJson(w, http.StatusBadRequest, dto.TopScoreResponse[[]dto.TopScorerData]{
			Status:  false,
			Message: err.Error(),
			Data:    []dto.TopScorerData{},
		})
		return
	}

	data, httpStatus := h.topScorerService.AddScorer(r.Context(), request)
	if httpStatus == http.StatusInternalServerError {
		helper.WriteJson(w, httpStatus, dto.TopScoreResponse[[]dto.TopScorerData]{
			Status:  false,
			Message: "internal server error",
			Data:    []dto.TopScorerData{},
		})
		return
	}

	helper.WriteJson(w, httpStatus, dto.TopScoreResponse[[]dto.TopScorerData]{
		Status:  true,
		Message: "success",
		Data:    data,
	})
}

func (h handler) UpdateGoal(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		helper.WriteJson(w, http.StatusMethodNotAllowed, dto.TopScoreResponse[dto.TopScorerData]{
			Status:  false,
			Message: "invalid http method",
			Data:    dto.TopScorerData{},
		})
		return
	}

	query := r.URL.Query()
	name := strings.TrimSpace(query.Get("name"))
	goal, err := strconv.Atoi(query.Get("goal"))

	if err != nil {
		helper.WriteJson(w, http.StatusBadRequest, dto.TopScoreResponse[dto.TopScorerData]{
			Status:  false,
			Message: "invalid data type for goal",
			Data:    dto.TopScorerData{},
		})
		return
	}

	if name == "" {
		helper.WriteJson(w, http.StatusBadRequest, dto.TopScoreResponse[dto.TopScorerData]{
			Status:  false,
			Message: "player name is required",
			Data:    dto.TopScorerData{},
		})
		return
	}

	if goal <= 0 {
		helper.WriteJson(w, http.StatusBadRequest, dto.TopScoreResponse[dto.TopScorerData]{
			Status:  false,
			Message: "goal must be filled",
			Data:    dto.TopScorerData{},
		})
		return
	}

	data, httpStatus := h.topScorerService.UpdateGoal(r.Context(), name, goal)
	if httpStatus == http.StatusInternalServerError {
		helper.WriteJson(w, httpStatus, dto.TopScoreResponse[dto.TopScorerData]{
			Status:  false,
			Message: "internal server error",
			Data:    dto.TopScorerData{},
		})
		return
	}

	helper.WriteJson(w, httpStatus, dto.TopScoreResponse[dto.TopScorerData]{
		Status:  true,
		Message: "success",
		Data:    data,
	})
}

func (h handler) UpdateTeam(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		helper.WriteJson(w, http.StatusMethodNotAllowed, dto.TopScoreResponse[dto.TopScorerData]{
			Status:  false,
			Message: "invalid http method",
			Data:    dto.TopScorerData{},
		})
		return
	}

	query := r.URL.Query()
	name := strings.TrimSpace(query.Get("name"))
	club := strings.TrimSpace(query.Get("club"))

	if name == "" {
		helper.WriteJson(w, http.StatusBadRequest, dto.TopScoreResponse[dto.TopScorerData]{
			Status:  false,
			Message: "player name is required",
			Data:    dto.TopScorerData{},
		})
		return
	}

	if club == "" {
		helper.WriteJson(w, http.StatusBadRequest, dto.TopScoreResponse[dto.TopScorerData]{
			Status:  false,
			Message: "player's club must define",
			Data:    dto.TopScorerData{},
		})
		return
	}

	data, httpStatus := h.topScorerService.UpdateTeam(r.Context(), name, club)
	if httpStatus == http.StatusInternalServerError {
		helper.WriteJson(w, httpStatus, dto.TopScoreResponse[dto.TopScorerData]{
			Status:  false,
			Message: "internal server error",
			Data:    dto.TopScorerData{},
		})
		return
	}

	helper.WriteJson(w, httpStatus, dto.TopScoreResponse[dto.TopScorerData]{
		Status:  true,
		Message: "success",
		Data:    data,
	})
}

func (h handler) GetScorer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.WriteJson(w, http.StatusMethodNotAllowed, dto.TopScoreResponse[dto.TopScorerData]{
			Status:  false,
			Message: "invalid http method",
			Data:    dto.TopScorerData{},
		})
		return
	}

	regex := regexp.MustCompile(`[^a-zA-Z0-9 ]`)
	name := regex.ReplaceAllString(r.PathValue("name"), " ")

	if name == "" {
		helper.WriteJson(w, http.StatusBadRequest, dto.TopScoreResponse[dto.TopScorerData]{
			Status:  false,
			Message: "player's name is required",
			Data:    dto.TopScorerData{},
		})
		return
	}

	data, httpStatus := h.topScorerService.FindScorer(r.Context(), name)
	if httpStatus == http.StatusNotFound {
		helper.WriteJson(w, httpStatus, dto.TopScoreResponse[dto.TopScorerData]{
			Status:  false,
			Message: "player not found",
			Data:    dto.TopScorerData{},
		})
		return
	}

	if httpStatus == http.StatusInternalServerError {
		helper.WriteJson(w, httpStatus, dto.TopScoreResponse[dto.TopScorerData]{
			Status:  false,
			Message: "internal server error",
			Data:    dto.TopScorerData{},
		})
		return
	}

	helper.WriteJson(w, httpStatus, dto.TopScoreResponse[dto.TopScorerData]{
		Status:  true,
		Message: "success",
		Data:    data,
	})
}

func (h handler) GetAllScorers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.WriteJson(w, http.StatusMethodNotAllowed, dto.TopScoreResponse[[]dto.TopScorerData]{
			Status:  false,
			Message: "invalid http method",
			Data:    []dto.TopScorerData{},
		})
		return
	}

	data, httpStatus := h.topScorerService.FindAllScorers(r.Context())
	if httpStatus == http.StatusInternalServerError {
		helper.WriteJson(w, httpStatus, dto.TopScoreResponse[[]dto.TopScorerData]{
			Status:  false,
			Message: "internal server error",
			Data:    []dto.TopScorerData{},
		})
		return
	}

	helper.WriteJson(w, httpStatus, dto.TopScoreResponse[[]dto.TopScorerData]{
		Status:  true,
		Message: "success",
		Data:    data,
	})
}

func (h handler) RemovePlayer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		helper.WriteJson(w, http.StatusMethodNotAllowed, dto.TopScoreResponse[dto.TopScorerData]{
			Status:  false,
			Message: "invalid http method",
			Data:    dto.TopScorerData{},
		})
		return
	}

	query := r.URL.Query()
	name := query.Get("name")

	if name == "" {
		helper.WriteJson(w, http.StatusBadRequest, dto.TopScoreResponse[dto.TopScorerData]{
			Status:  false,
			Message: "invalid http method",
			Data:    dto.TopScorerData{},
		})
		return
	}

	data, httpStatus := h.topScorerService.RemoveScorer(r.Context(), name)
	if httpStatus == http.StatusInternalServerError {
		helper.WriteJson(w, httpStatus, dto.TopScoreResponse[dto.TopScorerData]{
			Status:  false,
			Message: "internal server error",
			Data:    data,
		})
		return
	}

	helper.WriteJson(w, httpStatus, dto.TopScoreResponse[dto.TopScorerData]{
		Status:  true,
		Message: "success",
		Data:    data,
	})
}
