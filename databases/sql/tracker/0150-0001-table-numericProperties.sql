/*
 * 0003-0001-table-numericProperties.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * A simple key-value property table
 */


create table if not exists numericProperties
(
    id    uuid primary key not null,
    value integer          not null,
    foreign key (id) references propertyKeys (id)
        on delete cascade
);

