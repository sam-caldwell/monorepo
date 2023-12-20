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
create index if not exists ndxAvatarsHash on avatars (hash);
create index if not exists ndxAvatarsMimeType on avatars (mimeType);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * createAvatar() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function createAvatar(t mimeType, h varchar(1024)) returns uuid as
$$
declare
    Id uuid;
begin
    -- sanitize avatarHash
    Id := (select createEntity('avatar'::entityType));
    insert into avatars (id, mimeType, hash) values (Id, t, h);
    return Id;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * deleteAvatar() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function deleteAvatar(entityId uuid) returns integer as
$$
declare
    count integer;
begin
    if deleteAvatarPreCheck(entityId) then
        delete from avatars where id = entityId;
        get diagnostics count = ROW_COUNT;
        return count;
    else
        return 0; --error state
    end if;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * deleteAvatarPreCheck() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function deleteAvatarPreCheck(entityId uuid) returns boolean as
$$
begin
    /*
     * This is currently a placeholder.  This function should be overloaded later
     * as record consumers are defined so that we can prevent any record here from being
     * deleted until all related records are removed.
     */
    return true;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getAvatarById() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getAvatarById(entityId uuid) returns jsonb as
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
    where id = entityId
    limit 1;
    return result;
end;
$$ language plpgsql;
