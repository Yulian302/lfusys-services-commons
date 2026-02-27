package config

import (
	"fmt"
	"strings"
)

type DynamoDBConfig struct {
	UsersTableName   string
	UploadsTableName string
	FilesTableName   string
}

func (c *DynamoDBConfig) Validate() error {
	var errs []string

	if c.UsersTableName == "" {
		errs = append(errs, "DYNAMODB_USERS_TABLE_NAME is required")
	}

	if c.UploadsTableName == "" {
		errs = append(errs, "DYNAMODB_UPLOADS_TABLE_NAME is required")
	}

	if c.FilesTableName == "" {
		errs = append(errs, "DYNAMODB_FILES_TABLE_NAME is required")
	}

	if len(errs) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errs, "; "))
	}

	return nil
}
