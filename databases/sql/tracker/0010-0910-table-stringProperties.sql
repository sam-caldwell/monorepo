/*
 * 0150-0001-table-stringProperties.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * A simple key-value property table
 */


create table if not exists stringProperties
(
    id    uuid primary key not null,
    value text             not null,
    foreign key (id) references propertyKeys (id)
        on delete cascade
);
