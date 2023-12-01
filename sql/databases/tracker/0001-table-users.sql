/*
 * 0001-table-users.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Create user profiles for the tracker.
 */

create table if not exists users
(
    -- the userId used to uniquely identify the users --
    id           uuid primary key default gen_random_uuid(),
    -- the user's first and last names --
    first_name   varchar(64)  not null,
    last_name    varchar(64)  not null,
    avatarId     uuid         not null,
    -- the user's identifying email address (globally unique) --
    email        varchar(128) not null unique,
    -- an optional phone number --
    phone_number varchar(20)  null unique,
    -- descriptive text --
    description  text,
    foreign key (avatarId) references avatars(id)
);
