/*
 * 062-updateTeamIcon.sql
 * (c) 2023 Sam Caldwell.  See License.txt
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
