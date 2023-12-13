/*
 * 0002-0125-0000-table-icons.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * A table for identifying the URLs of icon files.
 */
create table if not exists icons
(
    id  uuid primary key default gen_random_uuid(),
    url text
    constraint prohibit_zero_uuid check (id <> '00000000-0000-0000-0000-000000000000')
);
