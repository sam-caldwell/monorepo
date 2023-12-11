/*
 * 0103-func-getIconsById.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function getIconsById(id uuid) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(jsonb_build_object(
            'id', id,
            'url', url
        )) as icons into result
    from icons
    where id == id;
    return result;
end;
$$ language plpgsql;
