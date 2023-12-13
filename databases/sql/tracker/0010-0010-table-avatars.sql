/*
 * 0010-0110-table-avatars.sql
 * (c) 2023 Sam Caldwell. See License.txt
 *
 * A table that will identify an avatar.  The uuid (id) will be used to identify the
 * actual file(name) on disk.
 */
create table if not exists avatars
(
    id       uuid primary key not null,
    hash     varchar(1024)    not null unique, -- deduplicate our file stores
    mimeType varchar(1024)    not null,
    created  timestamp        not null default now(),
    constraint prohibit_zero_uuid check (id <> '00000000-0000-0000-0000-000000000000'),
    foreign key (id) references entity (id) on delete restrict
);

