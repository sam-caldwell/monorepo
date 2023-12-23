/*
 * 062-updateTeamName.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function updateTeamName(teamId uuid, teamName varchar(64)) returns integer as
$$
declare
    count integer;
begin
    if not validname(teamName) then
        raise exception 'team name is invalid';
    end if;
    update teams
    set name=teamName
    where id = teamId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
