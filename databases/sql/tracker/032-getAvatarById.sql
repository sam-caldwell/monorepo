/*
 * 032-getAvatarById.sql
 * (c) 2023 Sam Caldwell. See License.txt
 */
create or replace function getAvatarById(entityId uuid) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_build_object(
                   'id', id,
                   'hash', hash,
                   'mimeType', mimeType
               ) as avatar
    into result
    from avatars
    where id = entityId
    limit 1;
    return result;
end;
$$ language plpgsql;

