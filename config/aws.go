package config

import "errors"

type AWSConfig struct {
	AccessKeyID     string
	SecretAccessKey string
	AccountID       string
	Region          string
	BucketName      string
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
