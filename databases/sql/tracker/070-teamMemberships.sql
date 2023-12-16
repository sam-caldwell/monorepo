/*
 * 070-teamMemberships.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Map users to teams to create team memberships.
 */
create table if not exists teamMemberships
(
    id uuid      primary key not null,
    -- the userId and teamId used for the mapping --
    userId        uuid      not null,
    teamId        uuid      not null,
    created       timestamp not null default now(),
    foreign key (userId) references users (id),
    foreign key (teamId) references teams (id)
);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * entity indexes
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create unique index if not exists ndxTeamMembershipsUserIdTeamId on teamMemberships (userId, teamId);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * addUserToTeam()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function addUserToTeam(userId uuid, teamId uuid) returns integer as
$$
declare
    count integer;
    associationId uuid;
begin
    associationId := (select createEntity('user'::entityType));
    insert into teamMemberships (id, userId, teamId) values (associationId, userId, teamId);
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * removeUserFromTeam()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function removeUserFromTeam(uid uuid, tid uuid) returns integer as
$$
declare
    count integer;
begin
    delete from teamMemberships where userId = uid and teamId = tid;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getTeamsForUser()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getTeamsForUser(thisUserId uuid) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(jsonb_build_object(
            'teamId', teamId
        )) as team
    into result
    from teamMemberships
    where userId = thisUserId;
    return result;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getUsersInTeam()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getUsersInTeam(thisTeamId uuid) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(jsonb_build_object(
            'userId', userId
        )) as workflow
    into result
    from teamMemberships
    where teamId = thisTeamId;
    return result;
end;
$$ language plpgsql;
