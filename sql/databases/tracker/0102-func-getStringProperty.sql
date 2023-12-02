/*
 * 0148-func-getStringProperty.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function getStringProperty(propertyName varchar(64)) returns text as
$$
declare
    result text;
begin
    select value as v into result
    from stringProperties
    where id in (select id
                 from propertyKeys
                 where name = propertyName);
    return result;
end;
$$ language plpgsql;
