package config

import (
	"fmt"
	"strings"
)

type AWSConfig struct {
	Region          string
	AccountID       string
	BucketName      string
	AccessKeyID     string
	SecretAccessKey string
}

func (c *AWSConfig) Validate() error {
	var errs []string

	if c.Region == "" {
		errs = append(errs, "AWS_REGION_NAME is required")
	}

	if c.AccountID == "" {
		errs = append(errs, "AWS_ACCOUNT_ID is required")
	}

	if c.BucketName == "" {
		errs = append(errs, "AWS_BUCKET is required")
	}

	if len(errs) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errs, "; "))
	}

	if c.AccessKeyID == "" {
		fmt.Println("WARN: AWS_ACCESS_KEY_ID is not defined - will use default credentials chain")
	}

	if c.SecretAccessKey == "" {
		fmt.Println("WARN: AWS_SECRET_ACCESS_KEY is not defined - will use default credentials chain")
	}

	return nil
}
