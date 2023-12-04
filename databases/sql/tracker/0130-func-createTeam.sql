/*
 * 0130-func-createTeam.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function createTeam(name varchar(64), iconId uuid, ownerId uuid, owner permissions,
                                      team permissions, everyone permissions, description text) returns uuid as
$$
declare
    teamId uuid;
begin
    teamId := gen_random_uuid();
    insert into users (id, name, iconId, ownerId, owner, team, everyone, description)
    values (teamId, name, iconId, ownerId, owner, team, description);
    return teamId;
end;
$$ language plpgsql;
