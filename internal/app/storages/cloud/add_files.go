package cloud

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"

	"github.com/levshindenis/GophKeeper/internal/app/models"
)

func (c *Cloud) AddFiles(userId string, binaries []models.CloudFile) error {
	fmt.Println("UserId: ", userId)
	for i := range binaries {
		if _, err := c.client.PutObject(
			context.Background(), userId, binaries[i].Filename, binaries[i].Data, binaries[i].Size,
			minio.PutObjectOptions{ContentType: "application/octet-stream"}); err != nil {
			return err
		}
	}
	return nil
}
