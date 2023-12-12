/*
 * 0106-func-updateUserName.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function updateUserName(userId uuid, newFirstName varchar(64), newLastName varchar(64)) returns integer as
$$
declare
    count integer;
begin
    update users set firstName=newFirstName, lastName=newLastName where id=userId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
