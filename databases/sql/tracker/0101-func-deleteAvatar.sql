/*
 * 0101-func-deleteAvatar.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function deleteAvatar(avatarId uuid) returns integer as
$$
declare
    count integer;
begin
    --An avatar cannot be deleted if it is in use (user)
    delete from avatars where id = avatarId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
