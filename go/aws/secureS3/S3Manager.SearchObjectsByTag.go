package secureS3

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsS3 "github.com/aws/aws-sdk-go-v2/service/s3"
)

// SearchObjectsByTag searches for objects with specific tags.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (s3 *Manager) SearchObjectsByTag(tagKey, tagValue string) ([]string, error) {
	var keys []string
	paginator := awsS3.NewListObjectsV2Paginator(s3.Client, &awsS3.ListObjectsV2Input{
		Bucket: aws.String(s3.Bucket),
	})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.TODO())
		if err != nil {
			return nil, err
		}
		for _, obj := range page.Contents {
			tags, err := s3.Client.GetObjectTagging(context.TODO(), &awsS3.GetObjectTaggingInput{
				Bucket: aws.String(s3.Bucket),
				Key:    obj.Key,
			})
			if err != nil {
				return nil, err
			}
			for _, tag := range tags.TagSet {
				if *tag.Key == tagKey && *tag.Value == tagValue {
					keys = append(keys, *obj.Key)
					break
				}
			}
		}
	}
	return keys, nil
}
