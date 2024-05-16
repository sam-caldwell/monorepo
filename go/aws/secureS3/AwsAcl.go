package secureS3

import "github.com/aws/aws-sdk-go-v2/service/s3/types"

type Acl string

func (o *Acl) Private() *Acl {
	*o = "private"
	return o
}

func (o *Acl) PublicRead() *Acl {
	*o = "public-read"
	return o
}

func (o *Acl) ObjectCannedACL() types.ObjectCannedACL {
	if *o == "" {
		return types.ObjectCannedACL("private")
	}
	return types.ObjectCannedACL(string(*o))
}
