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
