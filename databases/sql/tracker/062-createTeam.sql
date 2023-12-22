/*
 * 062-createTeam.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function createTeam(tName varchar(64), tIcon uuid, tOwner uuid, pOwner permissions,
                                      pTeam permissions, pEveryone permissions, tDescription text) returns uuid as
$$
declare
    teamId uuid;
begin
    teamId := (select createEntity('team'::entityType));
    insert into teams (id, name, iconId, ownerId, owner, team, everyone, description)
    values (teamId, tName, tIcon, tOwner, pOwner, pTeam, pEveryone, tDescription);
    return teamId;
end;
$$ language plpgsql;
