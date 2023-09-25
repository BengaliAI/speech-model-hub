package infrastructure

import (
	"context"
	"speech-model-hub/internal/configs"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	internalDB *mongo.Database
}

func (db *MongoDB) GetCollection(collectionName string) *mongo.Collection {
	collection := db.internalDB.Collection(collectionName)
	return collection
}

func (db *MongoDB) CheckExistanceByFilter(collectionName string, filter bson.M) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := db.GetCollection(collectionName)
	res := collection.FindOne(ctx, filter)
	if res.Err() != nil {
		if string(res.Err().Error()) == "mongo: no documents in result" {
			return false, nil
		} else {
			return false, res.Err()
		}
	}
	return true, nil
}

func (db *MongoDB) InsertIntoDB(collectionName string, data interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := db.GetCollection(collectionName)
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (db *MongoDB) UpdateOneIntoDB(collectionName string, data bson.D, filter bson.D, opts *options.UpdateOptions) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := db.GetCollection(collectionName)
	_, err := collection.UpdateOne(ctx, filter, data, opts)
	if err != nil {
		return err
	}
	return nil
}

func (db *MongoDB) DeleteOneFromDB(collectionName string, filter bson.D) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := db.GetCollection(collectionName)
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (db *MongoDB) FindOneFromDB(collectionName string, filter bson.D) (bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := db.GetCollection(collectionName)
	var result bson.M
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (db *MongoDB) FindManyFromDB(collectionName string, filter bson.D) ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := db.GetCollection(collectionName)
	var result []bson.M
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (db *MongoDB) DoesCollectionExist(collectionName string) (bool, error) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	collections, err := db.internalDB.ListCollections(ctx, bson.M{})
	if err != nil {
		return false, err
	}
	for collections.Next(ctx) {
		var collection bson.M
		if err := collections.Decode(&collection); err != nil {
			return false, err
		}
		if collection["name"] == collectionName {
			return true, nil
		}
	}
	return false, nil
}

func (db *MongoDB) ReturnIfExists(collectionName string, filter bson.D, opts *options.UpdateOptions) (bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := db.GetCollection(collectionName)
	var result bson.M
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (db *MongoDB) ReturnLatest(collectionName string, filter bson.D, opts *options.FindOneOptions) (bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := db.GetCollection(collectionName)
	opts.SetSort(bson.D{{Key: "updated_at", Value: -1}})
	var value bson.M
	err := collection.FindOne(ctx, filter, opts).Decode(&value)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func setupUniques(db *MongoDB) {
	setUniqueIndex(db, "models", bson.D{{Key: "url", Value: 1}})
	setUniqueIndex(db, "models", bson.D{{Key: "display_name", Value: 1}})
	setUniqueIndex(db, "models", bson.D{{Key: "name", Value: 1}})
}

func NewDBInstance(cfg configs.AppConfig) *MongoDB {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoDB.MongoDBURL))
	if err != nil {
		panic(err)
	}
	internalDB := client.Database(cfg.MongoDB.DatabaseName)
	db := &MongoDB{internalDB: internalDB}
	setupUniques(db)
	return db
}
