/*
 * 0150-0010-func-createIntegerProperty.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function createIntegerProperty(propertyName varchar(64), propertyValue integer) returns uuid as
$$
declare
    propertyId uuid;
begin
    propertyId := (select createPropertyKey(propertyName) as id);

    insert into numericProperties(id, value) values (propertyId, propertyValue);
    return propertyId;
end;
$$ language plpgsql;
