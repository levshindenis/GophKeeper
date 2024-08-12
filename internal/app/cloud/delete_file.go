package cloud

import (
	"context"
	"path/filepath"

	"github.com/minio/minio-go/v7"
)

func (c *Cloud) DeleteFile(userId string, filePath string) error {
	if err := c.Client.RemoveObject(context.Background(), userId, filepath.Base(filePath), minio.RemoveObjectOptions{}); err != nil {
		return err
	}
	return nil
}
