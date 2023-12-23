/*
 * 052-updateUserDescription.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function updateUserDescription(userId uuid, newDescription text) returns integer as
$$
declare
    count integer;
begin
    update users set description=newDescription where id = userId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
