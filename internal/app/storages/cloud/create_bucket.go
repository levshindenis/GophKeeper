package cloud

import (
	"context"

	"github.com/minio/minio-go/v7"
)

func (c *Cloud) CreateBucket(userId string) error {
	err := c.client.MakeBucket(context.Background(),
		userId, minio.MakeBucketOptions{Region: "us-east-1", ObjectLocking: true})
	if err != nil {
		return err
	}
	return nil
}
