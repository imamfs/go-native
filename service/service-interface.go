package service

import (
	"context"
	"go-native/dto"
)

type TopScoreService interface {
	AddScorer(ctx context.Context, request dto.TopScoreRequest) ([]dto.TopScorerData, int)
	UpdateGoal(ctx context.Context, name string, goal int) (dto.TopScorerData, int)
	UpdateTeam(ctx context.Context, name string, team string) (dto.TopScorerData, int)
	FindScorer(ctx context.Context, name string) (dto.TopScorerData, int)
	FindAllScorers(ctx context.Context) ([]dto.TopScorerData, int)
	RemoveScorer(ctx context.Context, name string) (dto.TopScorerData, int)
}
