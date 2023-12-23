/*
 * 062-updateTeamDescription.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function updateTeamDescription(teamId uuid, d text) returns integer as
$$
declare
    count integer;
begin
    update teams
    set description=d
    where id = teamId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
