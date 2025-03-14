package migrations

import (
	"context"

	"github.com/coreos/go-semver/semver"
	"go.mongodb.org/mongo-driver/mongo"
)

type migration interface {
	Version() semver.Version
	Update(ctx context.Context, db *mongo.Database) error
	Rollback(ctx context.Context, db *mongo.Database) error
}

func getAllMigrations() []migration {
	return []migration{
		migration000to001{},
	}
}

func Run(ctx context.Context, db *mongo.Database) error {
	// if db == nil {
	// 	return errors.New("empty collection")
	// }

	// for _, m := range getAllMigrations() {
	// 	if err := app.Update(ctx, db); err != nil {
	// 		return fmt.Errorf("run migration: %w", err)
	// 	}
	// }
	return nil
}
