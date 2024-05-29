package file

const (
	/*
	   Permission Scopes.

	       Permission scopes are values the permission values must be shifted
	          then combined with the OR operator to create a final POSIX permission.

	           For example:
	               Given an owner permission...

	                   r|w|x << AppliesToOwner (0x04)
	                   1 1 1 << 0x06
	                   1 1 1 0 0 0 0 0 0

	               ...a group permission...
	                   1 1 1 << AppliesToGroup (0x03)
	                   0 0 0 1 1 1 0 0 0

	               ...and permissions for everyone...

	                   1 1 1 << AppliesToEveryone (0x01)
	                   0 0 0 0 0 0 1 1 1

	               We can combine terms as...

	                    AppliesToOwner || AppliesToGroup  || AppliesToEveryone
	                   (1 1 1 << 0x06) || (1 1 1 << 0x03) || (1 1 1 << 0x03)
	                   0001 1100 0000 || 0000 0011 1000 | 0000 0000 0111

	                   As we can see, this combines to create a 9-bit permission.
	                       0001 1100 0000
	                       0000 0011 1000
	                       0000 0000 0111
	                       --------------
	                       0001 1111 1111
	*/

	//AppliesToOwner - permissions applied to owner
	AppliesToOwner = 0x06

	//AppliesToGroup - permissions applied to group
	AppliesToGroup = 0x03

	//AppliesToEveryone -  permissions applied to everyone
	AppliesToEveryone = 0x01

	/*
	   Permission flags.  Use bitwise OR to combine them.
	   This follows the POSIX rwx permission flags (see above).
	*/
)
