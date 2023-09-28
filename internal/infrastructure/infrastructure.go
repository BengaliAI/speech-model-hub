package infrastructure

import (
	"context"
	"speech-model-hub/internal/configs"
	"speech-model-hub/internal/domains"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"infrastructure",
	fx.Provide(
		NewDBInstance,
		fx.Annotate(
			NewModelRepository,
			fx.As(new(domains.IFModelRepository)),
		),
	),
)

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
