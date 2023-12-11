/*
 * 0106-func-deleteUsersById.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function deleteUsersById(userId uuid) returns integer as
$$
declare
    count integer;
begin
    delete from users where id = userId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
