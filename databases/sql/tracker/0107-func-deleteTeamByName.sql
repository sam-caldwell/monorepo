/*
 * 0107-func-deleteTeamByName.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function deleteTeamByName(teamName varchar(64)) returns uuid as
$$
declare
    count integer;
begin
    delete from teams where name=teamName;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
