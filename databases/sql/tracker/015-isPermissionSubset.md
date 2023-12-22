isPermissionSubset()
=============================

1. Given a [`permissions`](./010-permissions.md) object (`lhs`) representing an actor's permission claim and a second [`permissions`](./010-permissions.md) (`rhs`) object representing the set of expected [`permissions`](./010-permissions.md), return `true` if `lhs` is a subset of `rhs` or false otherwise.
2. An exception is thrown if the function has not been updated by any new [`permissions`](./010-permissions.md) values.
