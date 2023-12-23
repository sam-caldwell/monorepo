###  Table: `entity`

1. The [`entity`](./020-entity.sql) table identifies all objects in the database with a required [`entityType`](./010-entityType.md) and optional `context`.
2. There are two goals for the [`entity`](./020-entity.sql) table:
	1. Guarantee the UUID identifiers used are legitimate and unable to be easily spoofed by the application tier.
	2. Create a simple, write-only store of these [`entity`](./020-entity.sql) UUID values such that if the underlying entity (e.g. [`user`](./050-users.md),[`project`](./110-projects.md)) is deleted, the [`entity`](./020-entity.sql) is not deleted and will continue to provide a minimal audit trail.  For example:
		1. When a [`user`](./050-users.md) object is created, the object is identified by `entityId`. Any reference to a `user` can be validated in that the UUID describing the user is an `entityId` of type [`user`](./050-users.md).
		2. When activity occurs, the [`user`](./050-users.md) is identified in logs by `entityId` (though logs may provide additional context over the lifetime of the [`user`](./050-users.md)).
		3. When the [`user`](./050-users.md) is deleted from the system, the `entityId` identifying that [`user`](./050-users/md) remains in the database, and thus the logs continue with integrity to identify the [`user`](./050-users.md) and to provide context of the [`user`](./050-users.md) object's lifetime.
3. Once an `entity` is created, it cannot be deleted or updated, as protected by two trigger objects:
	1. [`preventEntityDelete`](./010-preventDelete.sql) and 
	2. [`preventEntityUpdate`](./010-preventUpdate.sql).

### Functions:
#### `createEntity()`
* Generate an `entityId` (UUID)
* Create `entity` record
* Return `entityId` to caller.
#### `getEntity()`
* lookup a single `entity` record and return as a JSON object.


#### <span style="color:yellow">Todo:</span>
1. We need an ability to get a list of entities based on type and/or context.
