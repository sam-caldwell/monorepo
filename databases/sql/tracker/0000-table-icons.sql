/*
 * 0000-table-icons.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * A table for identifying the URLs of icon files.
 */
create table if not exists icons
(
    id  uuid primary key default gen_random_uuid(),
    url text
);
