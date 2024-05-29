package secureS3

import "github.com/aws/aws-sdk-go-v2/service/s3/types"

// Acl - S3 Access control list
//
//	(c) 2023 Sam Caldwell.  MIT License
type Acl string

// Private - set private access
//
//	(c) 2023 Sam Caldwell.  MIT License
func (o *Acl) Private() *Acl {
	*o = "private"
	return o
}

// PublicRead - set public read
//
//	(c) 2023 Sam Caldwell.  MIT License
func (o *Acl) PublicRead() *Acl {
	*o = "public-read"
	return o
}

// ObjectCannedACL - set an object canned access control list
//
//	(c) 2023 Sam Caldwell.  MIT License
func (o *Acl) ObjectCannedACL() types.ObjectCannedACL {
	if *o == "" {
		return types.ObjectCannedACL("private")
	}
	return types.ObjectCannedACL(string(*o))
}
