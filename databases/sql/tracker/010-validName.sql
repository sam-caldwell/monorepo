/*
 * 010-validName.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * validate a given name string and return the boolean result.
 */
create or replace function validName(proposedName varchar(64)) returns boolean as
$$
declare
    result bool;
begin
    result := true; --ToDo: add regex validation to ensure all names meet naming conventions.
    return result;
end ;
$$ language plpgsql;
