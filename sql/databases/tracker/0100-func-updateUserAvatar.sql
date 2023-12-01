/*
 * 0100-func-updateUserAvatar.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function updateUserAvatar(userId uuid, avatarId varchar(256)) returns integer as
$$
declare
    count integer;
begin
    update users set avatarId=avatarId where id=userId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
