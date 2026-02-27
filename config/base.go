package config

import (
	"fmt"
	"strconv"
)

type ServiceName string

const (
	Gateway  ServiceName = "gateway"
	Sessions ServiceName = "sessions"
	Uploads  ServiceName = "uploads"
)

type ConfigOptions struct {
	LoadCors     bool
	LoadAWS      bool
	LoadJwtAuth  bool
	LoadOAuth    bool
	LoadDynamoDB bool
	LoadRedis    bool
	LoadSqs      bool
}

type Config struct {
	Env         Environment
	Tracing     bool
	TracingAddr string

	Cors CorsConfig

	AWS      AWSConfig
	JWT      JWTConfig
	OAuth    OAuthConfig
	DynamoDB DynamoDBConfig
	Sqs      SQSConfig
	Redis    RedisConfig
	Service  ServiceConfig
}

type ConfigBuilder struct {
	config *Config
	errors []error
}

func NewConfigBuilder() *ConfigBuilder {
	return &ConfigBuilder{
		config: &Config{},
		errors: []error{},
	}
}

func (c *Config) IsProduction() bool {
	return c.Env == EnvProduction
}

func (c *Config) IsDevelopment() bool {
	return c.Env == EnvDevelopment
}

func (c *Config) IsStaging() bool {
	return c.Env == EnvStaging
}

func (c *Config) IsTest() bool {
	return c.Env == EnvTest
}

func (c *Config) ValidateEnv() error {
	switch c.Env {
	case EnvProduction, EnvDevelopment, EnvStaging, EnvTest:
		return nil
	default:
		return fmt.Errorf("invalid environment %s", c.Env)
	}
}

func (b *ConfigBuilder) WithDevTools() *ConfigBuilder {
	tracing, _ := strconv.ParseBool(EnvVar("TRACING", "false"))
	b.config.Tracing = tracing
	b.config.TracingAddr = EnvVar("TRACING_ADDR", "jaeger:4317")

	return b
}

func (b *ConfigBuilder) WithCors() *ConfigBuilder {
	allowCredentials, _ := strconv.ParseBool(EnvVar("CORS_ALLOW_CREDENTIALS", ""))
	corsCfg := CorsConfig{
		Origins:     EnvVar("CORS_ALLOW_ORIGINS", ""),
		Headers:     EnvVar("CORS_ALLOW_HEADERS", ""),
		Methods:     EnvVar("CORS_ALLOW_METHODS", ""),
		Credentials: allowCredentials,
	}

	if err := corsCfg.Validate(); err != nil {
		b.errors = append(b.errors, fmt.Errorf("CORS: %w", err))
		return b
	}

	b.config.Cors = corsCfg
	return b
}

func (b *ConfigBuilder) WithAws() *ConfigBuilder {
	awsCfg := AWSConfig{
		AccessKeyID:     EnvVar("AWS_ACCESS_KEY_ID", ""),
		SecretAccessKey: EnvVar("AWS_SECRET_ACCESS_KEY", ""),
		AccountID:       EnvVar("AWS_ACCOUNT_ID", ""),
		Region:          EnvVar("AWS_REGION", ""),
		BucketName:      EnvVar("AWS_BUCKET_NAME", ""),
	}

	if err := awsCfg.Validate(); err != nil {
		b.errors = append(b.errors, fmt.Errorf("AWS: %w", err))
		return b
	}

	b.config.AWS = awsCfg
	return b
}

func (b *ConfigBuilder) WithOAuth() *ConfigBuilder {
	oAuthCfg := OAuthConfig{
		Github: GithubConfig{
			ClientID:     EnvVar("OAUTH2_GITHUB_CLIENT_ID", ""),
			ClientSecret: EnvVar("OAUTH2_GITHUB_CLIENT_SECRET", ""),
			RedirectURI:  EnvVar("OAUTH2_GITHUB_REDIRECT_URI", ""),
			ExchangeURL:  EnvVar("OAUTH2_GITHUB_EXCHANGE_URL", ""),
		},
		Google: GoogleConfig{
			ClientID:     EnvVar("OAUTH2_GOOGLE_CLIENT_ID", ""),
			ClientSecret: EnvVar("OAUTH2_GOOGLE_CLIENT_SECRET", ""),
			RedirectURI:  EnvVar("OAUTH2_GOOGLE_REDIRECT_URI", ""),
			ExchangeURL:  EnvVar("OAUTH2_GOOGLE_EXCHANGE_URL", ""),
		},
	}

	if err := oAuthCfg.Github.ValidateSecrets(); err != nil {
		b.errors = append(b.errors, fmt.Errorf("OAUTH GITHUB: %w", err))
		return b
	}

	if err := oAuthCfg.Google.ValidateSecrets(); err != nil {
		b.errors = append(b.errors, fmt.Errorf("OAUTH GOOGLE: %w", err))
		return b
	}

	b.config.OAuth = oAuthCfg
	return b
}

