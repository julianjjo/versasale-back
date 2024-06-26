package service

import (
	"context"

	model "github.com/julianjjo/versasale-back/internal/core/model"
	repository "github.com/julianjjo/versasale-back/internal/infrastructure/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SaveCustomer(client *mongo.Client, ctx context.Context, customer model.Customer) error {
	err := repository.SaveToMongoDB(client, ctx, "Customer", customer)
	if err != nil {
		return err
	}
	return nil
}

func EmailExists(client *mongo.Client, ctx context.Context, email string, collection string) (bool, error) {
	collectionRef := client.Database("versasale").Collection(collection)
	filter := bson.M{"email": email}

	var result bson.M
	err := collectionRef.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func DocumentIdExists(client *mongo.Client, ctx context.Context, documentId string, collection string) (bool, error) {
	collectionRef := client.Database("versasale").Collection(collection)
	filter := bson.M{"document_id": documentId}

	var result bson.M
	err := collectionRef.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
