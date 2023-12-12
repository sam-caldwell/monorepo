/*
 * 0108-func-removeUserFromTeam.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function removeUserFromTeam(userId uuid, teamId uuid) returns integer as
$$
declare
    count integer;
begin
    delete from teamMembership where userId==userId and teamId==teamId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
