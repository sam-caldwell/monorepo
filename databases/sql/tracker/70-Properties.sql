/*
 * 70-Properties.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * A simple key-value property table
 */

create table if not exists propertyKeys
(
    id      uuid primary key not null,
    name    varchar(64)      not null,
    -- --
    created timestamp        not null default now(),
        foreign key (id) references entity (id) on delete restrict
);

create table if not exists stringProperties
(
    id    uuid primary key not null,
    value text             not null,
    foreign key (id) references propertyKeys (id)
        on delete cascade
);

create table if not exists numericProperties
(
    id    uuid primary key not null,
    value integer          not null,
    foreign key (id) references propertyKeys (id)
        on delete cascade
);

create unique index ndxPropertyKeysName on propertykeys (name);
create unique index ndxPropertyKeysCreated on propertykeys (created);
