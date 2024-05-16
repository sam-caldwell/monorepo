package secureS3

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// FetchObject downloads an object from S3.
func (mgr *Manager) FetchObject(key string) (string, error) {
	resp, err := mgr.Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(mgr.Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return "", err
	}
	defer func() { _ = resp.Body.Close() }()

	buf := new(strings.Builder)

	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
