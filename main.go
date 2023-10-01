package main

import (
	"database/sql"
	"ecom/api"
	"ecom/db/sqlc"
	"ecom/db/util"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Can't load config", err)
	}

	db, err := sql.Open(config.DBDriver, config.DBUrl)
	if err != nil {
		log.Fatal("error connection to DB", err)
	}
	store := sqlc.NewStore(db)
	s3uploader := loadS3(&config)
	server, err := api.NewServer(config, store, s3uploader)
	if err != nil {
		log.Fatal("error while creating new server", err)
	}

	err = server.Start(config.HttpServerAddr)
	if err != nil {
		log.Fatal("error while running the server", err)
	}
}

func loadS3(config *util.Config) *manager.Uploader {
	s3Config := aws.Config{
		Region:      *aws.String(config.S3REGION),
		Credentials: credentials.NewStaticCredentialsProvider(config.AWSACCESSKEY, config.AWSSECRETKEY, ""),
	}
	s3 := s3.NewFromConfig(s3Config)
	uploader := manager.NewUploader(s3)
	return uploader
}
