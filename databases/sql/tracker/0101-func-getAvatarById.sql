/*
 * 0101-func-getAvatarById.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function getAvatarById(avatarId uuid) returns text as
$$
declare
    result text;
begin

    select url into result from avatars where id = avatarId limit 1;

    return result;
end;
$$ language plpgsql;
