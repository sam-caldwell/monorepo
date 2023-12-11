/*
 * 0102-func-createStringProperty.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function createStringProperty(propertyName varchar(64), propertyValue text) returns uuid as
$$
declare
    propertyId uuid;
begin
    propertyId := (select createPropertyKey(propertyName) as id);
    insert into stringProperties(id, value) values (propertyId, propertyValue);
    return propertyId;
end;
$$ language plpgsql;
