package app

// AppConfig struct
type AppConfig struct {
	ServiceName string
	Environment string
	InstanceID  int64
	ListenAddr  string
	JaegerURL   string
	MongoURL    string
	MongoDbName string
}

const (
	EnvProduction  = "production"
	EnvDevelopment = "development"
)

func CreateAppConfig() *AppConfig {
	return &AppConfig{
		ServiceName: "TracingApp",
		Environment: EnvProduction,
		InstanceID:  1,
		ListenAddr:  "localhost:8080",
		JaegerURL:   "http://localhost:14268/api/traces",
		MongoURL:    "mongodb://alex:pass@localhost",
		MongoDbName: "observability",
	}
}
