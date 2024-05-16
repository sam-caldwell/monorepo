package secureS3

import (
	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/sam-caldwell/monorepo/go/crypto/asymmetricCrypto"
)

// Manager is a struct to manage S3 operations.
type Manager struct {
	Client     *awss3.Client
	Bucket     string
	encryption asymmetricCrypto.Manager
}
