package db

import (
	"context"
	"os"

	authModels "github.com/will-kerwin/go-microservice-bookstore/auth/pkg/models"
	"github.com/will-kerwin/go-microservice-bookstore/pkg/models/events"
	"github.com/will-kerwin/go-microservice-bookstore/pkg/models/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func authenticatePassword(expected string, actual string) bool {
	return bcrypt.CompareHashAndPassword([]byte(expected), []byte(actual)) == nil
}

type MongoDbAuthRepository struct {
	client *mongo.Client
}

func NewAuthRepository(client *mongo.Client) *MongoDbAuthRepository {
	return &MongoDbAuthRepository{
		client: client,
	}
}

func (r *MongoDbAuthRepository) getCollection() *mongo.Collection {
	dbName := os.Getenv("DbName")

	return r.client.Database(dbName).Collection("users")
}

func (r *MongoDbAuthRepository) Add(ctx context.Context, user *events.CreateUserEvent) (*user.User, error) {

	hash, err := hashPassword(user.Password)

	if err != nil {
		return nil, err
	}

	userDoc := authModels.UserDocument{
		Username:  user.Username,
		Password:  hash,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	collection := r.getCollection()

	insertResult, err := collection.InsertOne(ctx, userDoc)

	if err != nil {
		return nil, err
	}

	userModel := userDoc.ToModel()
	userModel.ID = insertResult.InsertedID.(primitive.ObjectID).Hex()

	return userModel, nil

}

func (r *MongoDbAuthRepository) GetById(ctx context.Context, id string) (*user.User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var userDoc authModels.UserDocument
	collection := r.getCollection()

	filter := bson.M{"_id": objectId}

	err = collection.FindOne(ctx, filter).Decode(&userDoc)

	if err != nil {
		return nil, err
	}

	return userDoc.ToModel(), nil
}

func (r *MongoDbAuthRepository) Authenticate(ctx context.Context, username string, password string) (*user.User, error) {
	var userDoc authModels.UserDocument
	collection := r.getCollection()

	filter := bson.M{"username": username}

	err := collection.FindOne(ctx, filter).Decode(&userDoc)

	if err != nil {
		return nil, err
	}

	isAuthenticated := authenticatePassword(userDoc.Password, password)

	if !isAuthenticated {
		return nil, authModels.ErrUnauthenticated
	}

	return userDoc.ToModel(), nil

}

func (r *MongoDbAuthRepository) ValidateUsernameUnique(ctx context.Context, username string) (bool, error) {
	collection := r.getCollection()

	filter := bson.M{"username": username}

	count, err := collection.CountDocuments(ctx, filter)

	if err != nil {
		return false, err
	}

	return count == 0, nil
}
