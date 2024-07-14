package cloud

import (
	"context"

	"github.com/minio/minio-go/v7"
)

func (c *Cloud) DeleteFiles(userId string, arr []string) error {
	for i := range arr {
		if err := c.Client.RemoveObject(context.Background(), userId, arr[i], minio.RemoveObjectOptions{}); err != nil {
			return err
		}
	}
	return nil
}
