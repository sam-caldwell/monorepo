package secureS3

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
)

// DeleteObject - delete the identified object
//
//	(c) 2023 Sam Caldwell.  MIT License
func (s3 *Manager) DeleteObject(ctx context.Context, key string) error {
	_, err := s3.Client.DeleteObject(ctx, &awss3.DeleteObjectInput{
		Bucket: aws.String(s3.Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("unable to delete object %q from bucket %q, %v", key, s3.Bucket, err)
	}

	// Wait until the object is deleted
	err = s3.Client.WaitUntilObjectNotExists(ctx, &awss3.HeadObjectInput{
		Bucket: aws.String(s3.Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("error occurred while waiting for object %q to be deleted, %v", key, err)
	}

	return nil
}
