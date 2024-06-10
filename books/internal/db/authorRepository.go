package db

import (
	"context"
	"os"

	booksModels "github.com/will-kerwin/go-microservice-bookstore/books/pkg/models"
	"github.com/will-kerwin/go-microservice-bookstore/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDbAuthorRepository struct {
	client *mongo.Client
}

func NewAuthorRepository(client *mongo.Client) *MongoDbAuthorRepository {
	return &MongoDbAuthorRepository{
		client: client,
	}
}

func (r *MongoDbAuthorRepository) getCollection() *mongo.Collection {
	dbName := os.Getenv("DbName")
	return r.client.Database(dbName).Collection("authors")
}

func (r *MongoDbAuthorRepository) Add(ctx context.Context, author *models.Author) (*models.Author, error) {
	authorDoc := booksModels.AuthorDocument{
		Name:        author.Name,
		DateOfBirth: primitive.NewDateTimeFromTime(author.DateOfBirth),
	}

	collection := r.getCollection()

	insertResult, err := collection.InsertOne(ctx, authorDoc)

	if err != nil {
		return nil, err
	}

	author.ID = insertResult.InsertedID.(primitive.ObjectID).Hex()

	return author, nil
}

func (r *MongoDbAuthorRepository) Get(ctx context.Context) ([]*models.Author, error) {
	var authors []*models.Author = []*models.Author{}

	collection := r.getCollection()

	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var authorDoc booksModels.AuthorDocument

		if err := cursor.Decode(&authorDoc); err != nil {
			return nil, err
		}

		authors = append(authors, authorDoc.ToModel())
	}

	return authors, nil
}

func (r *MongoDbAuthorRepository) GetById(ctx context.Context, id string) (*models.Author, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var authorDoc booksModels.AuthorDocument
	collection := r.getCollection()

	filter := bson.M{"_id": objectId}

	err = collection.FindOne(ctx, filter).Decode(&authorDoc)

	if err != nil {
		return nil, err
	}

	return authorDoc.ToModel(), nil
}

func (r *MongoDbAuthorRepository) Delete(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	collection := r.getCollection()

	filter := bson.M{"_id": objectId}

	_, err = collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	return nil
}
