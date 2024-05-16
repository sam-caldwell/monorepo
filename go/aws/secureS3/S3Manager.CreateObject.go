package secureS3

import (
	"compress/gzip"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/sam-caldwell/monorepo/go/aws/tags"
	"github.com/sam-caldwell/monorepo/go/crypto/asymmetricCrypto"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"io"
)

// CreateObject uploads an object to S3.
// ToDo: Add ACL when creating object
func (mgr *Manager) CreateObject(key string, content *io.Reader, tags *tags.Tags, acl *Acl) (err error) {
	var gzipReader *io.Reader

	if encryptIo, err = asymmetricCrypto.NewReader(*content); err != nil {
		return err
	}
	if gzipReader, err = gzip.NewReader(*mgr.encryption.EncryptionReader()); err != nil {
		return err
	}

	_, err = mgr.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		ACL:     acl.ObjectCannedACL(),
		Bucket:  aws.String(mgr.Bucket),
		Key:     aws.String(key),
		Body:    *gzipReader,
		Tagging: aws.String(tags.String(words.Colon, words.Ampersand)),
	})
	return err
}
