/*
 * 0100-func-deleteAvatar.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function deleteAvatar(avatarId uuid) returns integer as
$$
declare
    count integer;
begin
    delete from avatars where id = avatarId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
