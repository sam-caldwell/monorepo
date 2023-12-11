/*
 * 0101-func-getAvatarById.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function getAvatarById(avatarId uuid) returns jsonb as
$$
declare
    result jsonb;
begin

    select jsonb_build_object(
            'id', id::text,
            'url', url
        ) as avatars into result
    from avatars
    where id = avatarId
    limit 1;

    return result;
end;
$$ language plpgsql;
