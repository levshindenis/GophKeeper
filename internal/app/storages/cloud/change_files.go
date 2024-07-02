package cloud

import (
	"context"

	"github.com/minio/minio-go/v7"

	"github.com/levshindenis/GophKeeper/internal/app/models"
)

func (c *Cloud) ChangeFiles(userId string, binaries []models.ChCloudFile) error {
	for i := range binaries {
		if err := c.client.RemoveObject(context.Background(),
			userId, binaries[i].OldFilename, minio.RemoveObjectOptions{}); err != nil {
			return err
		}

		if _, err := c.client.PutObject(
			context.Background(), userId, binaries[i].NewFilename, binaries[i].NewData, binaries[i].NewSize,
			minio.PutObjectOptions{ContentType: "application/octet-stream"}); err != nil {
			return err
		}
	}
	return nil
}
