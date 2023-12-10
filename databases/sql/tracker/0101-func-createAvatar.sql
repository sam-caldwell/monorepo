/*
 * 0101-func-createAvatar.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function createAvatar(avatarUrl text) returns uuid as $$
declare
    avatarId uuid;
begin
    avatarId := gen_random_uuid();
    insert into avatars (id,url) values(avatarId,avatarUrl);
    return avatarId;
end;
$$ language plpgsql;
