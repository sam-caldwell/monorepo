package secureS3

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// SearchObjectsByTag searches for objects with specific tags.
func (s *S3Manager) SearchObjectsByTag(tagKey, tagValue string) ([]string, error) {
	var keys []string
	paginator := s3.NewListObjectsV2Paginator(s.Client, &s3.ListObjectsV2Input{
		Bucket: aws.String(s.Bucket),
	})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.TODO())
		if err != nil {
			return nil, err
		}
		for _, obj := range page.Contents {
			tags, err := s.Client.GetObjectTagging(context.TODO(), &s3.GetObjectTaggingInput{
				Bucket: aws.String(s.Bucket),
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
