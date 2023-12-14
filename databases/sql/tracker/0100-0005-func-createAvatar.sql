/*
 * 0100-0005-func-createAvatar.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * When we create an Avatar, we must first register an entity uuid
 * and store any properties
 */

create or replace function createAvatar(avatarType varchar(1024), avatarHash text) returns uuid as
$$
declare
    avatarId uuid;
begin
    -- sanitize avatarType
    -- sanitize avatarHash
    avatarId := (select createEntity('avatar'));
    insert into avatars (id, mimeType, hash) values (avatarId,avatarType,avatarHash);
    return avatarId;
end;
$$ language plpgsql;
