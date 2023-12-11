/*
 * 0148-func-createIntegerProperty.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function createIntegerProperty(propertyName varchar(64), propertyValue integer) returns uuid as
$$
declare
    propertyId uuid;
begin
    propertyId := gen_random_uuid();
    insert into propertyKeys (pid, name) values (propertyId, propertyName);
    insert into numericProperties(id, value) values (propertyId, propertyValue);
    return propertyId;
end;
$$ language plpgsql;
