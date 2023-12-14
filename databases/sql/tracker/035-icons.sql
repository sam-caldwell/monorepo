/*
 * 035-icons.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * A table for identifying the URLs of icon files.
 */
create table if not exists icons
(
    id       uuid primary key not null,
    hash     varchar(1024)    not null unique, -- deduplicate our file stores
    mimeType varchar(1024)    not null,
    created  timestamp        not null default now(),
    foreign key (id) references entity (id) on delete restrict
);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * entity indexes
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create index if not exists ndxIconsCreated on icons (created);
create index if not exists ndxIconsHash on icons (hash);
create index if not exists ndxIconsMimeType on icons (mimeType);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * createIcon()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function createIcon(Type entityType, Hash varchar(1024)) returns uuid as
$$
declare
    Id uuid;
begin
    -- sanitize avatarHash
    Id := (select createEntity('icon'::entityType));
    insert into avatars (id, mimeType, hash) values (Id, Type, Hash);
    return Id;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * deleteIcons()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function deleteIcons(iconId uuid) returns integer as
$$
declare
    count integer;
begin
    -- an icon cannot be deleted if it is in use (project, workflow, ticketType, team, etc)
    delete from icons where id = iconId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
* getIconsById()
* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
*/
create or replace function getIconsById(iconId uuid) returns text as
$$
declare
    result text;
begin
    select jsonb_build_object(
                   'id', id,
                   'hash', hash,
                   'mimeType', mimeType
               ) as icons
    into result
    from icons
    where id = iconId
    limit 1;
    return result;
end
$$ language plpgsql;
