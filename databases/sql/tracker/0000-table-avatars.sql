/*
 * 0000-table-avatars.sql
 * (c) 2023 Sam Caldwell. See License.txt
 *
 * A table that stores avatar references.
 */
create table if not exists avatars
(
    id  uuid primary key default gen_random_uuid(),
    url text
);
