/*
 * 0003-0000-table-propertyKeys.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * A simple key-value property table
 */


create table if not exists propertyKeys
(
    id   uuid primary key default gen_random_uuid(),
    name varchar(64) not null unique
);
