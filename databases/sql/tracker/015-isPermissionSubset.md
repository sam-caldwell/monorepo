comparePermissions()
====================

```postgresql
create or replace function isPermissionSubset(lhs permissions, rhs permissions) returns boolean as
```
1. Assume the caller has a permission (lhs) and a permission (rhs)
2. The user wants to know if lhs is a subset of the permissions rhs defines.
