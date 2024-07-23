package cloud

import (
	"context"
	"github.com/minio/minio-go/v7"
	"path"
)

func (c *Cloud) GetFile(userId string, name string, dir string) error {
	if err := c.Client.FGetObject(context.Background(), userId, name, path.Join(dir, userId, name),
		minio.GetObjectOptions{}); err != nil {
		return err
	}
	return nil
}
