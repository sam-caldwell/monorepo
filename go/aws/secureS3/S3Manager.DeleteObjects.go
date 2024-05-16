package secureS3

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
)

func (mgr *Manager) DeleteObject(ctx context.Context, key string) error {
	_, err := mgr.Client.DeleteObject(ctx, &awss3.DeleteObjectInput{
		Bucket: aws.String(mgr.Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("unable to delete object %q from bucket %q, %v", key, mgr.Bucket, err)
	}

	// Wait until the object is deleted
	err = mgr.Client.WaitUntilObjectNotExists(ctx, &awss3.HeadObjectInput{
		Bucket: aws.String(mgr.Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("error occurred while waiting for object %q to be deleted, %v", key, err)
	}

	return nil
}
