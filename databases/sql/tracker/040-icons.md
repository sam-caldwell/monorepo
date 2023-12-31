Icons
=====

## Objectives
1. The `icons` table stores the identifying `entityId` (UUID) used to form the URL to an icon image file, along with its hash and mimeType.
2. Unlike the context-specific `Avatar` which should depict a single user, the `icons` table identifies images which are shared and used to represent system objects (e.g. projects, ticket types, etc.).
6. It is technically possible that a shared avatar image could exist if the hash were the same and that this would allow storage deduplication.  But there is no firm commitment on this matter as of now.
### Functions:
#### `createIcon()`
* Create an `iconId` (UUID) from a given `type` ([`mimeType`](./010-mimeType.md)) and `hash` (`varchar`).
* This function will generate the `iconId` as an [`entity`.`id`](./020-entity.md) using `entityType` `iconId`.
* Return the `iconId` to the caller.
#### `deleteIcon()`
* Given an `iconId` (UUID), perform a pre-check to ensure all dependencies have been cleaned up before deleting the iconId.
#### `getIconById()`
* Given an `iconId` (UUID), return a JSON object containing the record for the matching icon.