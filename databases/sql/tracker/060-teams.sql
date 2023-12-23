/*
 * 060-teams.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Create teams of users.
 */

create table if not exists teams
(
    -- the teamId used to uniquely identify the team
    id          uuid primary key not null,
    -- the unique team name --
    name        varchar(64)      not null,
    -- a uuid representing the icon for the workflow.
    iconId      uuid             not null,
    -- the owner is the user who created the team or who has team ownership --
    ownerId     uuid             not null,
    -- permissions are granted to the owner, team and everyone --
    owner       permissions      not null default 'delete'::permissions,
    team        permissions      not null default 'read'::permissions,
    everyone    permissions      not null default 'read'::permissions,
    -- --
    created     timestamp        not null default now(),
    -- descriptive text --
    description text,
    -- --
    constraint validateTeamName check (validName(name)),
    -- --
    foreign key (ownerId) references users (id),
    foreign key (iconId) references icons (id),
    foreign key (id) references entity (id) on delete restrict
);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * entity indexes
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create unique index if not exists ndxTeamsName on teams (name);
create index if not exists ndxTeamsCreated on teams (created);
create index if not exists ndxTeamsOwnerId on teams (ownerId);
create index if not exists ndxTeamsIconId on teams (iconId);
create index if not exists ndxTeamsOwner on teams (owner);
create index if not exists ndxTeamsTeam on teams (team);
create index if not exists ndxTeamsEveryone on teams (everyone);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateTeamDescription() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */

/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updatePermissions() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updatePermissions(teamId uuid, o permissions,t permissions, e permissions) returns integer as
$$
declare
    count integer;
begin
    update teams
    set owner=o,
        team=t,
        everyone=e
    where id = teamId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;

/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateTeamIcon() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateTeamIcon(teamId uuid, newIconId uuid) returns integer as
$$
declare
    count integer;
begin
    update teams
    set iconId=newIconId
    where id = teamId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;

/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateTeamOwner() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateTeamOwner(teamId uuid, newOwnerId uuid) returns integer as
$$
declare
    count integer;
begin
    update teams
    set ownerId=newOwnerId
    where id = teamId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;

/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateTeamName() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateTeamName(teamId uuid, teamName varchar(64)) returns integer as
$$
declare
    count integer;
begin
    update teams
    set name=teamName
    where id = teamId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * deleteIconPreCheck() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function deleteIconPreCheck(entityId uuid) returns boolean as
$$
begin
    /*
     * Overloaded by Workflows...
     */
    return (select count(id) from teams where iconId = entityId) == 0;
end;
$$ language plpgsql;
