package config

import "errors"

type DynamoDBConfig struct {
	UsersTableName   string
	UploadsTableName string
	FilesTableName   string
}

func (c *DynamoDBConfig) Validate() error {
	if c.UsersTableName == "" {
		return errors.New("USERS_TABLE_NAME is required")
	}

	if c.UploadsTableName == "" {
		return errors.New("UPLOADS_TABLE_NAME is required")
	}

	if c.FilesTableName == "" {
		return errors.New("FILES_TABLE_NAME is required")
	}

	return nil
}
