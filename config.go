package common

import (
	"github.com/Yulian302/lfusys-services-commons/db"
	"github.com/Yulian302/lfusys-services-commons/jwt"
	"github.com/Yulian302/lfusys-services-commons/services"
)

type Service struct {
	*services.UploadsConfig
}

type Config struct {
	HTTPAddr         string
	GRPCAddr         string
	SessionsGRPCAddr string
	Env              string
	Tracing          bool
	UploadServiceUrl string
	*AWSConfig
	*jwt.JWTConfig
	*db.DynamoDBConfig
	*Service
}

func LoadConfig() Config {
	httpAddr := EnvVar("HTTP_ADDR", ":8080")
	grpcAddr := EnvVar("GRPC_ADDR", "localhost:50051")
	sessionsGrpcAddr := EnvVar("SESSIONS_GRPC_ADDR", "localhost:50051")
	env := EnvVar("ENV", "DEV")
	awsAccessKeyId := EnvVar("AWS_ACCESS_KEY_ID", "")
	awsSecretAccessKey := EnvVar("AWS_SECRET_ACCESS_KEY", "")
	awsRegion := EnvVar("AWS_REGION", "eu-north-1")
	jwtSecretKey := EnvVar("JWT_SECRET_KEY", "")
	uploadServiceUrl := EnvVar("UPLOAD_SERVICE_URL", "http://localhost:8080")
	dbUsersTableName := EnvVar("DYNAMODB_USERS_TABLE_NAME", "users")
	dbUploadsTableName := EnvVar("DYNAMODB_UPLOADS_TABLE_NAME", "uploads")
	uploadsConfig := EnvVar("UPLOADS_ADDR", "localhost:8080")
	awsBucketName := EnvVar("AWS_BUCKET_NAME", "lfusyschunks")

	return Config{
		HTTPAddr:         httpAddr,
		GRPCAddr:         grpcAddr,
		SessionsGRPCAddr: sessionsGrpcAddr,
		Env:              env,
		Tracing:          env == "DEV",
		UploadServiceUrl: uploadServiceUrl,
		AWSConfig: &AWSConfig{
			AWS_ACCESS_KEY_ID:     awsAccessKeyId,
			AWS_SECRET_ACCESS_KEY: awsSecretAccessKey,
			AWS_REGION:            awsRegion,
		},
		JWTConfig: &jwt.JWTConfig{
			SECRET_KEY: jwtSecretKey,
		},
		DynamoDBConfig: &db.DynamoDBConfig{
			DynamoDbUsersTableName:   dbUsersTableName,
			DynamoDbUploadsTableName: dbUploadsTableName,
		},
		Service: &Service{
			UploadsConfig: &services.UploadsConfig{
				UPLOADS_ADDR:    uploadsConfig,
				AWS_BUCKET_NAME: awsBucketName,
			},
		},
	}
}
