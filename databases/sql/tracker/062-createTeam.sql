/*
 * 062-createTeam.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function createTeam(newTeamName varchar(64), tIcon uuid, tOwner uuid, pOwner permissions,
                                      pTeam permissions, pEveryone permissions, tDescription text) returns uuid as
$$
declare
    teamId uuid;
begin
    if not validName(newTeamName) then
        raise exception 'invalid team name';
    end if;
    if validatePermissionSet(pOwner, pTeam, pEveryone) then
        raise exception 'invalid permission set';
    end if;
    teamId := (select createEntity('team'::entityType));
    insert into teams (id, name, iconId, ownerId, owner, team, everyone, description)
    values (teamId, newTeamName, tIcon, tOwner, pOwner, pTeam, pEveryone, tDescription);
    return teamId;
end;
$$ language plpgsql;
