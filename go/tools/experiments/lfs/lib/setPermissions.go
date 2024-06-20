package lfs

import "os"

func SetPermissions(fileName *string) (err error) {
	//fix permissions
	var uid, gid int
	if uid, gid, err = GetUserGroupIds(); err != nil {
		return err
	}
	if err = os.Chown(*fileName, uid, gid); err != nil {
		return err
	}
	if err = os.Chmod(*fileName, 0640); err != nil {
		return err
	}
	return nil
}
