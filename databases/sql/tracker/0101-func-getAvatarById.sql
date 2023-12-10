/*
 * 0101-func-getAvatarById.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function getAvatarById(id uuid) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(jsonb_build_object(
            'id', id,
            'url', url
        )) as avatars into result
    from avatars
    where id == id;
    return result;
end;
$$ language plpgsql;
