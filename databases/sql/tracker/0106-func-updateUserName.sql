/*
 * 0106-func-updateUserName.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function updateUserName(userId uuid, firstName varchar(64), lastName varchar(64)) returns integer as
$$
declare
    count integer;
begin
    update users set firstName=firstName, lastName=lastName where id=userId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
