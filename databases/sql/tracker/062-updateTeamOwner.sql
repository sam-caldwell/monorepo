/*
 * 062-updateTeamOwner.sql
 * (c) 2023 Sam Caldwell.  See License.txt
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
