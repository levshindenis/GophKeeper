package cloud

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
)

func (c *Cloud) DeleteAccount(userId string) error {
	objectsCh := make(chan minio.ObjectInfo)

	go func() {
		defer close(objectsCh)
		for object := range c.Client.ListObjects(context.Background(), userId, minio.ListObjectsOptions{}) {
			if object.Err != nil {
				log.Fatalln(object.Err)
			}
			objectsCh <- object
		}
	}()

	opts := minio.RemoveObjectsOptions{
		GovernanceBypass: true,
	}

	for rErr := range c.Client.RemoveObjects(context.Background(), userId, objectsCh, opts) {
		return rErr.Err
	}

	err := c.Client.RemoveBucket(context.Background(), userId)
	if err != nil {
		return err
	}
	return nil
}
