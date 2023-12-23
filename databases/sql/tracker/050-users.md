### Table: Users
1. The [`users`](./050-users.md) table represents the users in the `Tracker` system.
2. A [`user`](./050-users.md) may be either a human or software service interacting with the system.
3. The [`user`](./050-users.md) records contain user-identifying information but *NEVER* authentication information.
### Functions:
#### `createUser()`
* Create a new user object, identifying its first and last name, phone number, email address and and optional description.
* Before creating the record, we validate the names, email addresses, and phone numbers.
#### `deleteUserById()`
* Delete a user record using a `user`.`id` (UUID).
#### `getUserById()`
* Given a user's `entityId` (UUID), retrieve the associated record and return it as a JSON object.
#### `getUsersByEmail()`
* Given an email address string, return a list of JSON objects containing any matching user objects.
#### `getUsersByPhone()`
Given an phone number string, return a list of JSON objects containing any matching user objects.
