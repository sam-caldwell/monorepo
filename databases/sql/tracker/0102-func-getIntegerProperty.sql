/*
 * 0148-func-getIntegerProperty.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function getIntegerProperty(propertyName varchar(64)) returns integer as
$$
declare
    result integer;
begin
    select value as v into result
    from numericProperties
    where id in (select id
                 from propertyKeys
                 where name = propertyName);
    return result;
end;
$$ language plpgsql;
