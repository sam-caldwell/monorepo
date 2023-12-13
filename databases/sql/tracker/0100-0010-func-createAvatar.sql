/*
 * 0100-0010-func-createAvatar.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function createAvatar(avatarUrl text) returns uuid as
$$
declare
    avatarId uuid;
begin
    if (select isvalidurl(avatarUrl) as result) then
        avatarId := gen_random_uuid();
        insert into avatars (id, url) values (avatarId, avatarUrl);
        return avatarId;
    else
        return '00000000-0000-0000-0000-000000000000';
    end if;
end;
$$ language plpgsql;
