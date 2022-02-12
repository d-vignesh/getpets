package db

import (
	"context"
	"os"

	"github.com/d-vignesh/getpets/pkg/domain"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mongoStore implements domain.PetDB with an mongo storage
type mongoStore struct {
	Client *mongo.Client
}

func NewMongoStore() (domain.PetDB, error) {
	uri := os.Getenv("mongo_uri")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return mongoStore{Client: client}, nil
}

func (ms mongoStore) Get(id uuid.UUID) (*domain.Pet, error) {
	col := ms.Client.Database("sellpetdb").Collection("pets")
	filter := bson.M{"id": id}
	pet := domain.Pet{}
	err := col.FindOne(context.TODO(), filter).Decode(&pet)
	if err != nil {
		return nil, err
	}
	return &pet, nil
}

func (ms mongoStore) List(category string) ([]*domain.Pet, error) {
	col := ms.Client.Database("sellpetdb").Collection("pets")
	filter := bson.M{"category": category}
	pets := []*domain.Pet{}
	cursor, err := col.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	if err := cursor.All(context.TODO(), &pets); err != nil {
		return nil, err
	}
	return pets, nil
}

func (ms mongoStore) Create(pet *domain.Pet) error {
	col := ms.Client.Database("sellpetdb").Collection("pets")
	_, err := col.InsertOne(context.TODO(), pet)
	return err
}

func (ms mongoStore) Delete(id uuid.UUID) error {
	col := ms.Client.Database("sellpetdb").Collection("pets")
	filter := bson.M{"id": id}
	_, err := col.DeleteOne(context.TODO(), filter)
	return err
}
