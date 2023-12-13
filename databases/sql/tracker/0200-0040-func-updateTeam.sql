/*
 * 0200-0040-func-updateTeam.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function updateTeam(teamId uuid, teamName varchar(64), iconId uuid, ownerId uuid, owner permissions,
                                      team permissions, everyone permissions, description text) returns integer as
$$
declare
    count integer;
begin
    update teams
    set name=teamName,
        iconId=iconId,
        ownerId=ownerId,
        owner=owner,
        team=team,
        everyone=everyone,
        description=description
    where id = teamId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
