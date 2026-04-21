package service

import (
	"context"
	"errors"
	"fmt"
	"go-native/database/query"
	"go-native/domain/repository"
	"go-native/dto"
	"net/http"
)

/*var Data = []dto.TopScorerData{
	{Name: "Erling Haaland", Club: "Manchester City", Position: "Forward", Goals: 22},
	{Name: "Igor Thiago", Club: "Brentford", Position: "Forward", Goals: 21},
	{Name: "Antoine Semenyo", Club: "Manchester City", Position: "Right Wing", Goals: 15},
}*/

type topScorerService struct {
	scorerRepository repository.TopScorerRepository
}

func NewTopScorerService(scorerRepository repository.TopScorerRepository) TopScoreService {
	return &topScorerService{scorerRepository}
}

func (s *topScorerService) AddScorer(ctx context.Context, request dto.TopScoreRequest) ([]dto.TopScorerData, int) {
	var scorers []dto.TopScorerData

	if err := s.scorerRepository.AddScorer(ctx, request.Data); err != nil {
		fmt.Println(err)
		return scorers, http.StatusInternalServerError
	}

	data, err := s.scorerRepository.FindAll(ctx)
	if err != nil {
		fmt.Println(err)
		return scorers, http.StatusInternalServerError
	}

	for _, v := range data {
		var scorer dto.TopScorerData
		scorer.Name = v.Name
		scorer.Club = v.Club
		scorer.Position = v.Position
		scorer.Goals = v.Goals
		scorers = append(scorers, scorer)
	}

	return scorers, http.StatusOK

	/*
		for _, scorer := range request.Data {
			Data = append(Data, scorer)
		}
		SortTopScorerDesc()
		return Data, nil*/
}

func (s *topScorerService) UpdateGoal(ctx context.Context, name string, goal int) (dto.TopScorerData, int) {
	var scorer dto.TopScorerData

	data, err := s.scorerRepository.UpdateGoal(ctx, name, goal)
	if err != nil {
		fmt.Println(err)
		return scorer, http.StatusInternalServerError
	}

	scorer.Name = data.Name
	scorer.Club = data.Club
	scorer.Position = data.Position
	scorer.Goals = data.Goals

	return scorer, http.StatusOK
}

func (s *topScorerService) UpdateTeam(ctx context.Context, name string, team string) (dto.TopScorerData, int) {
	var scorer dto.TopScorerData

	data, err := s.scorerRepository.UpdateTeam(ctx, name, team)
	if err != nil {
		fmt.Println(err)
		return scorer, http.StatusInternalServerError
	}

	scorer.Name = data.Name
	scorer.Club = data.Club
	scorer.Position = data.Position
	scorer.Goals = data.Goals

	return scorer, http.StatusOK
}

func (s *topScorerService) FindScorer(ctx context.Context, name string) (dto.TopScorerData, int) {
	var scorer dto.TopScorerData

	data, err := s.scorerRepository.FindScorerByName(ctx, name)
	if err != nil {
		if errors.Is(err, query.ErrPlayerNotFound) {
			fmt.Println(err)
			return scorer, http.StatusNotFound
		}
		fmt.Println(err)
		return scorer, http.StatusInternalServerError
	}

	scorer.Name = data.Name
	scorer.Club = data.Club
	scorer.Position = data.Position
	scorer.Goals = data.Goals

	return scorer, http.StatusOK
}

func (s *topScorerService) FindAllScorers(ctx context.Context) ([]dto.TopScorerData, int) {
	var scorers []dto.TopScorerData

	data, err := s.scorerRepository.FindAll(ctx)
	if err != nil {
		fmt.Println(err)
		return scorers, http.StatusInternalServerError
	}

	for _, v := range data {
		var scorer dto.TopScorerData
		scorer.Name = v.Name
		scorer.Club = v.Club
		scorer.Position = v.Position
		scorer.Goals = v.Goals
		scorers = append(scorers, scorer)
	}

	return scorers, http.StatusOK
}

func (s *topScorerService) RemoveScorer(ctx context.Context, name string) (dto.TopScorerData, int) {
	var scorer dto.TopScorerData

	if err := s.scorerRepository.RemoveScorer(ctx, name); err != nil {
		fmt.Println(err)
		return scorer, http.StatusInternalServerError
	}

	/*for i, v := range Data {
		if name == v.Name {
			Data = append(Data[:i], Data[i+1:]...)
		}
	}
	return nil*/
	return scorer, http.StatusOK
}

/*
func GetTopScorers() []dto.TopScorerData {
	SortTopScorerDesc()
	return Data
}

func SortTopScorerDesc() {
	sort.Slice(Data, func(i, j int) bool {
		return Data[i].Goals > Data[j].Goals
	})
}
*/
