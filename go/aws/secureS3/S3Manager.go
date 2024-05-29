package secureS3

import (
	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/sam-caldwell/monorepo/go/crypto/crypto"
)

// Manager is a struct to manage S3 operations.
//
//	(c) 2023 Sam Caldwell.  MIT License
type Manager struct {
	Client     *awss3.Client
	Bucket     string
	encryption crypto.Rsa
}
