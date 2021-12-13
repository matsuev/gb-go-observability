package app

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
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

func (as *AppStorage) FindAllUsers(ctx *gin.Context) (result *Users, err error) {
	span := trace.SpanFromContext(ctx.Request.Context())
	_, childSpan := span.TracerProvider().Tracer("TraceApp").Start(ctx.Request.Context(), "Storage")
	defer childSpan.End()

	childSpan.SetAttributes(attribute.Key("Collection").String("users"))
	childSpan.SetAttributes(attribute.Key("Method").String("GetUserByEmail"))

	filter := bson.M{}
	cursor, err := as.mdb.Collection("users").Find(ctx, filter)
	if err != nil {
		return
	}

	result = new(Users)
	err = cursor.All(ctx, result)
	return
}

func (as *AppStorage) GetUserByEmail(ctx context.Context, email string) (u *User, err error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.Key("collection").String("users"))
	span.SetAttributes(attribute.Key("method").String("GetUserByEmail"))

	u = new(User)
	err = as.mdb.Collection("users").FindOne(ctx, bson.M{}).Decode(u)
	return
}
