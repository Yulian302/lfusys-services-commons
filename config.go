package common

import (
	"errors"

	"github.com/Yulian302/lfusys-services-commons/db"
	"github.com/Yulian302/lfusys-services-commons/jwt"
)

type ServiceConfig struct {
	UploadsURL                    string
	UploadsAddr                   string
	UploadsNotificationsQueueName string
	GatewayAddr                   string
	SessionGRPCUrl                string
	SessionGRPCAddr               string
}

type Config struct {
	Env     string
	Tracing bool

	*AWSConfig
	*jwt.JWTConfig
	*db.DynamoDBConfig
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

func (c *AWSConfig) Validate() error {
	err := c.ValidateSecrets()
	if err != nil {
		return err
	}

	if c.Region == "" {
		return errors.New("AWS_REGION_NAME is required")
	}

	if c.BucketName == "" {
		return errors.New("AWS_BUCKET is required")
	}

	return nil
}

func (c *AWSConfig) ValidateSecrets() error {
	if c.AccessKeyID == "" {
		return errors.New("AWS_ACCESS_KEY_ID is required")
	}
	if c.SecretAccessKey == "" {
		return errors.New("AWS_SECRET_ACCESS_KEY is required")
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

	gatewayAddr := EnvVar("GATEWAY_ADDR", ":8080")

	sessionGrpcAddr := EnvVar("SESSION_GRPC_ADDR", ":50051")
	sessionGrpcUrl := EnvVar("SESSION_GRPC_URL", "localhost:50051")

	uploadsUrl := EnvVar("UPLOADS_URL", "http://localhost:8081")
	uploadsAddr := EnvVar("UPLOADS_ADDR", "localhost:8080")

	jwtRefreshSecretKey := EnvVar("JWT_REFRESH_SECRET_KEY", "")
	jwtSecretKey := EnvVar("JWT_SECRET_KEY", "")

	awsAccessKeyId := EnvVar("AWS_ACCESS_KEY_ID", "")
	awsSecretAccessKey := EnvVar("AWS_SECRET_ACCESS_KEY", "")
	awsAccountId := EnvVar("AWS_ACCOUNT_ID", "")
	awsRegion := EnvVar("AWS_REGION", "eu-north-1")
	awsBucketName := EnvVar("AWS_BUCKET_NAME", "lfusyschunks")

	// Dynamo DB
	usersTableName := EnvVar("DYNAMODB_USERS_TABLE_NAME", "users")
	uploadsTableName := EnvVar("DYNAMODB_UPLOADS_TABLE_NAME", "uploads")
	filesTableName := EnvVar("DYNAMODB_FILES_TABLE_NAME", "files")

	// notifications queue
	uploadsNotificationsQueue := EnvVar("UPLOADS_NOTIFICATIONS_QUEUE_NAME", "uploads_notifications")

	return Config{
		Env:     env,
		Tracing: tracing,
		AWSConfig: &AWSConfig{
			AccessKeyID:     awsAccessKeyId,
			SecretAccessKey: awsSecretAccessKey,
			AccountID:       awsAccountId,
			Region:          awsRegion,
			BucketName:      awsBucketName,
		},
		JWTConfig: &jwt.JWTConfig{
			SecretKey:        jwtSecretKey,
			RefreshSecretKey: jwtRefreshSecretKey,
		},
		DynamoDBConfig: &db.DynamoDBConfig{
			UsersTableName:   usersTableName,
			UploadsTableName: uploadsTableName,
			FilesTableName:   filesTableName,
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
