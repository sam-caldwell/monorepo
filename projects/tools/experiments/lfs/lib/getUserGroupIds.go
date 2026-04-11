package lfs

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
)

// GetUserGroupIds - Return the sudo user uid and gid
func GetUserGroupIds() (uid, gid int, err error) {
	var usr *user.User
	if sudoUser := os.Getenv("SUDO_USER"); sudoUser == "" {
		return os.Geteuid(), os.Getegid(), nil
	} else {
		if usr, err = user.Lookup(sudoUser); err != nil {
			return 0, 0, fmt.Errorf("failed to lookup user: %w", err)
		}
	}
	if uid, err = strconv.Atoi(usr.Uid); err != nil {
		return 0, 0, fmt.Errorf("failed to convert UID to int: %w", err)
	}
	if gid, err = strconv.Atoi(usr.Gid); err != nil {
		return 0, 0, fmt.Errorf("failed to convert GID to int: %w", err)
	}
	return uid, gid, err
}