func (b *ConfigBuilder) WithJWTAuth() *ConfigBuilder {
	jwtCfg := JWTConfig{

		RefreshSecretKey: EnvVar("JWT_REFRESH_SECRET_KEY", ""),
		SecretKey:        EnvVar("JWT_SECRET_KEY", ""),
	}

	if err := jwtCfg.ValidateSecrets(); err != nil {
		b.errors = append(b.errors, fmt.Errorf("JWT: %w", err))
		return b
	}

	b.config.JWT = jwtCfg
	return b
}

func (b *ConfigBuilder) WithDynamoDB() *ConfigBuilder {
	dbCfg := DynamoDBConfig{
		UsersTableName:   EnvVar("DYNAMODB_USERS_TABLE_NAME", ""),
		UploadsTableName: EnvVar("DYNAMODB_UPLOADS_TABLE_NAME", ""),
		FilesTableName:   EnvVar("DYNAMODB_FILES_TABLE_NAME", ""),
	}

	if err := dbCfg.Validate(); err != nil {
		b.errors = append(b.errors, fmt.Errorf("DynamoDB %w", err))
		return b
	}

	b.config.DynamoDB = dbCfg
	return b
}

func (b *ConfigBuilder) WithSQS() *ConfigBuilder {
	sqsCfg := SQSConfig{
		QueueName: EnvVar("UPLOADS_NOTIFICATIONS_QUEUE_NAME", ""),
	}

	if err := sqsCfg.Validate(); err != nil {
		b.errors = append(b.errors, fmt.Errorf("SQS %w", err))
		return b
	}

	b.config.Sqs = sqsCfg
	return b
}

func (b *ConfigBuilder) WithRedis() *ConfigBuilder {
	redisCfg := RedisConfig{
		HOST: EnvVar("REDIS_HOST", ""),
	}

	if err := redisCfg.Validate(); err != nil {
		b.errors = append(b.errors, fmt.Errorf("Redis %w", err))
		return b
	}

	b.config.Redis = redisCfg
	return b
}

func (b *ConfigBuilder) WithService(serviceName ServiceName) *ConfigBuilder {
	switch serviceName {
	case Gateway:
		b.WithGateway()
	case Sessions:
		b.WithSessions()
	case Uploads:
		b.WithUploads()
	default:
		b.errors = append(b.errors, fmt.Errorf("unknown service name: %s", serviceName))
	}

	return b
}

func (b *ConfigBuilder) WithGateway() *ConfigBuilder {
	cfg := GatewayConfig{
		Addr:            EnvVar("GATEWAY_ADDR", ""),
		SessionsGRPCUrl: EnvVar("SESSIONS_GRPC_URL", ""),
		FrontendUrl:     EnvVar("FRONTEND_URL", ""),
	}

	if err := cfg.Validate(); err != nil {
		b.errors = append(b.errors, fmt.Errorf("Gateway %w", err))
		return b
	}
	b.config.Service.Gateway = cfg
	return b
}

func (b *ConfigBuilder) WithSessions() *ConfigBuilder {
	cfg := SessionsConfig{
		Addr: EnvVar("SESSIONS_GRPC_ADDR", ""),
	}

	if err := cfg.Validate(); err != nil {
		b.errors = append(b.errors, fmt.Errorf("Sessions %w", err))
		return b
	}
	b.config.Service.Sessions = cfg
	return b
}

func (b *ConfigBuilder) WithUploads() *ConfigBuilder {
	cfg := UploadsConfig{
		Addr:        EnvVar("UPLOADS_ADDR", ""),
		FrontendUrl: EnvVar("FRONTEND_URL", ""),
	}

	if err := cfg.Validate(); err != nil {
		b.errors = append(b.errors, fmt.Errorf("Uploads %w", err))
		return b
	}
	b.config.Service.Uploads = cfg
	return b
}

func (b *ConfigBuilder) Build() (Config, error) {
	if len(b.errors) > 0 {
		return Config{}, fmt.Errorf("config errors: %v", b.errors)
	}
	return *b.config, nil
}

func LoadConfig(opts ConfigOptions, serviceName ServiceName) (Config, error) {
	b := NewConfigBuilder()

	b.config.Env = ParseEnvironment(EnvVar("ENV", string(EnvDevelopment)))

	// service-specific
	b.WithService(serviceName)

	if b.config.IsDevelopment() {
		b.WithDevTools()
	}

	if opts.LoadCors {
		b.WithCors()
	}

	if opts.LoadAWS {
		b.WithAws()
	}

	if opts.LoadJwtAuth {
		b.WithJWTAuth()
	}

	if opts.LoadOAuth {
		b.WithOAuth()
	}

	if opts.LoadDynamoDB {
		b.WithDynamoDB()
	}

	if opts.LoadSqs {
		b.WithSQS()
	}

	if opts.LoadRedis {
		b.WithRedis()
	}

	return b.Build()
}
