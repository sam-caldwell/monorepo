permissions
=================

An enumerated type which defines permissions which can be assigned to a user, team or everyone:

| permission | description | 
| ----------- | ----------- |
| `none` | Grant no permission (implicit deny) |
| `read` | Allow an object to be `read` |
| `create` | Allow an object to be `created` or `read` |
| `update` | Allow an object to be `updated`, `created` or `read` |
| `delete` | Allow an object to be `deleted`, `updated`, `created` or `read` |

## Related Functions
* [`isPermissionSubset()`](./015-isPermissionSubset.md)
