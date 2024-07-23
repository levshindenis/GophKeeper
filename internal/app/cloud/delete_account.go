package cloud

import (
	"context"
	"github.com/minio/minio-go/v7"
)

func (c *Cloud) DeleteAccount(userId string) error {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	objectCh := c.Client.ListObjects(ctx, userId, minio.ListObjectsOptions{})

	for elem := range objectCh {
		err := c.Client.RemoveObject(context.Background(), userId, elem.Key, minio.RemoveObjectOptions{GovernanceBypass: true})
		if err != nil {
			return err
		}
	}

	err := c.Client.RemoveBucket(context.Background(), userId)
	if err != nil {
		return err
	}
	return nil
}
