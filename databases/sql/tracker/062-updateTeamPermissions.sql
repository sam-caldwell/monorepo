/*
 * 062-updateTeamPermissions.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function updateTeamPermissions(teamId uuid, o permissions, t permissions, e permissions) returns integer as
$$
declare
    count integer;
begin
    if not validatePermissionSet(o, t, e) then
        raise exception 'invalid permission set';
    end if;
    update teams
    set owner=o,
        team=t,
        everyone=e
    where id = teamId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
