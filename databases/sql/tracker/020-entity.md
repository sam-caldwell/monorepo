entity table
============

## Objectives
1. The `entity` table identifies all objects in the database with a required `entityType` and optional `context`.
2. The goal of the `entity` table is to guarantee that identified entities are legitimate (not easily spoofed by exploitation of application vulnerabilities), and that the `entity` persists past the point where the identified object is deleted as a means of preserving its audit trail.  For example:
		1. When a `user` object is created, the object is identified by `entityId`. Any reference to a `user` can be validated in that the `uuid` describing the user is an `entityId` of type `user`.
		2. When activity occurs, the `user` is identified in logs by `entityId` (though logs may provide additional context over the lifetime of the `user`).
		3. When the `user` is deleted from the system, the `entityId` identifying that `user` remains in the database, and thus the logs continue with integrity to identify the `user` and to provide context of the `user` object's lifetime.
