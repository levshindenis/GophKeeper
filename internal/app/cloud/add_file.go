package cloud

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/minio/minio-go/v7"
)

func (c *Cloud) AddFile(login string, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		return err
	}

	_, err = c.Client.PutObject(context.Background(), strings.ToLower(login), filepath.Base(filePath), file, fileStat.Size(),
		minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return err
	}
	return nil
}
