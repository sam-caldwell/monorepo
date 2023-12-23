/*
 * 032-createAvatar.sql
 * (c) 2023 Sam Caldwell. See License.txt
 */
create or replace function createAvatar(t mimeType, h varchar(1024)) returns uuid as
$$
declare
    Id uuid;
begin
    -- sanitize avatarHash to verify it is proper length and hexadecimal.
    Id := (select createEntity('avatar'::entityType));
    insert into avatars (id, mimeType, hash) values (Id, t, h);
    return Id;
end;
$$ language plpgsql;
