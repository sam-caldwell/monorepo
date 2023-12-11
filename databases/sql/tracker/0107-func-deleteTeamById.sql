/*
 * 0107-func-deleteTeamById.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function deleteTeamById(teamId uuid) returns integer as
$$
declare
    count integer;
begin
    delete from teams where id=teamId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
