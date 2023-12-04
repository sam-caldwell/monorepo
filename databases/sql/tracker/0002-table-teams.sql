/*
 * 0002-table-teams.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Create teams of users.
 */

create table if not exists teams
(
    -- the teamId used to uniquely identify the team
    id          uuid primary key     default gen_random_uuid(),
    -- the unique team name --
    name        varchar(64) not null unique,
    -- a uuid representing the icon for the workflow.
    iconId      uuid        not null,
    -- the owner is the user who created the team or who has team ownership --
    ownerId     uuid        not null,
    -- permissions are granted to the owner, team and everyone --
    owner       permissions not null default 'delete',
    team        permissions not null default 'read',
    everyone    permissions not null default 'read',
    -- descriptive text --
    description text,
    foreign key (ownerId) references users (id),
    foreign key (iconId) references icons(id)
);

