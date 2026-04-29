package application

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type App struct {
	// Application fields go here
	router http.Handler
}

func New() *App {
	dsn := os.Getenv("DATABASE_URL")
	// db, err := sql.Open("pgx", dsn)
	// err2 := db.PingContext(context.Background())
	// if err2 != nil {
	// 	log.Fatalf("❌ Database not reachable: %v", err)
	// }

	// db.SetMaxOpenConns(10)
	// db.SetMaxIdleConns(5)

	sqldb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithDSN(dsn),
	))

	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	// Initialize application and router here
	app := &App{
		// Initialize fields as needed
		router: loadRoutes(db),
	}
	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: a.router,
	}

	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start server %w", err)
	}

	return nil
}
