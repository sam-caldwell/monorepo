/*
 * 0102-func-createPropertyKey.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function createPropertyKey(propertyName varchar(64)) returns uuid as
$$
declare
    propertyId uuid;
begin
    propertyId:=gen_random_uuid();
    insert into propertyKeys (id, name) values (propertyId, propertyName);
    return propertyId;
end
$$ language plpgsql;
