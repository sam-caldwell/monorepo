/*
 * 0010-0060-table-teamMembership.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Map users to teams to create team memberships.
 */
create table if not exists teamMembership
(
    -- the userId and teamId used for the mapping --
    userId uuid not null,
    teamId uuid not null,
    created  timestamp        not null default now(),
    primary key (userId, teamId),
    foreign key (userId) references users (id),
    foreign key (teamId) references teams (id)
);
