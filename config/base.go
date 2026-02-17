package config

import (
	"fmt"
)

type Config struct {
	Env         string
	Tracing     bool
	TracingAddr string
	FrontendURL string

	*CorsConfig

	*AWSConfig
	*JWTConfig
	*OAuthConfig
	*DynamoDBConfig
	*RedisConfig
	*ServiceConfig
}

func (c *Config) IsProduction() bool {
	return c.Env == "PROD"
}

func (c *Config) IsDevelopment() bool {
	return c.Env == "DEV"
}

func (c *Config) IsStaging() bool {
	return c.Env == "STAGING"
}

func (c *Config) Validate() error {

	err := c.AWSConfig.Validate()
	if err != nil {
		return err
	}
	return nil
}

func LoadConfig() Config {
	env := EnvVar("ENV", "DEV")
	tracingRaw := EnvVar("TRACING", "false")
	var tracing bool
	if tracingRaw == "true" {
		tracing = true
	}
	tracingAddr := EnvVar("TRACING_ADDR", "jaeger:4317")

	frontendURL := EnvVar("FRONTEND_URL", "http://localhost:3000")

	// cors
	allowedOrigins := EnvVar("CORS_ALLOWED_ORIGINS", "http://localhost:3000,http://127.0.0.1:3000")

	gatewayAddr := EnvVar("GATEWAY_ADDR", ":8080")

	sessionGrpcAddr := EnvVar("SESSION_GRPC_ADDR", ":50051")
	sessionGrpcUrl := EnvVar("SESSION_GRPC_URL", "localhost:50051")

	uploadsUrl := EnvVar("UPLOADS_URL", "http://localhost:8081/api/v1")
	uploadsAddr := EnvVar("UPLOADS_ADDR", ":8080")

	jwtRefreshSecretKey := EnvVar("JWT_REFRESH_SECRET_KEY", "")
	jwtSecretKey := EnvVar("JWT_SECRET_KEY", "")

	ghClientID := EnvVar("OAUTH2_GITHUB_CLIENT_ID", "")
	ghClientSecret := EnvVar("OAUTH2_GITHUB_CLIENT_SECRET", "")
	ghRedirectUri := EnvVar("OAUTH2_GITHUB_REDIRECT_URI", "")
	ghExchangeUrl := EnvVar("OAUTH2_GITHUB_EXCHANGE_URL", "")

	googleClientID := EnvVar("OAUTH2_GOOGLE_CLIENT_ID", "")
	googleClientSecret := EnvVar("OAUTH2_GOOGLE_CLIENT_SECRET", "")
	googleRedirectUri := EnvVar("OAUTH2_GOOGLE_REDIRECT_URI", "")
	googleExchangeUrl := EnvVar("OAUTH2_GOOGLE_EXCHANGE_URL", "")

	awsAccessKeyId := EnvVar("AWS_ACCESS_KEY_ID", "")
	awsSecretAccessKey := EnvVar("AWS_SECRET_ACCESS_KEY", "")
	awsAccountId := EnvVar("AWS_ACCOUNT_ID", "")
	awsRegion := EnvVar("AWS_REGION", "eu-north-1")
	awsBucketName := EnvVar("AWS_BUCKET_NAME", "lfusyschunks")

	// Dynamo DB
	usersTableName := EnvVar("DYNAMODB_USERS_TABLE_NAME", "users")
	uploadsTableName := EnvVar("DYNAMODB_UPLOADS_TABLE_NAME", "uploads")
	filesTableName := EnvVar("DYNAMODB_FILES_TABLE_NAME", "files")

	redisHost := EnvVar("REDIS_HOST", "redis:6379")

	// notifications queue
	uploadsNotificationsQueue := EnvVar("UPLOADS_NOTIFICATIONS_QUEUE_NAME", "uploads_notifications")

	return Config{
		Env:         env,
		Tracing:     tracing,
		TracingAddr: tracingAddr,
		FrontendURL: frontendURL,
		CorsConfig: &CorsConfig{
			Origins: allowedOrigins,
		},
		AWSConfig: &AWSConfig{
			AccessKeyID:     awsAccessKeyId,
			SecretAccessKey: awsSecretAccessKey,
			AccountID:       awsAccountId,
			Region:          awsRegion,
			BucketName:      awsBucketName,
		},
		JWTConfig: &JWTConfig{
			SecretKey:        jwtSecretKey,
			RefreshSecretKey: jwtRefreshSecretKey,
		},
		OAuthConfig: &OAuthConfig{
			GithubConfig: &GithubConfig{
				ClientID:     ghClientID,
				ClientSecret: ghClientSecret,
				RedirectURI:  ghRedirectUri,
				ExchangeURL:  ghExchangeUrl,
			},
			GoogleConfig: &GoogleConfig{
				ClientID:     googleClientID,
				ClientSecret: googleClientSecret,
				RedirectURI:  googleRedirectUri,
				ExchangeURL:  googleExchangeUrl,
			},
		},
		DynamoDBConfig: &DynamoDBConfig{
			UsersTableName:   usersTableName,
			UploadsTableName: uploadsTableName,
			FilesTableName:   filesTableName,
		},
		RedisConfig: &RedisConfig{
			HOST: redisHost,
		},
		ServiceConfig: &ServiceConfig{
			UploadsURL:                    uploadsUrl,
			GatewayAddr:                   gatewayAddr,
			UploadsAddr:                   uploadsAddr,
			SessionGRPCUrl:                sessionGrpcUrl,
			SessionGRPCAddr:               sessionGrpcAddr,
			UploadsNotificationsQueueName: uploadsNotificationsQueue,
		},
	}
}

func (c *Config) ValidateAllSecrets() error {
	if c.Env == "" {
		return fmt.Errorf("ENV is required")
	}

	if err := c.CorsConfig.Validate(); err != nil {
		return fmt.Errorf("CORS was not configured: %s", err.Error())
	}

	if err := c.AWSConfig.Validate(); err != nil {
		return fmt.Errorf("AWS was not configured: %s", err.Error())
	}

	if err := c.JWTConfig.ValidateSecrets(); err != nil {
		return fmt.Errorf("JWT was not configured: %s", err.Error())
	}

	if err := c.OAuthConfig.GithubConfig.ValidateSecrets(); err != nil {
		return fmt.Errorf("GITHUB OAUTH was not configured: %s", err.Error())
	}

	if err := c.OAuthConfig.GoogleConfig.ValidateSecrets(); err != nil {
		return fmt.Errorf("GOOGLE OAUTH was not configured: %s", err.Error())
	}

	if err := c.DynamoDBConfig.Validate(); err != nil {
		return fmt.Errorf("DYNAMODB was not configured: %s", err.Error())
	}

	if err := c.RedisConfig.ValidateSecrets(); err != nil {
		return fmt.Errorf("REDIS was not configured: %s", err.Error())
	}

	if err := c.ServiceConfig.Validate(); err != nil {
		return fmt.Errorf("SERVICES were not configured: %s", err.Error())
	}

	return nil
}
