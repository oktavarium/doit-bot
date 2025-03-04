package migrations

import (
	"context"

	"github.com/coreos/go-semver/semver"
	"go.mongodb.org/mongo-driver/mongo"
)

type migration000to001 struct{}

func (m migration000to001) Version() semver.Version {
	return semver.Version{Major: 0, Minor: 0, Patch: 1}
}

func (m migration000to001) Update(
	ctx context.Context,
	db *mongo.Database,
) error {
	return nil
}

func (m migration000to001) Rollback(ctx context.Context, db *mongo.Database) error {
	return nil
}
