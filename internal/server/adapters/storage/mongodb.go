package storage

import (
	"context"
	"fmt"

	"github.com/oktavarium/doit-bot/internal/server/adapters/storage/migrations"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type db struct {
	client *mongo.Client
	users  *mongo.Collection
	groups *mongo.Collection
	tasks  *mongo.Collection
	lists  *mongo.Collection
}

func New(uri string) (*db, error) {
	clientOpts := options.Client().ApplyURI(uri)

	ctx := context.Background()
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("new mongo client: %w", err)
	}

	database := client.Database(database)
	if err := migrations.Run(ctx, database); err != nil {
		return nil, fmt.Errorf("run migrations: %w", err)
	}

	return &db{
		client: client,
		users:  database.Collection(usersCollection),
		groups: database.Collection(groupsCollection),
		tasks:  database.Collection(tasksCollection),
		lists:  database.Collection(listsCollection),
	}, nil
}
