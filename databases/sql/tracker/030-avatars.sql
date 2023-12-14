/*
 * 030-avatars.sql
 * (c) 2023 Sam Caldwell. See License.txt
 *
 * A table that will identify an avatar.  The uuid (id) will be used to identify the
 * actual file(name) on disk.
 */
create table if not exists avatars
(
    id       uuid primary key not null,
    hash     varchar(1024)    not null, -- deduplicate our file stores
    mimeType varchar(1024)    not null,
    created  timestamp        not null default now(),
    foreign key (id) references entity (id) on delete restrict
);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * entity indexes
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create index if not exists ndxAvatarsCreated on avatars (created);
create unique index if not exists ndxAvatarsHash on avatars (hash);
create index if not exists ndxAvatarsMimeType on avatars (mimeType);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * createAvatar() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function createAvatar(Type entityType, Hash varchar(1024)) returns uuid as
$$
declare
    Id uuid;
begin
    -- sanitize avatarHash
    Id := (select createEntity('avatar'::entityType));
    insert into avatars (id, mimeType, hash) values (Id, Type, Hash);
    return Id;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * deleteAvatar() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function deleteAvatar(targetId uuid) returns integer as
$$
declare
    count integer;
begin
    count := (select count(id) from users where avatarId = targetId);
    if count > 0 then
        count := 0;
        raise exception 'cannot delete avatar if it is in use';
    end if;
    delete from avatars where id = targetId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getAvatarById() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getAvatarById(avatarId uuid) returns jsonb as
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
    where id = avatarId
    limit 1;
    return result;
end;
$$ language plpgsql;
