package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collection = "надо выделить коллекции и для каждой из них свои методы?"
)

type storage struct {
	client *mongo.Client
	users  *mongo.Collection
	groups *mongo.Collection
	tasks  *mongo.Collection
	lists  *mongo.Collection
}

func New(uri string) (*storage, error) {
	clientOpts := options.Client().ApplyURI(uri)

	ctx := context.Background()
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("new mongo client: %w", err)
	}

	db := client.Database(database)
	return &storage{
		client: client,
		users:  db.Collection(usersCollection),
		groups: db.Collection(groupsCollection),
		tasks:  db.Collection(tasksCollection),
		lists:  db.Collection(listsCollection),
	}, nil
}
