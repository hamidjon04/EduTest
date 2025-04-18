package minio

import (
	"context"
	"edutest/pkg/config"
	"fmt"
	"log/slog"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Minio interface {
	UploadFile(file multipart.File, fileName string) (string, error)
}

type MinioStruct struct {
	Client          *minio.Client
	BucketName      string
	PublicHost      string
	Log             *slog.Logger
}

func MinioConnect(minio MinioStruct, log *slog.Logger, cfg config.Config) Minio {
	return &MinioStruct{
		Client:     minio.Client,
		BucketName: minio.BucketName,
		PublicHost: cfg.MINIO_PUBLIC_HOST,
		Log:        log,
	}
}

func ConnectToMinio(cfg config.Config, log *slog.Logger) (*MinioStruct, error) {
	minioClient, err := minio.New(cfg.MINIO_HOST, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.MINIO_ROOT_USER, cfg.MINIO_ROOT_PASSWORD, ""),
		Secure: false,
	})
	if err != nil {
		log.Error(fmt.Sprintf("MinIO bilan ulanishda xatolik: %v", err))
		return nil, err
	}

	bucketName := "question-images"
	ctx := context.Background()
	exists, err := minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		log.Error(fmt.Sprintf("Bucket tekshirishda xatolik: %v", err))
	}
	if !exists {
		err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			log.Error(fmt.Sprintf("Bucket yaratilmadi: %v", err))
			return nil, err
		}
	}

	return &MinioStruct{
		Client:     minioClient,
		BucketName: bucketName,
		PublicHost: cfg.MINIO_PUBLIC_HOST,
		Log:        log,
	}, nil
}

func (m *MinioStruct) UploadFile(file multipart.File, fileName string) (string, error) {
	tempFile, err := os.CreateTemp("", fileName)
	if err != nil {
		m.Log.Error("Temporary fayl yaratishda xatolik", err)
		return "", err
	}
	defer os.Remove(tempFile.Name())

	_, err = tempFile.ReadFrom(file)
	if err != nil {
		m.Log.Error("Faylni o‘qishda xatolik", err)
		return "", err
	}

	objectName := uuid.NewString() + filepath.Ext(fileName)
	_, err = m.Client.FPutObject(context.Background(), m.BucketName, objectName, tempFile.Name(), minio.PutObjectOptions{
		ContentType: "image/jpeg",
	})
	if err != nil {
		m.Log.Error("Faylni MinIO ga yuklashda xatolik", err)
		return "", err
	}

	imageURL := fmt.Sprintf("%s/%s/%s", m.PublicHost, m.BucketName, objectName)
	return imageURL, nil
}
