package secureS3

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// FetchObject downloads an object from S3.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (s3 *Manager) FetchObject(key string) (string, error) {
	resp, err := s3.Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(s3.Bucket),
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
