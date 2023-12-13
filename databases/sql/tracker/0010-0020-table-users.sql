/*
 * 0010-0110-table-users.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Create user profiles for the tracker.
 */

create table if not exists users
(
    -- the userId used to uniquely identify the users --
    id          uuid primary key not null,
    -- the user's first and last names --
    firstName   varchar(64)      not null,
    lastName    varchar(64)      not null,
    avatarId    uuid             not null,
    created     timestamp        not null default now(),
    -- the user's identifying email address (globally unique) --
    email       varchar(256)     not null unique,
    -- an optional phone number --
    phoneNumber varchar(20)      null unique,
    -- descriptive text --
    description text,
    foreign key (avatarId) references avatars (id) on delete restrict,
    foreign key (id) references entity (id) on delete restrict
);
