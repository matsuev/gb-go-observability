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
	cache  *AppCache
}

// CreateAppStorage function
func CreateAppStorage(cfg *AppConfig) (as *AppStorage, err error) {
	as = new(AppStorage)

	if as.client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.MongoURL)); err != nil {
		return
	}

	as.cache = CreateAppCache(cfg)

	as.mdb = as.client.Database(cfg.MongoDbName)
	return
}

// CreateSpan function
func (as *AppStorage) CreateSpan(ctx context.Context, coll string, method string) (sp trace.Span) {
	span := trace.SpanFromContext(ctx)
	_, childSpan := span.TracerProvider().Tracer("TraceApp").Start(ctx, "Storage")
	defer childSpan.End()

	childSpan.SetAttributes(attribute.Key("Collection").String(coll))
	childSpan.SetAttributes(attribute.Key("Method").String(method))
	return childSpan
}

func (as *AppStorage) FindAllUsers(ctx *gin.Context) (result *Users, err error) {
	span := as.CreateSpan(ctx.Request.Context(), "users", "FindAllUsers")
	defer span.End()

	result = new(Users)

	// if result in cache, then return
	if err = as.cache.Get(ctx, ctx.Request.RequestURI, &result); err == nil {
		return
	}

	// else get result from database
	filter := bson.M{}
	cursor, err := as.mdb.Collection("users").Find(ctx, filter)
	if err != nil {
		return
	}

	if err = cursor.All(ctx, result); err != nil {
		return
	}

	// store result to cache and return
	as.cache.Set(ctx, ctx.Request.RequestURI, result)

	return
}

func (as *AppStorage) GetUserByEmail(ctx *gin.Context, email string) (u *User, err error) {
	span := as.CreateSpan(ctx.Request.Context(), "users", "GetUserByEmail")
	defer span.End()

	u = new(User)
	err = as.mdb.Collection("users").FindOne(ctx, bson.M{}).Decode(u)
	return
}
