package database

import (
	"context"
	"golang-clean-architecture/target/ent"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/cetnfurkan/core/config"
	"github.com/cetnfurkan/core/database"
)

func NewPostgresDatabase(cfg *config.Database) database.Database {
	return database.NewPostgresDatabase(cfg, createClient, database.WithCallback(migrate))
}

func createClient(driver *sql.Driver) *ent.Client {
	return ent.NewClient(ent.Driver(driver))
}

func migrate(client *ent.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return client.Schema.Create(ctx)
}
