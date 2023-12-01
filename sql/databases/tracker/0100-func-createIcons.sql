/*
 * 0100-func-createIcons.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function createIcons(iconUrl text) returns uuid as $$
declare
    iconId uuid;
begin
    iconId := gen_random_uuid();
    insert into avatars (id,url) values(iconId,iconUrl);
    return iconId;
end;
$$ language plpgsql;
