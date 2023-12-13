/*
 * 0250-0300-func-removeUserFromTeam.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function removeUserFromTeam(uid uuid, tid uuid) returns integer as
$$
declare
    count integer;
begin
    delete from teamMembership where userId=uid and teamId=tid;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
