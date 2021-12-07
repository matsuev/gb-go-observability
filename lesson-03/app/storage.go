package app

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

// AppStorage struct
type AppStorage struct {
	client *mongo.Client
	mdb    *mongo.Database
}

// CreateAppStorage function
func CreateAppStorage(cfg *AppConfig) (as *AppStorage, err error) {
	as = new(AppStorage)

	if as.client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.MongoURL)); err != nil {
		return
	}

	as.mdb = as.client.Database(cfg.MongoDbName)
	return
}

func (as *AppStorage) GetUserByEmail(ctx context.Context, email string) (u *User, err error) {
	tr := otel.Tracer("mongodb")
	_, span := tr.Start(ctx, "mongodb")
	span.SetAttributes(attribute.Key("collection").String("users"))
	span.SetAttributes(attribute.Key("method").String("GetUserByEmail"))
	defer span.End()

	u = new(User)
	err = as.mdb.Collection("users").FindOne(ctx, bson.M{}).Decode(u)
	return
}
