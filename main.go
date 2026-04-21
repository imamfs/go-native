package main

import (
	"context"
	"go-native/database"
	"go-native/database/query"
	"go-native/handler"
	"go-native/middleware"
	"go-native/service"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := query.NewQuery(db)
	service := service.NewTopScorerService(repo)
	handler := handler.NewTopScorerHandler(service)

	mux := http.NewServeMux()

	globalMiddleware := middleware.Chain(
		middleware.Recovery, middleware.Auth, middleware.Logger,
	)

	mux.Handle("/topscorers", globalMiddleware(http.HandlerFunc(handler.AddScorer)))
	mux.Handle("/topscorers/update/goal", globalMiddleware(http.HandlerFunc(handler.UpdateGoal)))
	mux.Handle("/topscorers/update/team", globalMiddleware(http.HandlerFunc(handler.UpdateTeam)))
	mux.Handle("/topscorers/data/player/{name}", globalMiddleware(http.HandlerFunc(handler.GetScorer)))
	mux.Handle("/topscorers/data/player", globalMiddleware(http.HandlerFunc(handler.GetAllScorers)))
	mux.Handle("/topscorers/remove", globalMiddleware(http.HandlerFunc(handler.RemovePlayer)))

	log.Println("Server run on port :8080")
	log.Fatal(http.ListenAndServe(":8080", globalMiddleware(mux)))
}
