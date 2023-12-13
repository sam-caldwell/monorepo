/*
 * 0100-0000-table-avatars.sql
 * (c) 2023 Sam Caldwell. See License.txt
 *
 * A table that stores avatar references.
 */
create table if not exists avatars
(
    id  uuid primary key default gen_random_uuid(),
    url text,
    constraint prohibit_zero_uuid check (id <> '00000000-0000-0000-0000-000000000000')
);
