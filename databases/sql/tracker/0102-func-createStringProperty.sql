/*
 * 0148-func-createStringProperty.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function createStringProperty(propertyName varchar(64), propertyValue text) returns uuid as
$$
declare
    propertyId uuid;
begin
    propertyId := gen_random_uuid();
    insert into propertyKeys (pid, name) values (propertyId, propertyName);
    insert into stringProperties(id, value) values (propertyId, propertyValue);
    return propertyId;
end;
$$ language plpgsql;
