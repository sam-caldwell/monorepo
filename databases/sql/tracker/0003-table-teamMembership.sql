/*
 * 0003-table-teamMembership.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Map users to teams to create team memberships.
 */
create table if not exists teamMembership
(
    -- the userId and teamId used for the mapping --
    userId uuid not null,
    teamId uuid not null,
    PRIMARY KEY (userId, teamId),
    FOREIGN KEY (userId) REFERENCES users (id),
    FOREIGN KEY (teamId) REFERENCES teams (id)
);
