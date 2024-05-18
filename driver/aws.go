package driver

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	privConf "github.com/backend-magang/halo-suster/config"
)

func InitS3Client(conf privConf.Config) *s3.Client {
	log.Println("[S3] initialized...")

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithSharedConfigProfile("default"),
		config.WithRegion(conf.S3Region),
	)

	if err != nil {
		log.Println("[S3] failed to Connect S3 Client, err: ", err)
		return nil
	}

	client := s3.NewFromConfig(cfg)

	log.Println("[S3] connected...")

	return client
}
