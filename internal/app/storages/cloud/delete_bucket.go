package cloud

import (
	"context"
)

func (c *Cloud) DeleteBucket(userId string) error {
	err := c.Client.RemoveBucket(context.Background(), userId)
	if err != nil {
		return err
	}
	return nil
}
