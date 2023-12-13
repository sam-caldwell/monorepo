/*
 * 0002-0010-func-getIconsById.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function getIconsById(iconId uuid) returns text as
$$
declare
    result text;
begin
    select url into result
    from icons
    where id = iconId;
    return result;
end
$$ language plpgsql;
