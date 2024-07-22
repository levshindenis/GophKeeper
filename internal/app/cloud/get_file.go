package cloud

import (
	"context"

	"github.com/minio/minio-go/v7"
)

func (c *Cloud) GetFile(userId string, name string) (*minio.Object, error) {
	object, err := c.Client.GetObject(context.Background(), userId, name, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer object.Close()

	return object, nil
}
