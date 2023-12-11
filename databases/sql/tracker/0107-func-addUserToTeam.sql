/*
 * 0107-func-addUserToTeam.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function addUserToTeam(userId uuid, teamId uuid) returns integer as
$$
declare
    count integer;
begin
    insert into teamMembership (userId, teamId) values (userId, teamId);
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
