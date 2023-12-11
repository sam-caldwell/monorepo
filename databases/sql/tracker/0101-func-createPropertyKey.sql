/*
 * 0102-func-createPropertyKey.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function createPropertyKey(pid uuid, propertyName varchar(64)) returns uuid as
$$
begin
    insert into propertyKeys (id, name) values (pid, propertyName);
    return pid;
end
$$ language plpgsql;
