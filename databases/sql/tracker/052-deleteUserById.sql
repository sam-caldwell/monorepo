/*
 * 052-deleteUsersById.sql
 * (c) 2023 Sam Caldwell. See License.txt
 */
create or replace function deleteUsersById(userId uuid) returns integer as
$$
declare
    count integer;
begin
    if deleteUsersByIdPreCheck(userId) then
        delete from users where id = userId;
        get diagnostics count = ROW_COUNT;
        return count;
    else
        return 0;
    end if;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * deleteUsersByIdPreCheck() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function deleteUsersByIdPreCheck(entityId uuid) returns boolean as
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
