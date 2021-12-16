package app

import "time"

// AppConfig struct
type AppConfig struct {
	ServiceName    string
	Environment    string
	InstanceID     int64
	ListenAddr     string
	JaegerURL      string
	MongoURL       string
	MongoDbName    string
	RedisServer    string
	RedisPassword  string
	RedisTTL       time.Duration
	LocalCacheSize int
	LocalCacheTTL  time.Duration
}

const (
	EnvProduction  = "production"
	EnvDevelopment = "development"
)

func CreateAppConfig() *AppConfig {
	return &AppConfig{
		ServiceName:    "TracingApp",
		Environment:    EnvProduction,
		InstanceID:     1,
		ListenAddr:     ":8080",
		JaegerURL:      "http://localhost:14268/api/traces",
		MongoURL:       "mongodb://alex:pass@localhost",
		MongoDbName:    "observability",
		RedisServer:    "localhost:6379",
		RedisPassword:  "",
		RedisTTL:       time.Hour,
		LocalCacheSize: 10000,
		LocalCacheTTL:  10 * time.Minute,
	}
}
