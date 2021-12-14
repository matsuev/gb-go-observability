package app

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (as *AppStorage) FindAllUsers(ctx *gin.Context, result *Users) (err error) {
	span := as.CreateSpan(ctx.Request.Context(), "users", "FindAllUsers")
	defer span.End()

	// if result in cache, then return
	if err = as.cache.Get(ctx, "users:all", &result); err == nil {
		return
	}

	// else get result from database
	filter := bson.M{}
	cursor, err := as.mdb.Collection("users").Find(ctx, filter)
	if err != nil {
		return
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, result); err != nil {
		return
	}

	// store result to cache and return
	as.cache.Set(ctx, "users:all", result)

	return
}

// GetUserByEmail fuction
func (as *AppStorage) GetUserByEmail(ctx *gin.Context, result *User) (err error) {
	span := as.CreateSpan(ctx.Request.Context(), "users", "GetUserByEmail")
	defer span.End()

	// if result in cache, then return
	if err = as.cache.Get(ctx, "user:email:"+result.Email, &result); err == nil {
		return
	}

	filter := bson.D{
		primitive.E{
			Key:   "email",
			Value: result.Email,
		},
	}
	if err = as.mdb.Collection("users").FindOne(ctx, filter).Decode(result); err != nil {
		return
	}

	// store result to cache and return
	as.cache.Set(ctx, "user:email:"+result.Email, result)

	return
}

// GetUserByUid function
func (as *AppStorage) GetUserByUid(ctx *gin.Context, user *User) (err error) {
	span := as.CreateSpan(ctx.Request.Context(), "users", "GetUserByUid")
	defer span.End()

	// if result in cache, then return
	if err = as.cache.Get(ctx, "user:uid:"+user.UID, user); err == nil {
		return
	}

	filter := bson.D{
		primitive.E{
			Key:   "uid",
			Value: user.UID,
		},
	}

	if err = as.mdb.Collection("users").FindOne(ctx, filter).Decode(user); err != nil {
		return
	}

	// store result to cache and return
	as.cache.Set(ctx, "user:uid:"+user.UID, user)

	return
}
