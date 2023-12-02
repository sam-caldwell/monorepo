/*
 * 0148-func-setIntegerProperty.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function setIntegerProperty(propertyName varchar(64), propertyValue integer) returns uuid as
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
