/*
 * 0003-0020-func-deletePropertyByName.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function deletePropertyByName(propertyName varchar(64)) returns integer as
$$
declare
    count integer;
begin
    delete from propertyKeys where id=(select id from propertyKeys where name=propertyName);
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
