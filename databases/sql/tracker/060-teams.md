### Table: Teams
1. A [`team`](./060-teams.md) is an object representing a group of [`users`](./050-users.md) objects which represents the collective permissions of the group and communications with the group.
2. The [`team`](./060-teams.md) table only identifies the `team`.
3. A [`team`](./060-teams.md) object has the following characteristics:
	1. A `team`.`name` (validated by the `validName()` function).
	2. An [`icon`](./040-icons.md) object.
	3. An `owner` ([`users`](./050-users.md).`id`)
	4. [`Permissions`](./010-permissions.md) fields:
		1. `owner`
		2. `team`
		3. `everyone`

### Functions
#### [`createTeam()`](./062-createTeams.sql)
	* creating a team must deconflict the given permissions and validate the team properties.
#### [`deleteTeam()`](./062-deleteTeam.sql)
	*  given a `teamId` (UUID) delete the matching record.
#### [`getTeamById()`](./060-deleteTeam.sql)
* Given a `teamId` (UUID) return a JSON object containing the matching record.
#### [`getTeamsByOwnerId()`](./062-getTeamsByOwnerId.sql)
* Given an `ownerId` (user UUID) return a list of matching JSON team objects.
#### [`updateTeamDescription()`](./062-updateTeamDescription.sql)
* Given a `teamId` (UUID) and description text string, update the matching record description.
