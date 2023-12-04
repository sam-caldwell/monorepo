/*
 * 0120-func-updateUserDescription.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function updateUserDescription(userId uuid, description text) returns integer as
$$
declare
    count integer;
begin
    update users set description=description where id = userId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
