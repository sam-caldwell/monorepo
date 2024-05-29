package secureS3

import (
	"compress/gzip"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/sam-caldwell/monorepo/go/aws/tags"
	"github.com/sam-caldwell/monorepo/go/crypto/crypto"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"io"
)

// CreateObject uploads an object to S3.
//
//	(c) 2023 Sam Caldwell.  MIT License
//
// ToDo: Add ACL when creating object
func (s3 *Manager) CreateObject(key string, content *io.Reader, tags *tags.Tags, acl *Acl) (err error) {
	var gzipReader *io.Reader

	if encryptIo, err = crypto.NewReader(*content); err != nil {
		return err
	}
	if gzipReader, err = gzip.NewReader(*s3.encryption.EncryptionReader()); err != nil {
		return err
	}

	_, err = s3.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		ACL:     acl.ObjectCannedACL(),
		Bucket:  aws.String(s3.Bucket),
		Key:     aws.String(key),
		Body:    *gzipReader,
		Tagging: aws.String(tags.String(words.Colon, words.Ampersand)),
	})
	return err
}
