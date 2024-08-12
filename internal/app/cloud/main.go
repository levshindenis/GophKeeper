package cloud

import "github.com/minio/minio-go/v7"

type Cloud struct {
	Client *minio.Client
}

func (c *Cloud) GetClient() *minio.Client {
	return c.Client
}
