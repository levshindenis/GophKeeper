package cloud

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/stretchr/testify/assert"

	"github.com/levshindenis/GophKeeper/internal/app/tools"
)

var (
	cloud     Cloud
	bucketArr []string
)

func TestMain(m *testing.M) {
	endpoint := "play.min.io"
	accessKeyID := "Q3AM3UQ867SPQQA43P2F"
	secretAccessKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	useSSL := true

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}

	cloud = Cloud{Client: client}

	f, err := os.Create("../tools/1.txt")
	if err != nil {
		log.Fatalf(err.Error())
	}
	if _, err = f.WriteString("Hello world!"); err != nil {
		log.Fatalf(err.Error())
	}
	f.Close()

	f, err = os.Create("../tools/2.txt")
	if err != nil {
		log.Fatalf(err.Error())
	}
	if _, err = f.WriteString("Goodbye!"); err != nil {
		log.Fatalf(err.Error())
	}
	f.Close()

	exitCode := m.Run()
	if err = os.RemoveAll("gifgijfngk"); err != nil {
		log.Fatalf(err.Error())
	}
	if err = os.RemoveAll(bucketArr[0]); err != nil {
		log.Fatalf(err.Error())
	}
	if err = os.Remove("../tools/1.txt"); err != nil {
		log.Fatalf(err.Error())
	}
	if err = os.Remove("../tools/2.txt"); err != nil {
		log.Fatalf(err.Error())
	}

	os.Exit(exitCode)
}

func TestCloud_CreateBucket(t *testing.T) {
	tests := []struct {
		name    string
		login   string
		wantErr bool
	}{
		{
			name:    "Good create bucket",
			login:   strings.ToLower(tools.GenerateShortKey(false)),
			wantErr: false,
		},
		{
			name:    "Bad create bucket",
			login:   tools.GenerateShortKey(true),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := cloud.CreateBucket(tt.login); (err != nil) != tt.wantErr {
				t.Errorf("CreateBucket() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				bucketArr = append(bucketArr, tt.login)
			}
		})
	}
}

func TestCloud_AddFile(t *testing.T) {
	tests := []struct {
		name     string
		login    string
		filePath string
		wantErr  bool
	}{
		{
			name:     "Good add file_1",
			login:    bucketArr[0],
			filePath: "../tools/1.txt",
			wantErr:  false,
		},
		{
			name:     "Good add file_2",
			login:    bucketArr[0],
			filePath: "../tools/2.txt",
			wantErr:  false,
		},
		{
			name:     "Bad login",
			login:    "gifgijfngk",
			filePath: "../tools/1.txt",
			wantErr:  true,
		},
		{
			name:     "Bad filePath",
			login:    bucketArr[0],
			filePath: "../tools/3.txt",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := cloud.AddFile(tt.login, tt.filePath)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestCloud_GetFile(t *testing.T) {
	tests := []struct {
		name     string
		login    string
		filePath string
		wantErr  bool
	}{
		{
			name:     "Good get",
			login:    bucketArr[0],
			filePath: "../tools/2.txt",
			wantErr:  false,
		},
		{
			name:     "Bad login",
			login:    "gifgijfngk",
			filePath: "../tools/2.txt",
			wantErr:  true,
		},
		{
			name:     "Bad fileName",
			login:    bucketArr[0],
			filePath: "../tools/32.txt",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := cloud.GetFile(tt.login, filepath.Base(tt.filePath), ".")
			assert.Equal(t, tt.wantErr, err != nil)
			_, err1 := os.Stat(path.Join(".", tt.login, filepath.Base(tt.filePath)))
			assert.Equal(t, tt.wantErr, err1 != nil)
		})
	}
}

func TestCloud_DeleteFile(t *testing.T) {
	tests := []struct {
		name     string
		login    string
		filePath string
		wantErr  bool
	}{
		{
			name:     "Good delete",
			login:    bucketArr[0],
			filePath: "../tools/2.txt",
			wantErr:  false,
		},
		{
			name:     "Bad login",
			login:    "gifgijfngk",
			filePath: "../tools/2.txt",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := cloud.DeleteFile(tt.login, tt.filePath)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestCloud_DeleteAccount(t *testing.T) {
	tests := []struct {
		name    string
		login   string
		wantErr bool
	}{
		{
			name:    "Good delete",
			login:   bucketArr[0],
			wantErr: false,
		},
		{
			name:    "Bad delete",
			login:   "pqmfojtnkfmokgms",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := cloud.DeleteAccount(tt.login)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
