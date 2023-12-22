/*
 * 062-deleteTeam.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function deleteTeam(teamId uuid) returns integer as
$$
declare
    count integer;
begin
    if deleteIconPreCheck(teamId) then
        delete from teams where id = teamId;
        get diagnostics count = ROW_COUNT;
        return count;
    else
        return 0; --error state
    end if;

end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * deleteTeamPreCheck() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function deleteTeamPreCheck(entityId uuid) returns boolean as
$$
begin
    /*
     * This is currently a placeholder.  This function should be overloaded later
     * as record consumers are defined so that we can prevent any record here from being
     * deleted until all related records are removed.
     */
    return true;
end;
$$ language plpgsql;
