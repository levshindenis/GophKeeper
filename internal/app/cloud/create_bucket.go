package cloud

import (
	"context"
	"strings"

	"github.com/minio/minio-go/v7"
)

func (c *Cloud) CreateBucket(login string) error {
	err := c.Client.MakeBucket(context.Background(),
		strings.ToLower(login), minio.MakeBucketOptions{Region: "us-east-1", ObjectLocking: true})
	if err != nil {
		return err
	}
	return nil
}
