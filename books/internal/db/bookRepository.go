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

type MongoDbBookRepository struct {
	client *mongo.Client
}

func NewBookRepository(client *mongo.Client) *MongoDbBookRepository {
	return &MongoDbBookRepository{
		client: client,
	}
}

func (r *MongoDbBookRepository) getCollection() *mongo.Collection {
	dbName := os.Getenv("DbName")
	return r.client.Database(dbName).Collection("books")
}

func (r *MongoDbBookRepository) Add(ctx context.Context, book *models.Book) (*models.Book, error) {

	authorId, err := primitive.ObjectIDFromHex(book.AuthorId)

	if err != nil {
		return nil, err
	}

	bookDoc := booksModels.BookDocument{
		Title:    book.Title,
		AuthorId: authorId,
		Synopsis: book.Synopsis,
		ImageUrl: book.ImageUrl,
		Genre:    book.Genre,
	}

	collection := r.getCollection()

	insertResult, err := collection.InsertOne(ctx, bookDoc)

	if err != nil {
		return nil, err
	}

	book.ID = insertResult.InsertedID.(primitive.ObjectID).Hex()

	return book, nil
}

func (r *MongoDbBookRepository) Update(ctx context.Context, id string, updateData *models.UpdateBookEventData) error {

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	collection := r.getCollection()

	updateFields := bson.M{}

	if updateData.Title != "" {
		updateFields["title"] = updateData.Title
	}

	if updateData.AuthorId != "" {
		authorOid, err := primitive.ObjectIDFromHex(updateData.AuthorId)

		if err != nil {
			updateFields["authorId"] = authorOid
		}
	}

	if updateData.Genre != "" {
		updateFields["genre"] = updateData.Genre
	}

	if updateData.Synopsis != "" {
		updateFields["synopsis"] = updateData.Synopsis
	}

	if updateData.ImageUrl != "" {
		updateFields["imageUrl"] = updateData.ImageUrl
	}

	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": updateFields}

	_, err = collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return err
	}

	return nil
}

func (r *MongoDbBookRepository) Get(ctx context.Context, title string, authorId string, genre string) ([]*models.Book, error) {
	var books []*models.Book = []*models.Book{}

	colllection := r.getCollection()

	filter := bson.M{}

	if title != "" {
		filter["title"] = bson.M{"$regex": title}
	}

	if genre != "" {
		filter["genre"] = genre
	}

	if authorId != "" {
		authorOid, err := primitive.ObjectIDFromHex(authorId)

		if err == nil {
			filter["authorId"] = authorOid
		}
	}

	cursor, err := colllection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var bookDoc booksModels.BookDocument

		if err := cursor.Decode(&bookDoc); err != nil {
			return nil, err
		}

		books = append(books, bookDoc.ToModel())
	}

	return books, nil
}

func (r *MongoDbBookRepository) GetById(ctx context.Context, id string) (*models.Book, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var bookDoc booksModels.BookDocument
	collection := r.getCollection()

	filter := bson.M{"_id": objectId}

	err = collection.FindOne(ctx, filter).Decode(&bookDoc)

	if err != nil {
		return nil, err
	}

	return bookDoc.ToModel(), nil
}

func (r *MongoDbBookRepository) Delete(ctx context.Context, id string) error {
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
