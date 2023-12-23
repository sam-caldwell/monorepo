/*
 * 070-teamMemberships.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Map users to teams to create team memberships.
 */
create table if not exists teamMemberships
(
    id      uuid primary key not null,
    -- the userId and teamId used for the mapping --
    userId  uuid             not null,
    teamId  uuid             not null,
    created timestamp        not null default now(),
    foreign key (userId) references users (id),
    foreign key (teamId) references teams (id)
);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * entity indexes
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create unique index if not exists ndxTeamMembershipsUserIdTeamId on teamMemberships (userId, teamId);
