package query

import (
	"context"
	"errors"
	"fmt"
	"go-native/database/model"
	"go-native/domain/repository"
	"go-native/dto"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type sqlQuery struct {
	pgpool *pgxpool.Pool
}

func NewQuery(pgpool *pgxpool.Pool) repository.TopScorerRepository {
	return &sqlQuery{pgpool}
}

var ErrPlayerNotFound = errors.New("player not found")

func (q *sqlQuery) AddScorer(ctx context.Context, scorer dto.TopScorerData) error {
	query := `INSERT INTO top_scorer (name, club, position, goals) VALUES ($1, $2, $3, $4)`

	result, err := q.pgpool.Exec(ctx, query, &scorer.Name, &scorer.Club, &scorer.Position, &scorer.Goals)
	if err != nil {
		return fmt.Errorf("failed to add scorer, error: %w\n", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("no new row added")
	}

	return nil
}

func (q *sqlQuery) UpdateGoal(ctx context.Context, name string, goal int) (model.TopScorer, error) {
	query := `UPDATE top_scorer SET goals = goals + $1 WHERE name = $2
	RETURNING name, club, position, goals`

	var scorer model.TopScorer
	err := q.pgpool.QueryRow(ctx, query, goal, name).Scan(&scorer.Name, &scorer.Club, &scorer.Position, &scorer.Goals)
	if err != nil {
		return model.TopScorer{}, fmt.Errorf("failed to update goal for player: %s, error: %w\n", name, err)
	}

	return scorer, nil
}

func (q *sqlQuery) UpdateTeam(ctx context.Context, name string, clubName string) (model.TopScorer, error) {
	query := `UPDATE top_scorer SET club = $1 WHERE name = $2
	RETURNING name, club, position, goals`

	var scorer model.TopScorer
	err := q.pgpool.QueryRow(ctx, query, clubName, name).Scan(&scorer.Name, &scorer.Club, &scorer.Position, &scorer.Goals)
	if err != nil {
		return model.TopScorer{}, fmt.Errorf("failed to update goal for player: %s, error: %w\n", name, err)
	}

	return scorer, nil
}

func (q *sqlQuery) FindScorerByName(ctx context.Context, name string) (model.TopScorer, error) {
	query := `SELECT name, club, position, goals FROM top_scorer WHERE name = $1`

	var scorer model.TopScorer
	err := q.pgpool.QueryRow(ctx, query, name).Scan(&scorer.Name, &scorer.Club, &scorer.Position, &scorer.Goals)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.TopScorer{}, ErrPlayerNotFound
		}
		return model.TopScorer{}, fmt.Errorf("failed to find player: %s, error: %w\n", name, err)
	}

	return scorer, nil
}

func (q *sqlQuery) FindAll(ctx context.Context) ([]model.TopScorer, error) {
	var scorers []model.TopScorer

	query := `SELECT name, club, position, goals FROM top_scorer ORDER BY goals DESC`
	rows, err := q.pgpool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all top scorer data, error: %w\n", err)
	}
	defer rows.Close()

	for rows.Next() {
		var scorer model.TopScorer
		if err := rows.Scan(&scorer.Name, &scorer.Club, &scorer.Position, &scorer.Goals); err != nil {
			return nil, fmt.Errorf("failed to scan top scorer list, error: %w\n", err)
		}
		scorers = append(scorers, scorer)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w\n", err)
	}

	return scorers, nil
}

func (q *sqlQuery) RemoveScorer(ctx context.Context, name string) error {
	query := `DELETE FROM top_scorer WHERE name = $1`

	result, err := q.pgpool.Exec(ctx, query, name)
	if err != nil {
		return fmt.Errorf("failed to remove scorer for name %s, error: %w", name, err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("No row deleted")
	}

	return nil
}
