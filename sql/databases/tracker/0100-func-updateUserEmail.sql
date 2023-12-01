/*
 * 0100-func-updateUserEmail.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function updateUserEmail(userId uuid, email varchar(256)) returns integer as
$$
declare
    count integer;
begin
    update users set email=email where id=userId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
