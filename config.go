package common

import "github.com/Yulian302/lfusys-services-commons/jwt"

type Config struct {
	HTTPAddr string
	GRPCAddr string
	Env      string
	Tracing  bool
	*AWSConfig
	*jwt.JWTConfig
}

func LoadConfig() Config {
	httpAddr := EnvVar("HTTP_ADDR", ":8080")
	grpcAddr := EnvVar("GRPC_ADDR", "localhost:50051")
	env := EnvVar("ENV", "DEV")
	awsAccessKeyId := EnvVar("AWS_ACCESS_KEY_ID", "")
	awsSecretAccessKey := EnvVar("AWS_SECRET_ACCESS_KEY", "")
	awsRegion := EnvVar("AWS_REGION", "eu-north-1")
	jwtSecretKey := EnvVar("JWT_SECRET_KEY", "")

	return Config{
		HTTPAddr: httpAddr,
		GRPCAddr: grpcAddr,
		Env:      env,
		Tracing:  env == "DEV",
		AWSConfig: &AWSConfig{
			AWS_ACCESS_KEY_ID:     awsAccessKeyId,
			AWS_SECRET_ACCESS_KEY: awsSecretAccessKey,
			AWS_REGION:            awsRegion,
		},
		JWTConfig: &jwt.JWTConfig{
			SECRET_KEY: jwtSecretKey,
		},
	}
}
