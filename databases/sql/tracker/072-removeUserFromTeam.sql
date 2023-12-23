/*
 * 072-removeUserFromTeam.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function removeUserFromTeam(uid uuid, tid uuid) returns integer as
$$
declare
    count integer;
begin
    if deleteTeamMembershipPreCheck(uid, tid) then
        delete from teamMemberships where userId = uid and teamId = tid;
        get diagnostics count = ROW_COUNT;
        return count;
    else
        return 0;
    end if;
end;
$$ language plpgsql;
/*
 *
 */
create or replace function deleteTeamMembershipPreCheck(uid uuid, tid uuid) returns boolean as
$$
begin
    return true;
end;
$$ language plpgsql;
