/*
 * 0120-func-updateUserPhone.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function updateUserPhone(userId uuid, phoneNumber varchar(20)) returns integer as
$$
declare
    count integer;
begin
    update users set phoneNumber=phoneNumber where id=userId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
