/*
 * 0010-0210-table-icons.sql
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
    constraint prohibit_zero_uuid check (id <> '00000000-0000-0000-0000-000000000000'),
    foreign key (id) references entity (id) on delete restrict
);

create index if not exists ndx_icons_hash on icons (hash);
create index if not exists ndx_icons_mimeType on icons (mimeType);
create index if not exists ndx_entity_context on entity (created);
