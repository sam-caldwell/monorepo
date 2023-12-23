### Table: TeamMemberships
1. The `TeamMemberships` table associates a `user` with a `team` to represent a user membership in a given team.

### Functions:
#### [`addUserToTeam()`](072-addUserToTeam.sql)
* Given a `userId` (UUID) and `teamId` (UUID), create a record in `teamMemberships` identified by `teamMemberships`.`id` (UUID) reflecting a team membership association.  
* The association identifier (`teamMemberships`.`id`) is registered in the `entity` table as a `teamAssociation` entity.
#### [`removeUserFromTeam()`](072-removeUserFromTeam.sql)
* Given a `userId` (UUID) and `teamId` (UUID), remove the matching record in `teamMemberships`
#### [`getTeamsForUser()`](072-getTeamsForUser.sql)
* Given a `userId` (UUID), return a list of teams (represented by their `teamId` UUIDs) to which the user belongs.
#### [`getUsersInTeam()`](072-getUsersInTeam.sql)
* Given a `teamId` (UUID), return a list of users (represented by their `userId` UUIDs) to which a team belongs.
