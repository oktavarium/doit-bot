package db

import (
	"context"
	"fmt"

	"github.com/oktavarium/doit-bot/internal/server/internal/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	database   = "test"
	collection = "bot"
)

type storage struct {
	client *mongo.Client
}

func New(uri string) (*storage, error) {
	clientOpts := options.Client().ApplyURI(uri)

	ctx := context.Background()
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("new mongo client: %w", err)
	}

	dd := client.Database("test")
	dd.CreateCollection(ctx, "trainers")

	return &storage{
		client: client,
	}, nil
}

func (db *storage) CreateTask(
	ctx context.Context,
	owner int64,
	summary string,
	assignee *int64,
) (string, error) {
	collection := db.client.Database("test").Collection("trainers")
	result, err := collection.InsertOne(ctx, dbTask{
		Owner:    owner,
		Assignee: assignee,
		Summary:  &summary,
	})
	if err != nil {
		return "", fmt.Errorf("insert one: %w", err)
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("get inserted task id")
	}

	return id.Hex(), nil
}

func (db *storage) UpdateTask(ctx context.Context,
	owner int64,
	id string,
	assignee *int64,
	summary *string,
	done *bool,
) error {
	collection := db.client.Database("test").Collection("trainers")

	bsonId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}

	filter := bson.D{{"_id", bsonId}, {"owner", owner}}
	updatePayload := bson.D{}
	if assignee != nil {
		updatePayload = append(updatePayload, bson.E{"assignee", assignee})
	}
	if summary != nil {
		updatePayload = append(updatePayload, bson.E{"summary", summary})
	}
	if done != nil {
		updatePayload = append(updatePayload, bson.E{"done", done})
	}

	update := bson.D{{"$set", updatePayload}}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("update one: %w", err)
	}

	return nil
}

func (db *storage) GetTask(ctx context.Context, id string) (*dto.Task, error) {
	collection := db.client.Database("test").Collection("trainers")
	var result dbTask
	bsonId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}

	if err := collection.FindOne(ctx, bson.M{"_id": bsonId}).Decode(&result); err != nil {
		return nil, fmt.Errorf("find task: %w", err)
	}

	return dbTaskToDtoTask(result), nil
}

func (db *storage) DeleteTask(ctx context.Context, owner int64, id string) error {
	collection := db.client.Database("test").Collection("trainers")

	bsonId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}

	filter := bson.D{{"_id", bsonId}, {"owner", owner}}

	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("delete one: %w", err)
	}

	return nil
}

func (db *storage) GetTasks(ctx context.Context, owner int64) ([]*dto.Task, error) {
	collection := db.client.Database("test").Collection("trainers")

	cursor, err := collection.Find(ctx, bson.M{"owner": owner})
	if err != nil {
		return nil, fmt.Errorf("find tasks: %w", err)
	}

	var tasks []dbTask
	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, fmt.Errorf("read cursor: %w", err)
	}

	result := make([]*dto.Task, 0, len(tasks))
	for _, task := range tasks {
		result = append(result, dbTaskToDtoTask(task))
	}

	return result, nil
}
