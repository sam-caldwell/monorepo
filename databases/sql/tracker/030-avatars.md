Avatars Table
=============

## Objectives
1. An `Avatar` is an image used by a `user` object to depict the user's account visually.
2. The `Avatar` database record identifies a `uuid`used to store the avatar file (rather than any source file name), along with it's file hash and mime type.
3. The `Avatar`should be unique to the `user` which makes it separate from the `icons` table records, which may be shared across different objects as a standard shared image library.
4. It is technically possible that a shared avatar image could exist if the hash were the same and that this would allow storage deduplication.  But there is no firm commitment on this matter as of now.