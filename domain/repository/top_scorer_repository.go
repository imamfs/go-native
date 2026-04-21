package repository

import (
	"context"
	"go-native/database/model"
	"go-native/dto"
)

type TopScorerRepository interface {
	AddScorer(ctx context.Context, scorer dto.TopScorerData) error
	UpdateGoal(ctx context.Context, name string, goal int) (model.TopScorer, error)
	UpdateTeam(ctx context.Context, name string, clubName string) (model.TopScorer, error)
	FindScorerByName(ctx context.Context, name string) (model.TopScorer, error)
	FindAll(ctx context.Context) ([]model.TopScorer, error)
	RemoveScorer(ctx context.Context, name string) error
}
