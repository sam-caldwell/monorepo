package secureS3

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
)

// NewS3Manager creates a new instance of S3Manager.
//
//	(c) 2023 Sam Caldwell.  MIT License
func NewS3Manager(bucket string, region string) (*Manager, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return nil, err
	}
	client := awss3.NewFromConfig(cfg)

	return &Manager{
		Client: client,
		Bucket: bucket,
	}, nil
}
